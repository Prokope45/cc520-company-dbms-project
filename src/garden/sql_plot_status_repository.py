"""SQL Server implementation of the IPlotRepository interface.

Author: Jared Paubel
"""
from typing import Optional

from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.i_plot_status_repository import i_plot_status_repository
from src.garden.models.plot_status import PlotStatus


class SqlPlotStatusRepository(i_plot_status_repository):
    """Persist and retrieve plot status through SQL Server stored procedures."""

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

    def get_status_by_id(self, status_id: int) -> Optional[PlotStatus]:
        """Get plot status by its ID."""
        if status_id is None or status_id <= 0:
           raise ValueError("Status ID must be positive integers.")

        sp_name = "Garden.GetPlotStatusById"
        inp_param_names = ["StatusId"]
        inp_param_values = [status_id]

        rows = self.executor.execute_stored_procedure(
            sp_name,
            input_param_names=inp_param_names,
            input_param_values=inp_param_values
        )

        if len(rows) >= 1:
            row = rows[0]
            return PlotStatus(
                status_id=row["StatusId"],
                status_name=row["StatusName"]
            )
        else:
            return None
