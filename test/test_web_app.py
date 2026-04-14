"""Tests for the Flask application.

Includes tests for page rendering and API endpoints, using fake repositories to isolate from database dependencies.
"""

import pytest

from src.data_access.RecordNotFoundException import RecordNotFoundException
from src.garden.i_plot_assignment_repository import IPlotAssignmentRepository
from src.garden.i_gardener_repository import IGardenerRepository
from src.garden.models.plot_assignment import PlotAssignment
from src.garden.models.gardener import Gardener
from src.garden.models.plot import Plot
from src.garden.models.plot_status import PlotStatus
from src.web.app import create_app


class FakeGardenerRepository(IGardenerRepository):
    """Provide deterministic gardener data for web-layer tests."""

    def create_gardener(self, first_name: str, last_name: str, phone: str | None, email: str | None, join_date: str | None):
        """Raise because create operations are outside the scope of these tests.

        Args:
            first_name: Gardener first name.
            last_name: Gardener last name.
            phone: Optional phone number.
            email: Optional email address.
            join_date: Optional join date.

        Returns:
            None: This method always raises.
        """
        raise NotImplementedError

    def retrieve_gardeners(self):
        """Return a fixed list of gardeners for API assertions.

        Returns:
            list[Gardener]: Deterministic gardeners payload for tests.
        """
        return [
            Gardener(1, "Ada", "Lovelace", "555-0101", "ada@example.com", "2024-03-12"),
            Gardener(2, "Grace", "Hopper", "555-0102", "grace@example.com", "2024-04-03"),
        ]

    def fetch_gardener(self, gardener_id: int):
        """Return a known gardener or raise when the identifier is unknown.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Gardener: Matching gardener model for known identifiers.
        """
        if gardener_id == 1:
            return Gardener(1, "Ada", "Lovelace", "555-0101", "ada@example.com", "2024-03-12")
        raise RecordNotFoundException(gardener_id)

    def get_gardener_by_email(self, email: str):
        """Raise because email lookups are outside the scope of these tests.

        Args:
            email: Email address to search by.

        Returns:
            None: This method always raises.
        """
        raise NotImplementedError


class FakePlotRepository(IPlotRepository):
    """Provide deterministic plot data for web-layer tests."""

    def retrieve_plots(self):
        """Return a fixed list of all plots for API assertions.

        Returns:
            list[Plot]: Deterministic plots payload for tests.
        """
        return [
            Plot(1, "A1", "Northwest Corner", 100, False, 1),
            Plot(2, "A2", "Northwest Corner - A2", 100, False, 1),
            Plot(3, "B1", "Northeast Corner", 100, True, 1),
            Plot(4, "B2", "Northeast Corner - B2", 100, True, 2),
        ]

    def retrieve_available_plots(self):
        """Return a fixed list of available plots for API assertions.

        Returns:
            list[Plot]: Deterministic available plots payload for tests.
        """
        return [
            Plot(1, "A1", "Northwest Corner", 100, False, 1),
            Plot(3, "B1", "Northeast Corner", 100, True, 1),
        ]

    def update_plot_status(self, plot_id: int, new_status_id: int) -> bool:
        """Update the status of a specific plot.

        Args:
            plot_id: Plot identifier.
            new_status_id: New status identifier.

        Returns:
            bool: True if update succeeded, False otherwise.
        """
        # Simulate success for valid plot IDs
        return plot_id in [1, 2, 3, 4]

    def retrieve_plot_status(self, plot_id: int) -> PlotStatus | None:
        """Return the plot status for a given plot ID.

        Args:
            plot_id: Plot identifier.

        Returns:
            PlotStatus: Plot status for known plot IDs, None otherwise.
        """
        status_map = {
            1: PlotStatus(1, "Active"),
            2: PlotStatus(2, "Maintenance"),
            3: PlotStatus(1, "Active"),
            4: PlotStatus(2, "Maintenance"),
        }
        return status_map.get(plot_id)


class FakePlotAssignmentRepository(IPlotAssignmentRepository):
    """Provide deterministic assignment data for web-layer tests."""

    def get_assignments(self, gardener_id: int):
        """Return fixed assignments for the known test gardener.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            list[PlotAssignment]: Deterministic assignment payload for tests.
        """
        if gardener_id == 1:
            return [
                PlotAssignment(1, 2, 1, "2024-04-01", None, "Spring crops", "A2", "Northwest Corner - A2", 100, False, "Active"),
            ]
        return []

    def save_assignment(self, plot_id: int, gardener_id: int, start_date: str, end_date: str | None, notes: str | None):
        """Raise because write operations are outside the scope of these tests.

        Args:
            plot_id: Plot identifier.
            gardener_id: Gardener identifier.
            start_date: Assignment start date.
            end_date: Optional assignment end date.
            notes: Optional assignment notes.

        Returns:
            None: This method always raises.
        """
        raise NotImplementedError


