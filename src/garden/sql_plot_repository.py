"""SQL Server implementation of the IPlotRepository interface.

Author: Jared Paubel
"""
from typing import List

from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.i_plot_repository import IPlotRepository
from src.garden.models.plot import Plot


class SqlPlotRepository(IPlotRepository):
    """Persist and retrieve plot data through SQL Server stored procedures."""

    def __init__(self, server: str = "", database: str = "", trusted: bool = True):
        """Initialize the repository with a SQL command executor.

        Args:
            server: Optional SQL Server host override.
            database: Optional database name override.
            trusted: Whether trusted authentication should be used.

        Returns:
            None: Constructor initializes repository state.
        """
        self.executor = SqlCommandExecutor(server=server, database=database, trusted=trusted)

    def get_all_plots(self) -> List[Plot]:
        """Return all plots from the database."""
        sp_name = "Garden.GetAllPlots"
        rows = self.executor.execute_stored_procedure(sp_name)

        plots = []
        for row in rows:
            plots.append(Plot(
                plot_id=row["PlotId"],
                plot_tag=row["PlotTag"],
                location_description=row["LocationDescription"],
                size_sq_ft=row["SizeSqFt"],
                is_raised_bed=row["IsRaisedBed"],
                status_id=row["StatusId"],
            ))
        return plots

    def get_available_plots(self) -> List[Plot]:
        """Return available plots from the database."""
        sp_name = "Garden.GetAvailablePlots"
        rows = self.executor.execute_stored_procedure(sp_name)

        plots = []
        for row in rows:
            plots.append(Plot(
                plot_id=row["PlotId"],
                plot_tag=row["PlotTag"],
                location_description=row["LocationDescription"],
                size_sq_ft=row["SizeSqFt"],
                is_raised_bed=row["IsRaisedBed"],
                status_id=row["StatusId"],
            ))
        return plots

    def update_plot_status(
        self,
        plot_id: int,
        new_status_id: int
    ) -> bool:
        """Update status of a specific plot."""
        if (
            (plot_id is None or plot_id <= 0) or
            (new_status_id is None or new_status_id <= 0)
        ):
           raise ValueError("Plot ID and Status ID must be positive integers.")

        sp_name = "Garden.UpdatePlotStatus"
        inp_param_names = ["PlotId", "NewStatusId"]
        inp_param_values = [plot_id, new_status_id]

        rows = self.executor.execute_stored_procedure(
            sp_name,
            input_param_names=inp_param_names,
            input_param_values=inp_param_values
        )

        # Return true if execution was successful; false otherwise
        return rows is not None and rows > 0
