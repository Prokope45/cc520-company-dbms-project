"""SQL Server implementation of the IPlotRepository interface.

Author: Jared Paubel
"""
from typing import List, Optional

from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.i_plot_repository import IPlotRepository
from src.garden.models.plot import Plot
from src.garden.models.plot_status import PlotStatus


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

    def __translate_plot_status(self, row: Any) -> Plot:
        """Map a database row object into a ``PlotStatus`` model.

        Args:
            row: Database row object with plot status fields.

        Returns:
            PlotStatus: Translated plot status model.
        """
        return PlotStatus(row.StatusId, row.StatusName)

    def __translate_plot(self, row: Any) -> Plot:
        """Map a database row object into a ``Plot`` model.

        Args:
            row: Database row object with plot fields.

        Returns:
            Plot: Translated plot model.
        """
        return Plot(row.PlotId, row.PlotTag, row.LocationDescription, row.SizeSqFt, row.IsRaisedBed, row.StatusId)

    def __translate_plots(self, rows: List[Any]) -> List[Plot]:
        """Map multiple database rows into a list of ``Plot`` models.

        Args:
            rows: Database row objects with plot fields.

        Returns:
            List[Plot]: Translated plot models.
        """
        plots = []
        for row in rows:
            plots.append(self.__translate_plot(row))
        return plots

    def retrieve_plots(self) -> List[Plot]:
        """Return all plots from the database."""
        sp_name = "Garden.RetrievePlots"
        rows = self.executor.execute_stored_procedure(sp_name)

        if len(rows) >= 1:
            return self.__translate_plots(rows)
        return None

    def retrieve_available_plots(self) -> List[Plot]:
        """Return available plots from the database."""
        sp_name = "Garden.RetrieveAvailablePlots"
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
        try:
            self.executor.execute_stored_procedure(
                sp_name,
                input_param_names=inp_param_names,
                input_param_values=inp_param_values
            )
        except Exception:
            return False
        return True

    def retrieve_plot_status(self, plot_id: int) -> Optional[PlotStatus]:
        sp_name = "Garden.RetrievePlotStatus"
        inp_param_names = ['PlotId']
        inp_param_values = [plot_id]

        rows = self.executor.execute_stored_procedure(
            sp_name,
            input_param_names=inp_param_names,
            input_param_values=inp_param_values
        )

        if len(rows) == 1:
            return self.__translate_plot_status(rows[0])
        else:
            return None