@pytest.fixture
def app_with_fakes():
    """Return a Flask app configured with fake repositories for API tests.

    Returns:
        Flask: Configured app instance using fake repositories.
    """
    app = create_app()
    app.config["GARDENER_REPOSITORY"] = FakeGardenerRepository()
    app.config["PLOT_ASSIGNMENT_REPOSITORY"] = FakePlotAssignmentRepository()
    app.config["PLOT_REPOSITORY"] = FakePlotRepository()
    return app


@pytest.fixture
def client():
    """Return a default Flask test client for page rendering tests.

    Returns:
        FlaskClient: Test client for route requests.
    """
    return create_app().test_client()


@pytest.fixture
def client_with_fakes(app_with_fakes):
    """Return a Flask test client backed by fake repositories.

    Args:
        app_with_fakes: Flask app fixture configured with fakes.

    Returns:
        FlaskClient: Test client for route requests.
    """
    return app_with_fakes.test_client()


class TestWebApp:
    """Verify Flask page rendering and API responses."""

    def test_index_page_loads(self, client):
        """Confirm the landing page renders expected content.

        Args:
            client: Default Flask test client.

        Returns:
            None: Assertions validate behavior.
        """
        response = client.get("/")

        assert response.status_code == 200
        assert b"Community Garden Repository Example" in response.data
        assert b'href="/gardeners"' in response.data
        assert b'css/index.css' in response.data

    def test_gardeners_page_loads(self, client):
        """Confirm the gardeners page renders expected content.

        Args:
            client: Default Flask test client.

        Returns:
            None: Assertions validate behavior.
        """
        response = client.get("/gardeners")

        assert response.status_code == 200
        assert b"Gardeners" in response.data
        assert b'href="/"' in response.data
        assert b'css/gardeners.css' in response.data
        assert b'id="gardeners-table"' in response.data
        assert b'datatables' in response.data

    def test_retrieve_gardeners_api(self, client_with_fakes):
        """Confirm the gardeners API returns the fake repository payload.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.get("/api/gardeners")

        assert response.status_code == 200
        payload = response.get_json()
        assert len(payload) == 2
        assert payload[0]["gardenerId"] == 1
        assert payload[0]["email"] == "ada@example.com"

    def test_retrieve_assignments_for_gardener_api(self, client_with_fakes):
        """Confirm assignments API returns gardener and assignment payloads.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.get("/api/gardeners/1/assignments")

        assert response.status_code == 200
        payload = response.get_json()
        assert payload["gardener"]["gardenerId"] == 1
        assert len(payload["assignments"]) == 1
        assert payload["assignments"][0]["locationDescription"] == "Northwest Corner - A2"

    def test_retrieve_assignments_for_missing_gardener_returns_404(self, client_with_fakes):
        """Confirm missing gardeners produce a 404 response.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.get("/api/gardeners/99/assignments")

        assert response.status_code == 404


class TestPlotAPI:
    """Verify plot-related API endpoints."""

    def test_retrieve_plots_api(self, client_with_fakes):
        """Confirm the retrieve_plots API returns all plots with status.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.get("/api/plots")

        assert response.status_code == 200
        payload = response.get_json()
        assert len(payload) == 4
        assert payload[0]["plotId"] == 1
        assert payload[0]["plotTag"] == "A1"
        assert payload[0]["locationDescription"] == "Northwest Corner"
        assert payload[0]["sizeSqFt"] == 100
        assert payload[0]["isRaisedBed"] is False
        assert payload[0]["status"] == "Active"

    def test_retrieve_available_plots_api(self, client_with_fakes):
        """Confirm the retrieve_available_plots API returns only available plots with status.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.get("/api/plots")

        assert response.status_code == 200
        payload = response.get_json()
        assert len(payload) == 2
        assert payload[0]["plotId"] == 1
        assert payload[0]["plotTag"] == "A1"
        assert payload[0]["status"] == "Active"

    def test_update_plot_status_success(self, client_with_fakes):
        """Confirm updating plot status returns success response.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.put("/api/plots/1/status", json={"status": 2})

        assert response.status_code == 200
        payload = response.get_json()
        assert payload["message"] == "Plot 1 status successfully updated to maintenance."
        assert payload["newStatus"] == "Maintenance"

    def test_update_plot_status_success_active(self, client_with_fakes):
        """Confirm updating plot status to active returns correct response.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.put("/api/plots/1/status", json={"status": 1})

        assert response.status_code == 200
        payload = response.get_json()
        assert payload["message"] == "Plot 1 status successfully updated to active."
        assert payload["newStatus"] == "Active"

    def test_update_plot_status_invalid_plot_id(self, client_with_fakes):
        """Confirm updating status for invalid plot returns error.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.put("/api/plots/99/status", json={"status": 1})

        assert response.status_code == 500
        payload = response.get_json()
        assert "Failed to update status for plot 99" in payload["error"]

    def test_update_plot_status_missing_json(self, client_with_fakes):
        """Confirm updating status without JSON body returns error.

        Args:
            client_with_fakes: Flask test client configured with fake repositories.

        Returns:
            None: Assertions validate behavior.
        """
        response = client_with_fakes.put("/api/plots/1/status")

        assert response.status_code == 500
        payload = response.get_json()
        assert "Failed to update status for plot 1" in payload["error"]
