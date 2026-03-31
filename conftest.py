"""
Pytest configuration for the project.

Author: Josh Weese
"""
import pytest
import subprocess
import sys

from src.garden.sql_plot_assignment_repository import SqlPlotAssignmentRepository
from src.garden.sql_gardener_repository import SqlGardenerRepository


@pytest.fixture(scope="session", autouse=True)
def setup():
    """Run one-time test session setup.

    Returns:
        None: Fixture performs setup side effects only.
    """
    reset_test_database()


def reset_test_database():
    """Rebuild the database so tests start from a known state.

    Returns:
        None: Function performs process/database side effects only.
    """
    print("\n\nReloading database....")

    try:
        cmd = [
            sys.executable,
            "-m",
            "src",
            "rebuild-db",
            "rebuild",
            "--trusted",
        ]
        result = subprocess.run(cmd, capture_output=True, text=True, check=False)
        print(result.stdout)
        if result.stderr:
            print("STDERR:", result.stderr, file=sys.stderr)
        if result.returncode != 0:
            print(f"Database rebuild failed with exit code {result.returncode}")
    except Exception as e:
        print(f"Error rebuilding database: {e}")

    print("done.\n\n")


@pytest.fixture(scope="function")
def assignment_repo():
    """Provide an assignment repository and reset database after each test.

    Returns:
        Iterator[SqlPlotAssignmentRepository]: Function-scoped repository fixture.
    """
    yield SqlPlotAssignmentRepository()
    reset_test_database()


@pytest.fixture(scope="function")
def gardener_repo():
    """Provide a gardener repository and reset database after each test.

    Returns:
        Iterator[SqlGardenerRepository]: Function-scoped repository fixture.
    """
    yield SqlGardenerRepository()
    reset_test_database()
