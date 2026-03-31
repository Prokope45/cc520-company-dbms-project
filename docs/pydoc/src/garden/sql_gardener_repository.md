# src.garden.sql_gardener_repository

Generated from Python docstrings using `pydoc`.

```text
module src.garden.sql_gardener_repository in src.garden

NAME
    src.garden.sql_gardener_repository - SQL Server implementation of the IGardenerRepository interface.

DESCRIPTION
    Author: Josh Weese

CLASSES
    src.garden.i_gardener_repository.IGardenerRepository(abc.ABC)
        SqlGardenerRepository
    
    class SqlGardenerRepository(src.garden.i_gardener_repository.IGardenerRepository)
     |  SqlGardenerRepository(server: str = '', database: str = '', trusted: bool = True)
     |  
     |  Persist and retrieve gardeners through SQL Server stored procedures.
     |  
     |  Method resolution order:
     |      SqlGardenerRepository
     |      src.garden.i_gardener_repository.IGardenerRepository
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
     |  create_gardener(self, first_name: str, last_name: str, phone: str | None, email: str | None, join_date: str | None) -> Optional[src.garden.models.gardener.Gardener]
     |      Create a new gardener record and return the resulting model.
     |      
     |      Args:
     |          first_name: Gardener first name.
     |          last_name: Gardener last name.
     |          phone: Optional phone number.
     |          email: Optional email address.
     |          join_date: Optional join date.
     |      
     |      Returns:
     |          Optional[Gardener]: Created gardener model, or ``None`` when no output row is returned.
     |  
     |  fetch_gardener(self, gardener_id: int) -> Optional[src.garden.models.gardener.Gardener]
     |      Return one gardener by identifier or raise when the record is absent.
     |      
     |      Args:
     |          gardener_id: Gardener identifier.
     |      
     |      Returns:
     |          Optional[Gardener]: Matching gardener model.
     |  
     |  get_gardener_by_email(self, email: str) -> Optional[src.garden.models.gardener.Gardener]
     |      Return one gardener matching the supplied email address.
     |      
     |      Args:
     |          email: Email address to search by.
     |      
     |      Returns:
     |          Optional[Gardener]: Matching gardener model, or ``None`` when absent.
     |  
     |  retrieve_gardeners(self) -> Optional[List[src.garden.models.gardener.Gardener]]
     |      Return all gardeners currently stored in the backing database.
     |      
     |      Returns:
     |          Optional[List[Gardener]]: Gardener models when rows exist; otherwise ``None``.
     |  
     |  translate_gardener(self, row: Any) -> src.garden.models.gardener.Gardener
     |      Map a database row object into a ``Gardener`` model.
     |      
     |      Args:
     |          row: Database row object with gardener fields.
     |      
     |      Returns:
     |          Gardener: Translated gardener model.
     |  
     |  translate_gardeners(self, rows: List[Any]) -> List[src.garden.models.gardener.Gardener]
     |      Map multiple database rows into a list of ``Gardener`` models.
     |      
     |      Args:
     |          rows: Database row objects with gardener fields.
     |      
     |      Returns:
     |          List[Gardener]: Translated gardener models.
     |  
     |  ----------------------------------------------------------------------
     |  Data and other attributes defined here:
     |  
     |  __abstractmethods__ = frozenset()
     |  
     |  ----------------------------------------------------------------------
     |  Data descriptors inherited from src.garden.i_gardener_repository.IGardenerRepository:
     |  
     |  __dict__
     |      dictionary for instance variables (if defined)
     |  
     |  __weakref__
     |      list of weak references to the object (if defined)

DATA
    Any = typing.Any
        Special type indicating an unconstrained type.
        
        - Any is compatible with every type.
        - Any assumed to have all methods.
        - All values assumed to be instances of Any.
        
        Note that all the above statements are true from the point of view of
        static type checkers. At runtime, Any should not be used with instance
        or class checks.
    
    List = typing.List
        A generic version of list.
    
    Optional = typing.Optional
        Optional type.
        
        Optional[X] is equivalent to Union[X, None].

FILE
    /home/runner/work/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/src/garden/sql_gardener_repository.py
```
