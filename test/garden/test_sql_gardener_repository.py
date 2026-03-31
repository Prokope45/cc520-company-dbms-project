"""Unit tests for SqlGardenerRepository without external DB dependencies.

Author: Josh Weese

"""

from types import SimpleNamespace

import pytest

from src.data_access.RecordNotFoundException import RecordNotFoundException
from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.sql_gardener_repository import SqlGardenerRepository


class FakeExecutor(SqlCommandExecutor):
    """Return canned stored procedure results for repository tests."""

    def __init__(self):
        """Initialize fake result storage and call tracking.

        Returns:
            None: Constructor initializes in-memory state.
        """
        self.results_by_sp = {}
        self.calls = []

    def execute_stored_procedure(self, procedure_name, input_param_names=None, input_param_values=None,
                                 output_param=None, connection=None):
        """Record a stored procedure call and return the configured fake result set.

        Args:
            procedure_name: Stored procedure name.
            input_param_names: Optional input parameter names.
            input_param_values: Optional input parameter values.
            output_param: Optional output parameter metadata.
            connection: Optional externally-managed connection.

        Returns:
            list[Any]: Configured fake result set.
        """
        self.calls.append((procedure_name, input_param_names, input_param_values, output_param))
        return self.results_by_sp.get(procedure_name, [])


@pytest.fixture
def gardener_repo():
    """Return a gardener repository wired to a fake executor.

    Returns:
        SqlGardenerRepository: Repository instance with fake executor injected.
    """
    repo = SqlGardenerRepository()
    repo.executor = FakeExecutor()
    return repo


class TestSqlGardenerRepository:
    """Verify the SQL gardener repository maps and validates data correctly."""

    def test_create_gardener(self, gardener_repo):
        """Ensure gardener creation returns a mapped ``Gardener`` instance.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.CreateGardener"] = [SimpleNamespace(GardenerId_var=42)]

        gardener = gardener_repo.create_gardener("Chili", "Heeler", "555-0101", "chili@test.com", "2024-03-01")

        assert gardener is not None
        assert gardener.gardener_id == 42
        assert gardener.first_name == "Chili"
        assert gardener.last_name == "Heeler"
        assert gardener.email == "chili@test.com"

    def test_fetch_none_gardener(self, gardener_repo):
        """Ensure missing gardeners raise ``RecordNotFoundException``.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.FetchGardener"] = []

        with pytest.raises(RecordNotFoundException):
            gardener_repo.fetch_gardener(0)

    def test_fetch_gardener(self, gardener_repo):
        """Ensure a fetched row is translated into a ``Gardener`` instance.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.FetchGardener"] = [
            SimpleNamespace(GardenerId=7, FirstName="Chili", LastName="Heeler", Phone="555-0101", Email="chili@test.com", JoinDate="2024-03-01")
        ]

        gardener = gardener_repo.fetch_gardener(gardener_id=7)

        assert gardener.gardener_id == 7
        assert gardener.first_name == "Chili"
        assert gardener.last_name == "Heeler"
        assert gardener.email == "chili@test.com"

    def test_get_gardener(self, gardener_repo):
        """Ensure email lookup returns the mapped ``Gardener`` instance.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.GetGardenerByEmail"] = [
            SimpleNamespace(GardenerId=5, FirstName="Chili", LastName="Heeler", Phone="555-0101", Email="chili@test.com", JoinDate="2024-03-01")
        ]

        gardener = gardener_repo.get_gardener_by_email(email="chili@test.com")

        assert gardener is not None
        assert gardener.gardener_id == 5
        assert gardener.email == "chili@test.com"

    def test_retrieve_gardeners(self, gardener_repo):
        """Ensure multiple gardeners are translated into a list of models.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.RetrieveGardeners"] = [
            SimpleNamespace(GardenerId=1, FirstName="A", LastName="One", Phone="555-0001", Email="a@test.com", JoinDate="2024-01-01"),
            SimpleNamespace(GardenerId=2, FirstName="B", LastName="Two", Phone="555-0002", Email="b@test.com", JoinDate="2024-01-02"),
            SimpleNamespace(GardenerId=3, FirstName="C", LastName="Three", Phone="555-0003", Email="c@test.com", JoinDate="2024-01-03"),
        ]

        gardeners = gardener_repo.retrieve_gardeners()

        assert gardeners is not None
        assert len(gardeners) == 3
        assert gardeners[0].gardener_id == 1
        assert gardeners[1].email == "b@test.com"
        assert gardeners[2].last_name == "Three"

    def test_create_gardener_returns_none_when_no_single_output(self, gardener_repo):
        """Ensure create returns ``None`` when the stored procedure yields no single output row.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.CreateGardener"] = []

        assert gardener_repo.create_gardener("Chili", "Heeler", "555-0101", "chili@test.com", "2024-03-01") is None

    @pytest.mark.parametrize(
        "first,last,expected",
        [
            ("", "Heeler", "First name cannot be empty."),
            ("Chili", "", "Last name cannot be empty."),
            (None, "Heeler", "First name cannot be empty."),
            ("Chili", None, "Last name cannot be empty."),
        ],
    )
    def test_create_gardener_validates_required_fields(self, gardener_repo, first, last, expected):
        """Ensure required fields are validated before issuing the stored procedure call.

        Args:
            gardener_repo: Repository fixture under test.
            first: Candidate first name.
            last: Candidate last name.
            expected: Expected validation message pattern.

        Returns:
            None: Assertions validate behavior.
        """
        with pytest.raises(ValueError, match=expected):
            gardener_repo.create_gardener(first, last, "555-0101", "chili@test.com", "2024-03-01")

    def test_retrieve_gardeners_returns_none_when_no_rows(self, gardener_repo):
        """Ensure retrieving gardeners returns ``None`` when no rows are returned.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.RetrieveGardeners"] = []

        assert gardener_repo.retrieve_gardeners() is None

    def test_get_gardener_returns_none_when_missing(self, gardener_repo):
        """Ensure email lookup returns ``None`` when no row is returned.

        Args:
            gardener_repo: Repository fixture under test.

        Returns:
            None: Assertions validate behavior.
        """
        gardener_repo.executor.results_by_sp["Garden.GetGardenerByEmail"] = []

        assert gardener_repo.get_gardener_by_email("missing@test.com") is None

