"""Repository interface for plot data operations.

Author: Jared Paubel

"""
from abc import ABC, abstractmethod
from typing import Optional

from src.garden.models.plot_status import PlotStatus


class IPlotRepository(ABC):
    """Define the contract for plot status repository implementations."""

    @abstractmethod
    def get_status_by_id(self, status_id: int) -> Optional[PlotStatus]:
        """Retrieve plot status record by ID."""
        pass
