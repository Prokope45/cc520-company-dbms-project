"""Repository interface for plot-assignment-related data operations.

Author: Josh Weese

"""
from abc import ABC, abstractmethod
from typing import Optional, List

from src.garden.models.plot_assignment import PlotAssignment


class IPlotAssignmentRepository(ABC):
    """Define the contract for plot assignment repository implementations."""

    @abstractmethod
    def get_assignments(self, gardener_id: int) -> Optional[List[PlotAssignment]]:
        """Return assignments associated with the supplied gardener identifier.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Optional[List[PlotAssignment]]: Matching assignment models, or ``None`` when absent.
        """
        pass

    @abstractmethod
    def save_assignment(self, plot_id: int, gardener_id: int, start_date: str,
                        end_date: str | None, notes: str | None):
        """Persist a plot assignment for a gardener.

        Args:
            plot_id: Plot identifier.
            gardener_id: Gardener identifier.
            start_date: Assignment start date.
            end_date: Optional assignment end date.
            notes: Optional assignment notes.

        Returns:
            None: This method performs persistence side effects only.
        """
        pass

