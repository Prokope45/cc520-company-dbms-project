"""Tests for database rebuild helpers.

Author: GPT-5.4
Reviewed by: Josh Weese
"""

from pathlib import Path

import src.db.rebuild as rebuild


class FakeExecutor(rebuild.SqlCommandExecutor):
    """Capture executed SQL batches for rebuild helper tests."""

    def __init__(self):
        """Initialize the fake executor with an empty execution log.

        Args:
            self: Fake executor instance.

        Returns:
            None: Constructor initializes in-memory state.
        """
        self.executed = []

    def execute_query(self, sql, connection=None):
        """Record each executed batch instead of sending it to a database.

        Args:
            sql: SQL batch text.
            connection: Optional externally-managed connection.

        Returns:
            None: This method records calls for assertions.
        """
        self.executed.append(sql)


def test_execute_files_strips_utf8_bom(tmp_path: Path):
    """Ensure SQL files are read without a leading UTF-8 BOM marker.

    Args:
        tmp_path: Pytest temporary directory fixture.

    Returns:
        None: Assertions validate behavior.
    """
    sql_file = tmp_path / "script.sql"
    sql_file.write_text("\ufeffSELECT 1;", encoding="utf-8")
    executor = FakeExecutor()
    original_root = rebuild.ROOT

    rebuild.ROOT = tmp_path
    try:
        rebuild._execute_files(executor, [sql_file])
    finally:
        rebuild.ROOT = original_root

    assert executor.executed == ["SELECT 1;"]


def test_data_files_include_person_and_address_seed_files_in_order():
    """Ensure the seed file list preserves the intended execution order.

    Returns:
        None: Assertions validate behavior.
    """
    relative_data_files = [file_path.relative_to(rebuild.ROOT).as_posix() for file_path in rebuild.DATA_FILES]

    assert relative_data_files == [
        "src/garden/sql/Data/Garden.PlotStatus.sql",
        "src/garden/sql/Data/Garden.Plots.sql",
        "src/garden/sql/Data/Garden.Gardeners.sql",
        "src/garden/sql/Data/Garden.PlotAssignments.sql",
    ]


def test_split_sql_batches_respects_go_lines():
    """Ensure GO delimiters split SQL scripts into executable batches.

    Returns:
        None: Assertions validate behavior.
    """
    sql_text = """
    SELECT 1;
    GO
    SELECT 2;
    GO 5

    """

    assert rebuild._split_sql_batches(sql_text) == ["SELECT 1;", "SELECT 2;"]


def test_execute_files_executes_all_batches(tmp_path: Path):
    """Ensure every SQL batch from a file is executed in sequence.

    Args:
        tmp_path: Pytest temporary directory fixture.

    Returns:
        None: Assertions validate behavior.
    """
    sql_file = tmp_path / "script.sql"
    sql_file.write_text("SELECT 1;\nGO\nSELECT 2;", encoding="utf-8")
    executor = FakeExecutor()
    original_root = rebuild.ROOT

    rebuild.ROOT = tmp_path
    try:
        rebuild._execute_files(executor, [sql_file])
    finally:
        rebuild.ROOT = original_root

    assert executor.executed == ["SELECT 1;", "SELECT 2;"]


def test_run_rebuild_executes_all_file_groups(monkeypatch):
    """Ensure rebuild mode executes schema, tables, procedures, and seed files.

    Args:
        monkeypatch: Pytest fixture for patching module attributes.

    Returns:
        None: Assertions validate behavior.
    """
    calls = []

    class FakeSqlCommandExecutor:
        def __init__(self, **kwargs):
            self.kwargs = kwargs

    def fake_execute_files(executor, files):
        calls.append((executor.kwargs, list(files)))

    monkeypatch.setattr(rebuild, "SqlCommandExecutor", FakeSqlCommandExecutor)
    monkeypatch.setattr(rebuild, "_execute_files", fake_execute_files)

    rebuild.run_rebuild(server="s", database="d", user="u", password="p", trusted=False)

    assert len(calls) == 4
    assert calls[0][1] == rebuild.SCHEMA_FILES
    assert calls[1][1] == rebuild.TABLE_FILES
    assert calls[2][1] == rebuild.PROCEDURE_FILES
    assert calls[3][1] == rebuild.DATA_FILES
    assert calls[0][0] == {
        "server": "s",
        "database": "d",
        "user": "u",
        "password": "p",
        "trusted": False,
    }


def test_run_seed_only_executes_data_files(monkeypatch):
    """Ensure seed-only mode executes only the configured data files.

    Args:
        monkeypatch: Pytest fixture for patching module attributes.

    Returns:
        None: Assertions validate behavior.
    """
    calls = []

    class FakeSqlCommandExecutor:
        def __init__(self, **kwargs):
            self.kwargs = kwargs

    def fake_execute_files(executor, files):
        calls.append((executor.kwargs, list(files)))

    monkeypatch.setattr(rebuild, "SqlCommandExecutor", FakeSqlCommandExecutor)
    monkeypatch.setattr(rebuild, "_execute_files", fake_execute_files)

    rebuild.run_seed_only(server="s", database="d", user="u", password="p", trusted=True)

    assert len(calls) == 1
    assert calls[0][1] == rebuild.DATA_FILES
    assert calls[0][0]["trusted"] is True


def test_run_from_cli_calls_seed_only(monkeypatch):
    """Ensure CLI seed-only mode dispatches to the seed-only runner.

    Args:
        monkeypatch: Pytest fixture for patching module attributes.

    Returns:
        None: Assertions validate behavior.
    """
    called = {"seed_only": 0, "rebuild": 0}

    def fake_seed_only(**_kwargs):
        called["seed_only"] += 1

    def fake_rebuild(**_kwargs):
        called["rebuild"] += 1

    monkeypatch.setattr(rebuild, "run_seed_only", fake_seed_only)
    monkeypatch.setattr(rebuild, "run_rebuild", fake_rebuild)

    rc = rebuild.run_from_cli(["seed-only", "--server", "s", "--database", "d", "--user", "u", "--password", "p", "--no-trusted"])

    assert rc == 0
    assert called == {"seed_only": 1, "rebuild": 0}


def test_run_from_cli_calls_rebuild_by_default(monkeypatch):
    """Ensure the CLI defaults to rebuild mode when no mode is provided.

    Args:
        monkeypatch: Pytest fixture for patching module attributes.

    Returns:
        None: Assertions validate behavior.
    """
    called = {"seed_only": 0, "rebuild": 0}

    def fake_seed_only(**_kwargs):
        called["seed_only"] += 1

    def fake_rebuild(**_kwargs):
        called["rebuild"] += 1

    monkeypatch.setattr(rebuild, "run_seed_only", fake_seed_only)
    monkeypatch.setattr(rebuild, "run_rebuild", fake_rebuild)

    rc = rebuild.run_from_cli([])

    assert rc == 0
    assert called == {"seed_only": 0, "rebuild": 1}