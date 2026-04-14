"""Database rebuild/seed command runner.

Author: Josh Weese
"""

from __future__ import annotations

import argparse
import re
from os import getenv
from pathlib import Path
from typing import Iterable

from dotenv import load_dotenv

from src.data_access.sql_command_executor import SqlCommandExecutor

load_dotenv()


ROOT = Path(__file__).resolve().parents[2]
SQL_ROOT = ROOT / "src" / "garden" / "sql"
GARDEN_SCHEMA = "Garden"

SCHEMA_FILES = [
    Path(f"{SQL_ROOT}/Schemas/Garden.sql"),
]

TABLE_FILES = [
    Path(f"{SQL_ROOT}/Tables/Garden.DropTables.sql"),
    Path(f"{SQL_ROOT}/Tables/Garden.PlotStatus.sql"),
    Path(f"{SQL_ROOT}/Tables/Garden.Gardeners.sql"),
    Path(f"{SQL_ROOT}/Tables/Garden.Plots.sql"),
    Path(f"{SQL_ROOT}/Tables/Garden.PlotAssignments.sql"),
]

PROCEDURE_FILES = [
    Path(f"{SQL_ROOT}/Procedures/Garden.CreateGardener.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.RetrieveGardeners.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.FetchGardener.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.GetGardenerByEmail.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.SavePlotAssignment.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.RetrieveAssignmentsForGardener.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.RetrievePlots.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.RetrievePlotStatus.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.RetrieveAvailablePlots.sql"),
    Path(f"{SQL_ROOT}/Procedures/Garden.UpdatePlotStatus.sql"),
]

DATA_FILES = [
    Path(f"{SQL_ROOT}/Data/Garden.PlotStatus.sql"),
    Path(f"{SQL_ROOT}/Data/Garden.Plots.sql"),
    Path(f"{SQL_ROOT}/Data/Garden.Gardeners.sql"),
    Path(f"{SQL_ROOT}/Data/Garden.PlotAssignments.sql"),
]

GO_LINE = re.compile(r"^\s*GO(?:\s+\d+)?\s*$", re.IGNORECASE)


def _split_sql_batches(sql_text: str) -> list[str]:
    """Split a SQL script into executable batches using GO delimiters.

    Args:
        sql_text: Raw SQL script text.

    Returns:
        list[str]: Ordered SQL batches with delimiters removed.
    """
    batches: list[str] = []
    current_lines: list[str] = []
    for line in sql_text.splitlines():
        if GO_LINE.match(line):
            batch = "\n".join(current_lines).strip()
            if batch:
                batches.append(batch)
            current_lines = []
            continue
        current_lines.append(line)

    tail = "\n".join(current_lines).strip()
    if tail:
        batches.append(tail)
    return batches


def _execute_files(executor: SqlCommandExecutor, files: Iterable[Path]) -> None:
    """Execute each SQL file in order, running every batch it contains.

    Args:
        executor: SQL command executor used to run statements.
        files: Ordered file paths to execute.

    Returns:
        None: This function performs database side effects only.
    """
    for file_path in files:
        sql_text = file_path.read_text(encoding="utf-8").lstrip("\ufeff")
        print(f"Executing {file_path.relative_to(ROOT)}")
        for batch in _split_sql_batches(sql_text):
            executor.execute_query(batch)


def run_rebuild(*, server: str = "", database: str = "", user: str = "", password: str = "", trusted: bool = True) -> None:
    """Run schema + table + procedure + seed scripts in order.

    Args:
        server: Optional SQL Server host override.
        database: Optional database name override.
        user: Optional SQL login username.
        password: Optional SQL login password.
        trusted: Whether trusted authentication should be used when credentials are absent.

    Returns:
        None: This function performs database side effects only.
    """
    resolved_server = server or getenv('DB_SERVER', '')
    resolved_database = database or getenv('DB_DATABASE', '')
    resolved_user = user or getenv('DB_USER', '')
    resolved_password = password or getenv('DB_PASSWORD', '')
    executor = SqlCommandExecutor(
        server=resolved_server,
        database=resolved_database,
        user=resolved_user,
        password=resolved_password,
        trusted=trusted,
    )

    _execute_files(executor, SCHEMA_FILES)
    _execute_files(executor, TABLE_FILES)
    _execute_files(executor, PROCEDURE_FILES)
    _execute_files(executor, DATA_FILES)


def run_seed_only(*, server: str = "", database: str = "", user: str = "", password: str = "", trusted: bool = True) -> None:
    """Run only data seed scripts.

    Args:
        server: Optional SQL Server host override.
        database: Optional database name override.
        user: Optional SQL login username.
        password: Optional SQL login password.
        trusted: Whether trusted authentication should be used when credentials are absent.

    Returns:
        None: This function performs database side effects only.
    """
    executor = SqlCommandExecutor(
        server=server,
        database=database,
        user=user,
        password=password,
        trusted=trusted,
    )
    _execute_files(executor, DATA_FILES)


def run_from_cli(argv: list[str] | None = None) -> int:
    """Parse command-line arguments and run the requested database action.

    Args:
        argv: Optional argument vector to parse instead of ``sys.argv``.

    Returns:
        int: Process-style exit code.
    """
    parser = argparse.ArgumentParser(description="Rebuild or seed the Garden database schema.")
    parser.add_argument("mode", choices=["rebuild", "seed-only"], nargs="?", default="rebuild")
    parser.add_argument("--server", default="")
    parser.add_argument("--database", default="")
    parser.add_argument("--user", default="")
    parser.add_argument("--password", default="")
    parser.add_argument("--trusted", action=argparse.BooleanOptionalAction, default=True)

    args = parser.parse_args(argv)

    print(f"Running database mode: {args.mode}")
    if args.mode == "seed-only":
        run_seed_only(
            server=args.server,
            database=args.database,
            user=args.user,
            password=args.password,
            trusted=args.trusted,
        )
    else:
        run_rebuild(
            server=args.server,
            database=args.database,
            user=args.user,
            password=args.password,
            trusted=args.trusted,
        )

    print("Database command completed.")
    return 0
