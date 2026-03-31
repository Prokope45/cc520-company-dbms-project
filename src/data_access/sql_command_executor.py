"""
File: sql_command_executor.py
Author: Josh Weese

This file contains a class that serves as the data access layer for a MS SQL database.
You will need to have the odbc driver for SQL server installed, which can be found here https://docs.microsoft.com/en-us/sql/connect/odbc/download-odbc-driver-for-sql-server?view=sql-server-ver15

"""
from typing import Any, Mapping, Sequence
from contextlib import contextmanager

from os import getenv

from dotenv import load_dotenv
import mssql_python

load_dotenv()


class SqlCommandExecutor:
    """
    This is a general purpose class designed to interface directly with a MS SQL Server database.
    """

    def __init__(self, server: str = "", database: str = "", user: str = "",
                 password: str = "", trusted: bool = True):
        """
        Constructor to initialize information to connect/work with the database.
        Args:
            server: The server to connect to.
            database: The database on the server to use.
            user: The user to connect with.
            password: The password for the user.
            trusted: Whether or not this is a trusted connection.
        """
        resolved_database = database or getenv('DB_DATABASE')
        resolved_server = server or getenv('DB_SERVER')
        resolved_user = user or getenv('DB_USER')
        resolved_password = password or getenv('DB_PASSWORD')

        params = []
        if resolved_database:
            params.append(f"Database={resolved_database}")
        if resolved_server:
            params.append(f"Server={resolved_server}")

        # Prefer SQL authentication when credentials are available.
        if resolved_user and resolved_password:
            params.append(f"UID={resolved_user}")
            params.append(f"PWD={resolved_password}")
        elif trusted:
            params.append("Trusted_Connection=yes")

        params.append("TrustServerCertificate=yes")
        params.append("Encrypt=yes")
        self._connection_string = ";".join(params)

    def _connect(self) -> Any:
        """Create a new database connection using the configured connection string.

        Returns:
            Any: A live database connection instance.
        """
        return mssql_python.connect(self._connection_string, autocommit=False)  # type: ignore[attr-defined]

    @contextmanager
    def transaction_scope(self):
        """Yield a connection that commits on success and rolls back on failure.

        Returns:
            Any: A context-managed connection for grouped operations.
        """
        connection = self._connect()
        try:
            yield connection
            connection.commit()
        except Exception:
            connection.rollback()
            raise
        finally:
            connection.close()

    @staticmethod
    def stringify_inp_params(params: Sequence[str]) -> str:
        """Convert stored procedure input parameter names into SQL placeholders.

        Args:
            params: Stored procedure parameter names.

        Returns:
            str: SQL placeholder fragment for EXEC invocation.
        """
        return ", ".join([f"@{val} = ?" for val in params])

    @staticmethod
    def stringify_out_params(params: Mapping[str, Sequence[str]]) -> tuple[str, str, str]:
        """Build SQL fragments required to declare, bind, and select output parameters.

        Args:
            params: Mapping containing ``sp_local``, ``sp_local_types``, and ``sp_out`` lists.

        Returns:
            tuple[str, str, str]: Output bind fragment, variable declaration fragment, and select fragment.
        """
        sp_out = ", ".join([f"@{val} = @{params['sp_local'][i]} OUTPUT" for i, val in enumerate(params["sp_out"])])
        sp_local = ", ".join([f"@{val} {params['sp_local_types'][i]}" for i, val in enumerate(params["sp_local"])])
        sp_select = ", ".join([f"@{val} AS {val}_var" for val in params["sp_local"]])
        return sp_out, sp_local, sp_select

    def execute_stored_procedure(self, procedure_name: str,
                                 input_param_names: Sequence[str] | None = None,
                                 input_param_values: Sequence[Any] | None = None,
                                 output_param: Mapping[str, Sequence[str]] | None = None,
                                 connection=None):
        """Execute a stored procedure and return all row-producing result sets.

        Args:
            procedure_name: Fully-qualified stored procedure name.
            input_param_names: Optional ordered input parameter names.
            input_param_values: Optional ordered input parameter values.
            output_param: Optional output parameter mapping metadata.
            connection: Optional externally-managed shared connection.

        Returns:
            list[Any]: Combined rows from all result sets.
        """
        owns_connection = connection is None
        active_connection = connection or self._connect()
        try:
            cursor = active_connection.cursor()
            if input_param_names is None and output_param is None:
                sql = f"""
                        EXEC {procedure_name};
                    """
                cursor.execute(sql)
            elif input_param_names is not None and output_param is None:
                if input_param_values is None:
                    raise ValueError("input_param_values cannot be None when input_param_names is provided")
                sql = f"""
                        EXEC {procedure_name} {self.stringify_inp_params(input_param_names)};
                    """
                cursor.execute(sql, input_param_values)
            elif input_param_names is None and output_param is not None:
                sp_out, sp_local, sp_select = self.stringify_out_params(output_param)
                sql = f"""
                        DECLARE {sp_local};
                        EXEC {procedure_name} {sp_out};
                        SELECT {sp_select};            
                    """
                cursor.execute(sql)
            else:
                if input_param_values is None:
                    raise ValueError("input_param_values cannot be None when input_param_names is provided")
                assert input_param_names is not None
                assert output_param is not None
                sp_out, sp_local, sp_select = self.stringify_out_params(output_param)
                sql = f"""
                        DECLARE {sp_local};
                        EXEC {procedure_name} {self.stringify_inp_params(input_param_names)}, {sp_out};
                        SELECT {sp_select};            
                    """
                cursor.execute(sql, input_param_values)
            results = self.get_all_rows(cursor)
            if owns_connection:
                active_connection.commit()
            return results
        except Exception:
            if owns_connection:
                active_connection.rollback()
            raise
        finally:
            if owns_connection:
                active_connection.close()

    def execute_query(self, sql, connection=None):
        """Execute a raw SQL statement with optional shared connection reuse.

        Args:
            sql: SQL statement to execute.
            connection: Optional externally-managed shared connection.

        Returns:
            None: This method performs side effects only.
        """
        owns_connection = connection is None
        active_connection = connection or self._connect()
        try:
            cursor = active_connection.cursor()
            cursor.execute(sql)
            if owns_connection:
                active_connection.commit()
        except Exception:
            if owns_connection:
                active_connection.rollback()
            raise
        finally:
            if owns_connection:
                active_connection.close()

    @staticmethod
    def get_all_rows(cursor: Any):
        """Read and combine rows from every result set available on the cursor.

        Args:
            cursor: Database cursor positioned on the first result set.

        Returns:
            list[Any]: Aggregated rows from every row-producing result set.
        """
        results = []
        while True:
            if cursor.description is not None:
                results += cursor.fetchall()
            if not cursor.nextset():
                break
        return results
