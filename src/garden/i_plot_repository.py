"""Repository interface for plot data operations.

Author: Jared Paubel

"""
from abc import ABC, abstractmethod
from typing import List, Optional

from src.garden.models.plot import Plot
from src.garden.models.plot_status import PlotStatus

class IPlotRepository(ABC):
    """Define the contract for plot repository implementations."""

    @abstractmethod
    def retrieve_plots(self) -> List[Plot]:
        """Retrieve all plots from the database."""
        pass

    @abstractmethod
    def retrieve_available_plots(self) -> List[Plot]:
        """Review only plots taht are currently available."""
        pass

    @abstractmethod
    def update_plot_status(
        self,
        plot_id: int,
        new_status_id: int
    ) -> bool:
        """Update the status of a specific plot."""
        pass

    @abstractmethod
    def retrieve_plot_status(
        self,
        plot_id: int
    ) -> Optional[PlotStatus]:
        """Update the status of a specific plot status."""
        pass
