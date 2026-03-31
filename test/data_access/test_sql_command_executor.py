"""Unit tests for SqlCommandExecutor connection string assembly.

Author: GPT-5.4
Reviewed by: Josh Weese
"""

import pytest

from src.data_access.sql_command_executor import SqlCommandExecutor


class TestSqlCommandExecutor:
    """Verify connection string assembly and execution lifecycle behavior."""

    @staticmethod
    def _build_fake_db(should_fail: bool = False):
        """Create fake connection/cursor objects used by execution-path tests.

        Args:
            should_fail: Whether fake cursor execution should raise runtime errors.

        Returns:
            Any: Fake connection instance with counters and captured executions.
        """
        class FakeCursor:
            """Capture executed SQL and optionally simulate database errors."""

            def __init__(self):
                self.description = [object()]
                self.executions = []
                self.nextset_calls = 0

            def execute(self, sql, params=None):
                self.executions.append((sql, params))
                if should_fail:
                    raise RuntimeError("db error")

            def fetchall(self):
                return ["row"]

            def nextset(self):
                self.nextset_calls += 1
                return False

        class FakeConnection:
            """Track commit, rollback, and close calls for assertions."""

            def __init__(self):
                self.cursor_instance = FakeCursor()
                self.commit_calls = 0
                self.rollback_calls = 0
                self.close_calls = 0

            def cursor(self):
                return self.cursor_instance

            def commit(self):
                self.commit_calls += 1

            def rollback(self):
                self.rollback_calls += 1

            def close(self):
                self.close_calls += 1

        return FakeConnection()

    def test_prefers_sql_auth_when_credentials_are_present(self, monkeypatch):
        """Ensure SQL credentials override trusted-connection settings.

        Args:
            monkeypatch: Pytest fixture used to set environment variables.

        Returns:
            None: Assertions validate behavior.
        """
        monkeypatch.setenv("SERVER", "localhost")
        monkeypatch.setenv("DATABASE", "cc520")
        monkeypatch.setenv("USER", "cc520-admin")
        monkeypatch.setenv("PASSWORD", "database3ssentials!")

        executor = SqlCommandExecutor(trusted=True)

        assert "Server=localhost" in executor._connection_string
        assert "Database=cc520" in executor._connection_string
        assert "UID=cc520-admin" in executor._connection_string
        assert "PWD=database3ssentials!" in executor._connection_string
        assert "TrustedConnection=yes" not in executor._connection_string
        assert "Trusted_Connection=yes" not in executor._connection_string

    def test_uses_trusted_connection_when_credentials_are_absent(self, monkeypatch):
        """Ensure trusted connection is used when explicit credentials are missing.

        Args:
            monkeypatch: Pytest fixture used to set environment variables.

        Returns:
            None: Assertions validate behavior.
        """
        monkeypatch.setenv("SERVER", "localhost")
        monkeypatch.setenv("DATABASE", "cc520")
        monkeypatch.delenv("USER", raising=False)
        monkeypatch.delenv("PASSWORD", raising=False)

        executor = SqlCommandExecutor(trusted=True)

        assert "Server=localhost" in executor._connection_string
        assert "Database=cc520" in executor._connection_string
        assert "Trusted_Connection=yes" in executor._connection_string
        assert "UID=" not in executor._connection_string
        assert "PWD=" not in executor._connection_string

    def test_get_all_rows_collects_results_after_non_row_result_sets(self):
        """Ensure row-producing result sets after non-row sets are still collected.

        Returns:
            None: Assertions validate behavior.
        """
        class FakeCursor:
            def __init__(self):
                self.index = 0
                self.descriptions = [None, [object()], [object()]]
                self.result_sets = [[], [1, 2], [3]]

            @property
            def description(self):
                return self.descriptions[self.index]

            def fetchall(self):
                return self.result_sets[self.index]

            def nextset(self):
                self.index += 1
                return self.index < len(self.descriptions)

        assert SqlCommandExecutor.get_all_rows(FakeCursor()) == [1, 2, 3]

    def test_get_all_rows_collects_first_result_set(self):
        """Ensure a single initial result set is returned correctly.

        Returns:
            None: Assertions validate behavior.
        """
        class FakeCursor:
            def __init__(self):
                self.index = 0
                self.descriptions = [[object()]]
                self.result_sets = [[1]]

            @property
            def description(self):
                return self.descriptions[self.index]

            def fetchall(self):
                return self.result_sets[self.index]

            def nextset(self):
                self.index += 1
                return False

        assert SqlCommandExecutor.get_all_rows(FakeCursor()) == [1]

    def test_stringify_inp_params_formats_names(self):
        """Ensure input parameter names are converted into placeholder syntax.

        Returns:
            None: Assertions validate behavior.
        """
        assert SqlCommandExecutor.stringify_inp_params(["Foo", "Bar"]) == "@Foo = ?, @Bar = ?"

    def test_stringify_out_params_formats_segments(self):
        """Ensure output parameter fragments are generated consistently.

        Returns:
            None: Assertions validate behavior.
        """
        sp_out, sp_local, sp_select = SqlCommandExecutor.stringify_out_params({
            "sp_local": ["PersonId"],
            "sp_local_types": ["int"],
            "sp_out": ["PersonId"],
        })

        assert sp_out == "@PersonId = @PersonId OUTPUT"
        assert sp_local == "@PersonId int"
        assert sp_select == "@PersonId AS PersonId_var"

    def test_execute_stored_procedure_without_params(self, monkeypatch):
        """Ensure stored procedures with no parameters execute and commit successfully.

        Args:
            monkeypatch: Pytest fixture used to patch database connection creation.

        Returns:
            None: Assertions validate behavior.
        """
        fake_connection = self._build_fake_db()
        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", lambda *_a, **_k: fake_connection)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        results = executor.execute_stored_procedure("Garden.RetrieveGardeners")

        sql, params = fake_connection.cursor_instance.executions[0]
        assert "EXEC Garden.RetrieveGardeners" in sql
        assert params is None
        assert results == ["row"]
        assert fake_connection.commit_calls == 1
        assert fake_connection.close_calls == 1

    def test_execute_stored_procedure_with_input_params(self, monkeypatch):
        """Ensure input parameters are passed through to the stored procedure call."""
        fake_connection = self._build_fake_db()
        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", lambda *_a, **_k: fake_connection)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        executor.execute_stored_procedure(
            "Garden.FetchGardener",
            input_param_names=["GardenerId"],
            input_param_values=[7],
        )

        sql, params = fake_connection.cursor_instance.executions[0]
        assert "@GardenerId = ?" in sql
        assert params == [7]

    def test_execute_stored_procedure_with_output_params_only(self, monkeypatch):
        """Ensure output-only stored procedures declare and select output values."""
        fake_connection = self._build_fake_db()
        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", lambda *_a, **_k: fake_connection)
        out_param = {
            "sp_local": ["GardenerId"],
            "sp_local_types": ["int"],
            "sp_out": ["GardenerId"],
        }

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        executor.execute_stored_procedure("Garden.CreateGardener", output_param=out_param)

        sql, params = fake_connection.cursor_instance.executions[0]
        assert "DECLARE @GardenerId int" in sql
        assert "SELECT @GardenerId AS GardenerId_var" in sql
        assert params is None

    def test_execute_stored_procedure_with_input_and_output_params(self, monkeypatch):
        """Ensure mixed input and output parameters are formatted correctly."""
        fake_connection = self._build_fake_db()
        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", lambda *_a, **_k: fake_connection)
        out_param = {
            "sp_local": ["GardenerId"],
            "sp_local_types": ["int"],
            "sp_out": ["GardenerId"],
        }

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        executor.execute_stored_procedure(
            "Garden.CreateGardener",
            input_param_names=["FirstName", "LastName", "Email"],
            input_param_values=["Chili", "Heeler", "chilih@test.com"],
            output_param=out_param,
        )

        sql, params = fake_connection.cursor_instance.executions[0]
        assert "@FirstName = ?" in sql
        assert "@GardenerId = @GardenerId OUTPUT" in sql
        assert params == ["Chili", "Heeler", "chilih@test.com"]

    def test_execute_stored_procedure_raises_on_missing_input_values(self):
        """Ensure missing input values raise a validation error."""
        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)

        with pytest.raises(ValueError):
            executor.execute_stored_procedure("Garden.FetchGardener", input_param_names=["GardenerId"])

    def test_execute_stored_procedure_rolls_back_on_error(self, monkeypatch):
        """Ensure stored procedure failures roll back owned connections."""
        fake_connection = self._build_fake_db(should_fail=True)
        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", lambda *_a, **_k: fake_connection)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)

        with pytest.raises(RuntimeError):
            executor.execute_stored_procedure("Garden.RetrieveGardeners")

        assert fake_connection.commit_calls == 0
        assert fake_connection.rollback_calls == 1
        assert fake_connection.close_calls == 1

    def test_execute_stored_procedure_external_connection_not_finalized(self):
        """Ensure externally supplied connections are not finalized by the executor."""
        fake_connection = self._build_fake_db()
        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)

        executor.execute_stored_procedure("Garden.RetrieveGardeners", connection=fake_connection)

        assert fake_connection.commit_calls == 0
        assert fake_connection.rollback_calls == 0
        assert fake_connection.close_calls == 0

    def test_execute_query_owns_connection_and_commits(self, monkeypatch):
        """Ensure raw queries commit and close when the executor owns the connection."""
        class FakeCursor:
            def execute(self, *_args, **_kwargs):
                return None

        class FakeConnection:
            def __init__(self):
                self.commit_calls = 0
                self.rollback_calls = 0
                self.close_calls = 0

            def cursor(self):
                return FakeCursor()

            def commit(self):
                self.commit_calls += 1

            def rollback(self):
                self.rollback_calls += 1

            def close(self):
                self.close_calls += 1

        fake_connection = FakeConnection()

        def fake_connect(*_args, **_kwargs):
            return fake_connection

        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", fake_connect)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        executor.execute_query("SELECT 1")

        assert fake_connection.commit_calls == 1
        assert fake_connection.rollback_calls == 0
        assert fake_connection.close_calls == 1

    def test_execute_query_with_external_connection_does_not_finalize(self):
        """Ensure raw queries do not finalize externally supplied connections."""
        class FakeCursor:
            def execute(self, *_args, **_kwargs):
                return None

        class FakeConnection:
            def __init__(self):
                self.commit_calls = 0
                self.rollback_calls = 0
                self.close_calls = 0

            def cursor(self):
                return FakeCursor()

            def commit(self):
                self.commit_calls += 1

            def rollback(self):
                self.rollback_calls += 1

            def close(self):
                self.close_calls += 1

        fake_connection = FakeConnection()
        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)

        executor.execute_query("SELECT 1", connection=fake_connection)

        assert fake_connection.commit_calls == 0
        assert fake_connection.rollback_calls == 0
        assert fake_connection.close_calls == 0

    def test_execute_query_rolls_back_on_error(self, monkeypatch):
        """Ensure raw query failures roll back owned connections."""
        class FakeCursor:
            def execute(self, *_args, **_kwargs):
                raise RuntimeError("db error")

        class FakeConnection:
            def __init__(self):
                self.commit_calls = 0
                self.rollback_calls = 0
                self.close_calls = 0

            def cursor(self):
                return FakeCursor()

            def commit(self):
                self.commit_calls += 1

            def rollback(self):
                self.rollback_calls += 1

            def close(self):
                self.close_calls += 1

        fake_connection = FakeConnection()
        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", lambda *_a, **_k: fake_connection)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        with pytest.raises(RuntimeError):
            executor.execute_query("SELECT 1")

        assert fake_connection.commit_calls == 0
        assert fake_connection.rollback_calls == 1
        assert fake_connection.close_calls == 1

    def test_transaction_scope_commits_then_closes(self, monkeypatch):
        """Ensure the transaction scope commits and closes after successful work."""
        class FakeConnection:
            def __init__(self):
                self.commit_calls = 0
                self.rollback_calls = 0
                self.close_calls = 0

            def commit(self):
                self.commit_calls += 1

            def rollback(self):
                self.rollback_calls += 1

            def close(self):
                self.close_calls += 1

        fake_connection = FakeConnection()

        def fake_connect(*_args, **_kwargs):
            return fake_connection

        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", fake_connect)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)
        with executor.transaction_scope() as connection:
            assert connection is fake_connection

        assert fake_connection.commit_calls == 1
        assert fake_connection.rollback_calls == 0
        assert fake_connection.close_calls == 1

    def test_transaction_scope_rolls_back_then_closes(self, monkeypatch):
        """Ensure the transaction scope rolls back and closes after failures."""
        class FakeConnection:
            def __init__(self):
                self.commit_calls = 0
                self.rollback_calls = 0
                self.close_calls = 0

            def commit(self):
                self.commit_calls += 1

            def rollback(self):
                self.rollback_calls += 1

            def close(self):
                self.close_calls += 1

        fake_connection = FakeConnection()

        def fake_connect(*_args, **_kwargs):
            return fake_connection

        monkeypatch.setattr("src.data_access.sql_command_executor.mssql_python.connect", fake_connect)

        executor = SqlCommandExecutor(server="localhost", database="cc520", trusted=True)

        try:
            with executor.transaction_scope():
                raise RuntimeError("boom")
        except RuntimeError:
            pass

        assert fake_connection.commit_calls == 0
        assert fake_connection.rollback_calls == 1
        assert fake_connection.close_calls == 1