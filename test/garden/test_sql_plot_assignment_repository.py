"""Unit tests for SqlPlotAssignmentRepository without external DB dependencies.

Author: Josh Weese

"""

from types import SimpleNamespace

import pytest

from src.data_access.sql_command_executor import SqlCommandExecutor
from src.garden.sql_plot_assignment_repository import SqlPlotAssignmentRepository


class FakeExecutor(SqlCommandExecutor):
       """Return canned stored procedure results for assignment repository tests."""

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
def assignment_repo():
       """Return an assignment repository wired to a fake executor.

       Returns:
              SqlPlotAssignmentRepository: Repository instance with fake executor injected.
       """
       repo = SqlPlotAssignmentRepository()
       repo.executor = FakeExecutor()
       return repo


class TestSqlPlotAssignmentRepository:
       """Verify the SQL assignment repository maps and validates data correctly."""

       def test_save_assignment(self, assignment_repo):
              """Ensure save operations pass the expected parameters to the executor.

              Args:
                     assignment_repo: Repository fixture under test.

              Returns:
                     None: Assertions validate behavior.
              """
              assignment_repo.save_assignment(2, 1, "2024-04-01", None, "Spring crops")

              procedure_name, _, values, _ = assignment_repo.executor.calls[-1]
              assert procedure_name == "Garden.SavePlotAssignment"
              assert values == [2, 1, "2024-04-01", None, "Spring crops"]

       def test_retrieve_assignments(self, assignment_repo):
              """Ensure assignment rows are translated into models.

              Args:
                     assignment_repo: Repository fixture under test.

              Returns:
                     None: Assertions validate behavior.
              """
              assignment_repo.executor.results_by_sp["Garden.RetrieveAssignmentsForGardener"] = [
                     SimpleNamespace(AssignmentId=1, PlotId=2, GardenerId=1, StartDate="2024-04-01", EndDate=None,
                                     Notes="Spring crops", PlotTag="A2", LocationDescription="Northwest Corner - A2", SizeSqFt=100,
                                     IsRaisedBed=False, Status="Active"),
                     SimpleNamespace(AssignmentId=2, PlotId=3, GardenerId=1, StartDate="2024-05-01", EndDate="2024-10-15",
                                     Notes="Summer only", PlotTag="B1", LocationDescription="East Fence - B1", SizeSqFt=90,
                                     IsRaisedBed=True, Status="Maintenance"),
              ]

              assignments = assignment_repo.get_assignments(1)

              assert assignments is not None
              assert len(assignments) == 2
              assert assignments[0].plot_id == 2
              assert assignments[1].status == "Maintenance"
              assert assignments[1].end_date == "2024-10-15"

       def test_retrieve_assignments_returns_none_when_no_rows(self, assignment_repo):
              """Ensure no rows produce a ``None`` result.

              Args:
                     assignment_repo: Repository fixture under test.

              Returns:
                     None: Assertions validate behavior.
              """
              assignment_repo.executor.results_by_sp["Garden.RetrieveAssignmentsForGardener"] = []

              assert assignment_repo.get_assignments(1) is None

       @pytest.mark.parametrize(
              "plot_id,gardener_id,start_date,end_date,notes,expected",
              [
                     (0, 1, "2024-04-01", None, "x", "Plot ID must be a positive integer."),
                     (2, 0, "2024-04-01", None, "x", "Gardener ID must be a positive integer."),
                     (2, 1, "", None, "x", "Start date cannot be empty."),
                     (2, 1, None, None, "x", "Start date cannot be empty."),
              ],
       )
       def test_save_assignment_validations(self, assignment_repo, plot_id, gardener_id, start_date, end_date, notes, expected):
              """Ensure invalid assignment data raises the expected validation error.

              Args:
                     assignment_repo: Repository fixture under test.
                     plot_id: Candidate plot identifier.
                     gardener_id: Candidate gardener identifier.
                     start_date: Candidate start date value.
                     end_date: Candidate end date value.
                     notes: Candidate notes value.
                     expected: Expected validation message pattern.

              Returns:
                     None: Assertions validate behavior.
              """
              with pytest.raises(ValueError, match=expected):
                     assignment_repo.save_assignment(plot_id, gardener_id, start_date, end_date, notes)
