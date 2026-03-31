# src.garden.i_plot_assignment_repository

Generated from Python docstrings using `pydoc`.

```text
module src.garden.i_plot_assignment_repository in src.garden

NAME
    src.garden.i_plot_assignment_repository - Repository interface for plot-assignment-related data operations.

DESCRIPTION
    Author: Josh Weese

CLASSES
    abc.ABC(builtins.object)
        IPlotAssignmentRepository
    
    class IPlotAssignmentRepository(abc.ABC)
     |  Define the contract for plot assignment repository implementations.
     |  
     |  Method resolution order:
     |      IPlotAssignmentRepository
     |      abc.ABC
     |      builtins.object
     |  
     |  Methods defined here:
     |  
     |  get_assignments(self, gardener_id: int) -> Optional[List[src.garden.models.plot_assignment.PlotAssignment]]
     |      Return assignments associated with the supplied gardener identifier.
     |      
     |      Args:
     |          gardener_id: Gardener identifier.
     |      
     |      Returns:
     |          Optional[List[PlotAssignment]]: Matching assignment models, or ``None`` when absent.
     |  
     |  save_assignment(self, plot_id: int, gardener_id: int, start_date: str, end_date: str | None, notes: str | None)
     |      Persist a plot assignment for a gardener.
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
     |  Data descriptors defined here:
     |  
     |  __dict__
     |      dictionary for instance variables (if defined)
     |  
     |  __weakref__
     |      list of weak references to the object (if defined)
     |  
     |  ----------------------------------------------------------------------
     |  Data and other attributes defined here:
     |  
     |  __abstractmethods__ = frozenset({'get_assignments', 'save_assignment'}...

DATA
    List = typing.List
        A generic version of list.
    
    Optional = typing.Optional
        Optional type.
        
        Optional[X] is equivalent to Union[X, None].

FILE
    /home/runner/work/database-repository-pattern-example/database-repository-pattern-example/src/garden/i_plot_assignment_repository.py
```
