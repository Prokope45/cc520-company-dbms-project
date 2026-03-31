# src.data_access.sql_command_executor

Generated from Python docstrings using `pydoc`.

```text
module src.data_access.sql_command_executor in src.data_access

NAME
    src.data_access.sql_command_executor

DESCRIPTION
    File: sql_command_executor.py
    Author: Josh Weese

    This file contains a class that serves as the data access layer for a MS SQL database.
    You will need to have the odbc driver for SQL server installed, which can be found here https://docs.microsoft.com/en-us/sql/connect/odbc/download-odbc-driver-for-sql-server?view=sql-server-ver15

CLASSES
    builtins.object
        SqlCommandExecutor

    class SqlCommandExecutor(builtins.object)
     |  SqlCommandExecutor(
     |      server: str = '',
     |      database: str = '',
     |      user: str = '',
     |      password: str = '',
     |      trusted: bool = True
     |  )
     |
     |  This is a general purpose class designed to interface directly with a MS SQL Server database.
     |
     |  Methods defined here:
     |
     |  __init__(
     |      self,
     |      server: str = '',
     |      database: str = '',
     |      user: str = '',
     |      password: str = '',
     |      trusted: bool = True
     |  )
     |      Constructor to initialize information to connect/work with the database.
     |      Args:
     |          server: The server to connect to.
     |          database: The database on the server to use.
     |          user: The user to connect with.
     |          password: The password for the user.
     |          trusted: Whether or not this is a trusted connection.
     |
     |  execute_query(self, sql, connection=None)
     |      Execute a raw SQL statement with optional shared connection reuse.
     |
     |      Args:
     |          sql: SQL statement to execute.
     |          connection: Optional externally-managed shared connection.
     |
     |      Returns:
     |          None: This method performs side effects only.
     |
     |  execute_stored_procedure(
     |      self,
     |      procedure_name: str,
     |      input_param_names: Sequence[str] | None = None,
     |      input_param_values: Sequence[Any] | None = None,
     |      output_param: Mapping[str, Sequence[str]] | None = None,
     |      connection=None
     |  )
     |      Execute a stored procedure and return all row-producing result sets.
     |
     |      Args:
     |          procedure_name: Fully-qualified stored procedure name.
     |          input_param_names: Optional ordered input parameter names.
     |          input_param_values: Optional ordered input parameter values.
     |          output_param: Optional output parameter mapping metadata.
     |          connection: Optional externally-managed shared connection.
     |
     |      Returns:
     |          list[Any]: Combined rows from all result sets.
     |
     |  transaction_scope(self)
     |      Yield a connection that commits on success and rolls back on failure.
     |
     |      Returns:
     |          Any: A context-managed connection for grouped operations.
     |
     |  ----------------------------------------------------------------------
     |  Static methods defined here:
     |
     |  get_all_rows(cursor: Any)
     |      Read and combine rows from every result set available on the cursor.
     |
     |      Args:
     |          cursor: Database cursor positioned on the first result set.
     |
     |      Returns:
     |          list[Any]: Aggregated rows from every row-producing result set.
     |
     |  stringify_inp_params(params: Sequence[str]) -> str
     |      Convert stored procedure input parameter names into SQL placeholders.
     |
     |      Args:
     |          params: Stored procedure parameter names.
     |
     |      Returns:
     |          str: SQL placeholder fragment for EXEC invocation.
     |
     |  stringify_out_params(params: Mapping[str, Sequence[str]]) -> tuple[str, str, str]
     |      Build SQL fragments required to declare, bind, and select output parameters.
     |
     |      Args:
     |          params: Mapping containing ``sp_local``, ``sp_local_types``, and ``sp_out`` lists.
     |
     |      Returns:
     |          tuple[str, str, str]: Output bind fragment, variable declaration fragment, and select fragment.
     |
     |  ----------------------------------------------------------------------
     |  Data descriptors defined here:
     |
     |  __dict__
     |      dictionary for instance variables
     |
     |  __weakref__
     |      list of weak references to the object

DATA
    Mapping = typing.Mapping
        A generic version of collections.abc.Mapping.

    Sequence = typing.Sequence
        A generic version of collections.abc.Sequence.

FILE
    /workspaces/database-repository-pattern-example/src/data_access/sql_command_executor.py
```
