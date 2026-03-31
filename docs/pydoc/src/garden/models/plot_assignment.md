# src.garden.models.plot_assignment

Generated from Python docstrings using `pydoc`.

```text
module src.garden.models.plot_assignment in src.garden.models

NAME
    src.garden.models.plot_assignment - Data model for garden plot assignments.

DESCRIPTION
    Author: Josh Weese

CLASSES
    builtins.object
        PlotAssignment
    
    class PlotAssignment(builtins.object)
     |  PlotAssignment(assignment_id: int, plot_id: int, gardener_id: int, start_date: str, end_date: str | None, notes: str | None, plot_tag: str | None, location_description: str | None, size_sq_ft: int | None, is_raised_bed: bool | None, status: str | None)
     |  
     |  Represent a gardener's assignment to a plot.
     |  
     |  Methods defined here:
     |  
     |  __init__(self, assignment_id: int, plot_id: int, gardener_id: int, start_date: str, end_date: str | None, notes: str | None, plot_tag: str | None, location_description: str | None, size_sq_ft: int | None, is_raised_bed: bool | None, status: str | None)
     |      Initialize a plot assignment model with assignment and plot details.
     |      
     |      Args:
     |          assignment_id: Unique assignment identifier.
     |          plot_id: Plot identifier.
     |          gardener_id: Gardener identifier.
     |          start_date: Assignment start date.
     |          end_date: Optional assignment end date.
     |          notes: Optional assignment notes.
     |          plot_tag: Optional plot tag (e.g., A1, B2).
     |          location_description: Optional plot location description.
     |          size_sq_ft: Optional plot size in square feet.
     |          is_raised_bed: Optional raised-bed flag.
     |          status: Optional plot status.
     |      
     |      Returns:
     |          None: Constructor initializes instance state.
     |  
     |  ----------------------------------------------------------------------
     |  Readonly properties defined here:
     |  
     |  assignment_id
     |      Return the assignment identifier.
     |      
     |      Returns:
     |          int: Assignment identifier.
     |  
     |  end_date
     |      Return the optional assignment end date.
     |      
     |      Returns:
     |          str | None: End date when present.
     |  
     |  gardener_id
     |      Return the gardener identifier.
     |      
     |      Returns:
     |          int: Gardener identifier.
     |  
     |  is_raised_bed
     |      Return whether the plot is a raised bed.
     |      
     |      Returns:
     |          bool | None: Raised-bed flag when present.
     |  
     |  location_description
     |      Return the optional plot location description.
     |      
     |      Returns:
     |          str | None: Plot location description when present.
     |  
     |  notes
     |      Return optional notes for the assignment.
     |      
     |      Returns:
     |          str | None: Assignment notes when present.
     |  
     |  plot_id
     |      Return the plot identifier.
     |      
     |      Returns:
     |          int: Plot identifier.
     |  
     |  plot_tag
     |      Return the optional plot tag.
     |      
     |      Returns:
     |          str | None: Plot tag (e.g., A1, B2) when present.
     |  
     |  size_sq_ft
     |      Return the optional plot size in square feet.
     |      
     |      Returns:
     |          int | None: Plot size when present.
     |  
     |  start_date
     |      Return the assignment start date.
     |      
     |      Returns:
     |          str: Start date value.
     |  
     |  status
     |      Return the optional plot status.
     |      
     |      Returns:
     |          str | None: Plot status when present.
     |  
     |  ----------------------------------------------------------------------
     |  Data descriptors defined here:
     |  
     |  __dict__
     |      dictionary for instance variables (if defined)
     |  
     |  __weakref__
     |      list of weak references to the object (if defined)

FILE
    /home/runner/work/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/src/garden/models/plot_assignment.py
```
