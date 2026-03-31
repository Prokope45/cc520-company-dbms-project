# src.garden.sql_plot_assignment_repository

Generated from Python docstrings using `pydoc`.

```text
module src.garden.sql_plot_assignment_repository in src.garden

NAME
    src.garden.sql_plot_assignment_repository - SQL Server implementation of the IPlotAssignmentRepository interface.

DESCRIPTION
    Author: Josh Weese

CLASSES
    src.garden.i_plot_assignment_repository.IPlotAssignmentRepository(abc.ABC)
        SqlPlotAssignmentRepository

    class SqlPlotAssignmentRepository(src.garden.i_plot_assignment_repository.IPlotAssignmentRepository)
     |  SqlPlotAssignmentRepository(
     |      server: str = '',
     |      database: str = '',
     |      trusted: bool = True
     |  )
     |
     |  Persist and retrieve plot assignments through SQL Server stored procedures.
     |
     |  Method resolution order:
     |      SqlPlotAssignmentRepository
     |      src.garden.i_plot_assignment_repository.IPlotAssignmentRepository
     |      abc.ABC
     |      builtins.object
     |
     |  Methods defined here:
     |
     |  __init__(self, server: str = '', database: str = '', trusted: bool = True)
     |      Initialize the repository with a SQL command executor.
     |
     |      Args:
     |          server: Optional SQL Server host override.
     |          database: Optional database name override.
     |          trusted: Whether trusted authentication should be used.
     |
     |      Returns:
     |          None: Constructor initializes repository state.
     |
     |  get_assignments(self, gardener_id: int) -> Optional[List[PlotAssignment]]
     |      Return all assignments associated with the supplied gardener identifier.
     |
     |      Args:
     |          gardener_id: Gardener identifier.
     |
     |      Returns:
     |          Optional[List[PlotAssignment]]: Assignment models when rows exist; otherwise ``None``.
     |
     |  save_assignment(
     |      self,
     |      plot_id: int,
     |      gardener_id: int,
     |      start_date: str,
     |      end_date: str | None,
     |      notes: str | None
     |  )
     |      Validate and save a plot assignment for the supplied gardener.
     |
     |      Args:
     |          plot_id: Plot identifier.
     |          gardener_id: Gardener identifier.
     |          start_date: Assignment start date.
     |          end_date: Optional assignment end date.
     |          notes: Optional assignment notes.
     |
     |      Returns:
     |          None: This method performs persistence side effects only.
     |
     |  ----------------------------------------------------------------------
     |  Static methods defined here:
     |
     |  translate_assignments(rows: List[Any]) -> List[PlotAssignment]
     |      Map database rows into ``PlotAssignment`` models.
     |
     |      Args:
     |          rows: Database row objects with assignment fields.
     |
     |      Returns:
     |          List[PlotAssignment]: Translated assignment models.
     |
     |  ----------------------------------------------------------------------
     |  Data and other attributes defined here:
     |
     |  __abstractmethods__ = frozenset()
     |
     |  ----------------------------------------------------------------------
     |  Data descriptors inherited from src.garden.i_plot_assignment_repository.IPlotAssignmentRepository:
     |
     |  __dict__
     |      dictionary for instance variables
     |
     |  __weakref__
     |      list of weak references to the object

DATA
    List = typing.List
        A generic version of list.

    Optional = typing.Optional
        Optional[X] is equivalent to Union[X, None].

FILE
    /workspaces/database-repository-pattern-example/src/garden/sql_plot_assignment_repository.py
```
