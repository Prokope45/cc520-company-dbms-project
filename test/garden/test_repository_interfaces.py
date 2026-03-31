"""Tests that execute abstract interface method bodies for coverage.

Author: Josh Weese   
"""

from src.garden.i_plot_assignment_repository import IPlotAssignmentRepository
from src.garden.i_gardener_repository import IGardenerRepository


class ConcretePlotAssignmentRepository(IPlotAssignmentRepository):
    """Concrete test double used to exercise abstract assignment repository methods."""

    def get_assignments(self, gardener_id: int):
        """Delegate to the abstract base implementation for coverage.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Optional[List[PlotAssignment]]: Result from the abstract base implementation.
        """
        return super().get_assignments(gardener_id)

    def save_assignment(self, plot_id: int, gardener_id: int, start_date: str,
                        end_date: str | None, notes: str | None):
        """Delegate to the abstract base implementation for coverage.

        Args:
            plot_id: Plot identifier.
            gardener_id: Gardener identifier.
            start_date: Assignment start date.
            end_date: Optional assignment end date.
            notes: Optional assignment notes.

        Returns:
            None: Result from the abstract base implementation.
        """
        return super().save_assignment(plot_id, gardener_id, start_date, end_date, notes)


class ConcreteGardenerRepository(IGardenerRepository):
    """Concrete test double used to exercise abstract gardener repository methods."""

    def create_gardener(self, first_name: str, last_name: str, phone: str | None,
                        email: str | None, join_date: str | None):
        """Delegate to the abstract base implementation for coverage.

        Args:
            first_name: Gardener first name.
            last_name: Gardener last name.
            phone: Optional phone number.
            email: Optional email address.
            join_date: Optional join date.

        Returns:
            Optional[Gardener]: Result from the abstract base implementation.
        """
        return super().create_gardener(first_name, last_name, phone, email, join_date)

    def retrieve_gardeners(self):
        """Delegate to the abstract base implementation for coverage.

        Returns:
            Optional[List[Gardener]]: Result from the abstract base implementation.
        """
        return super().retrieve_gardeners()

    def fetch_gardener(self, gardener_id: int):
        """Delegate to the abstract base implementation for coverage.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Optional[Gardener]: Result from the abstract base implementation.
        """
        return super().fetch_gardener(gardener_id)

    def get_gardener_by_email(self, email: str):
        """Delegate to the abstract base implementation for coverage.

        Args:
            email: Email address to search by.

        Returns:
            Optional[Gardener]: Result from the abstract base implementation.
        """
        return super().get_gardener_by_email(email)


def test_assignment_interface_method_bodies_are_executable():
    """Ensure assignment repository abstract method bodies remain executable for coverage.

    Returns:
        None: Assertions validate behavior.
    """
    repository = ConcretePlotAssignmentRepository()

    assert repository.get_assignments(1) is None
    assert repository.save_assignment(2, 1, "2024-04-01", None, "Spring crops") is None


def test_gardener_interface_method_bodies_are_executable():
    """Ensure gardener repository abstract method bodies remain executable for coverage.

    Returns:
        None: Assertions validate behavior.
    """
    repository = ConcreteGardenerRepository()

    assert repository.create_gardener("Chili", "Heeler", "555-0101", "chili@test.com", "2024-03-01") is None
    assert repository.retrieve_gardeners() is None
    assert repository.fetch_gardener(1) is None
    assert repository.get_gardener_by_email("chili@test.com") is None
