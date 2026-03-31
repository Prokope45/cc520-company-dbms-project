"""SQL Server implementation of the IPlotAssignmentRepository interface.

Author: Josh Weese

"""
from typing import List, Optional, Any

from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.i_plot_assignment_repository import IPlotAssignmentRepository
from src.garden.models.plot_assignment import PlotAssignment


class SqlPlotAssignmentRepository(IPlotAssignmentRepository):
    """Persist and retrieve plot assignments through SQL Server stored procedures."""

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

    def get_assignments(self, gardener_id: int) -> Optional[List[PlotAssignment]]:
        """Return all assignments associated with the supplied gardener identifier.

        Args:
            gardener_id: Gardener identifier.

        Returns:
            Optional[List[PlotAssignment]]: Assignment models when rows exist; otherwise ``None``.
        """
        sp_name = 'Garden.RetrieveAssignmentsForGardener'
        inp_param_names = ['GardenerId']
        inp_param_values = [gardener_id]
        rows = self.executor.execute_stored_procedure(sp_name,
                                                      input_param_names=inp_param_names,
                                                      input_param_values=inp_param_values)
        if len(rows) >= 1:
            return self.translate_assignments(rows)
        else:
            return None

    def save_assignment(self, plot_id: int, gardener_id: int, start_date: str,
                        end_date: str | None, notes: str | None):
        """Validate and save a plot assignment for the supplied gardener.

        Args:
            plot_id: Plot identifier.
            gardener_id: Gardener identifier.
            start_date: Assignment start date.
            end_date: Optional assignment end date.
            notes: Optional assignment notes.

        Returns:
            None: This method performs persistence side effects only.
        """
        if plot_id is None or plot_id <= 0:
            raise ValueError("Plot ID must be a positive integer.")
        if gardener_id is None or gardener_id <= 0:
            raise ValueError("Gardener ID must be a positive integer.")
        if start_date is None or start_date == "":
            raise ValueError("Start date cannot be empty.")

        sp_name = 'Garden.SavePlotAssignment'
        inp_param_names = ['PlotId', 'GardenerId', 'StartDate', 'EndDate', 'Notes']
        inp_param_values = [plot_id, gardener_id, start_date, end_date, notes]
        self.executor.execute_stored_procedure(sp_name,
                                               input_param_names=inp_param_names,
                                               input_param_values=inp_param_values)

    @staticmethod
    def translate_assignments(rows: List[Any]) -> List[PlotAssignment]:
        """Map database rows into ``PlotAssignment`` models.

        Args:
            rows: Database row objects with assignment fields.

        Returns:
            List[PlotAssignment]: Translated assignment models.
        """
        assignments = []
        for row in rows:
            assignments.append(
                PlotAssignment(
                    row.AssignmentId,
                    row.PlotId,
                    row.GardenerId,
                    row.StartDate,
                    row.EndDate,
                    row.Notes,
                    row.PlotTag,
                    row.LocationDescription,
                    row.SizeSqFt,
                    row.IsRaisedBed,
                    row.Status,
                )
            )
        return assignments
