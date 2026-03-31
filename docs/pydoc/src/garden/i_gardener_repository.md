# src.garden.i_gardener_repository

Generated from Python docstrings using `pydoc`.

```text
module src.garden.i_gardener_repository in src.garden

NAME
    src.garden.i_gardener_repository - Repository interface for gardener-related data operations.

DESCRIPTION
    Author: Josh Weese

CLASSES
    abc.ABC(builtins.object)
        IGardenerRepository
    
    class IGardenerRepository(abc.ABC)
     |  Define the contract for gardener repository implementations.
     |  
     |  Method resolution order:
     |      IGardenerRepository
     |      abc.ABC
     |      builtins.object
     |  
     |  Methods defined here:
     |  
     |  create_gardener(self, first_name: str, last_name: str, phone: str | None, email: str | None, join_date: str | None) -> Optional[src.garden.models.gardener.Gardener]
     |      Create a gardener record and return the created model when successful.
     |      
     |      Args:
     |          first_name: Gardener first name.
     |          last_name: Gardener last name.
     |          phone: Optional phone number.
     |          email: Optional email address.
     |          join_date: Optional join date.
     |      
     |      Returns:
     |          Optional[Gardener]: Created gardener model, or ``None`` when creation fails.
     |  
     |  fetch_gardener(self, gardener_id: int) -> Optional[src.garden.models.gardener.Gardener]
     |      Return a single gardener by identifier.
     |      
     |      Args:
     |          gardener_id: Gardener identifier.
     |      
     |      Returns:
     |          Optional[Gardener]: Matching gardener model.
     |  
     |  get_gardener_by_email(self, email: str) -> Optional[src.garden.models.gardener.Gardener]
     |      Return a gardener matching the supplied email address.
     |      
     |      Args:
     |          email: Email address to search by.
     |      
     |      Returns:
     |          Optional[Gardener]: Matching gardener model, or ``None`` when absent.
     |  
     |  retrieve_gardeners(self) -> Optional[List[src.garden.models.gardener.Gardener]]
     |      Return all known gardeners or ``None`` when no rows exist.
     |      
     |      Returns:
     |          Optional[List[Gardener]]: Gardener models when rows exist; otherwise ``None``.
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
     |  __abstractmethods__ = frozenset({'create_gardener', 'fetch_gardener', ...

DATA
    List = typing.List
        A generic version of list.
    
    Optional = typing.Optional
        Optional type.
        
        Optional[X] is equivalent to Union[X, None].

FILE
    /home/runner/work/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/src/garden/i_gardener_repository.py
```
