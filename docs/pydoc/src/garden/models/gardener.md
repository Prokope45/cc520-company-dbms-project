# src.garden.models.gardener

Generated from Python docstrings using `pydoc`.

```text
module src.garden.models.gardener in src.garden.models

NAME
    src.garden.models.gardener

DESCRIPTION
    Data model for a gardener.
    Author: Josh Weese

CLASSES
    builtins.object
        Gardener
    
    class Gardener(builtins.object)
     |  Gardener(gardener_id: int, first_name: str, last_name: str, phone: str | None, email: str | None, join_date: str | None)
     |  
     |  Represent a gardener record returned by the repository layer.
     |  
     |  Methods defined here:
     |  
     |  __init__(self, gardener_id: int, first_name: str, last_name: str, phone: str | None, email: str | None, join_date: str | None)
     |      Initialize a gardener model with contact and membership details.
     |      
     |      Args:
     |          gardener_id: Unique gardener identifier.
     |          first_name: Gardener first name.
     |          last_name: Gardener last name.
     |          phone: Optional phone number.
     |          email: Optional email address.
     |          join_date: Optional join date.
     |      
     |      Returns:
     |          None: Constructor initializes instance state.
     |  
     |  ----------------------------------------------------------------------
     |  Readonly properties defined here:
     |  
     |  email
     |      Return the gardener's email address.
     |      
     |      Returns:
     |          str | None: Email address value when present.
     |  
     |  first_name
     |      Return the gardener's first name.
     |      
     |      Returns:
     |          str: First name value.
     |  
     |  gardener_id
     |      Return the unique identifier for the gardener.
     |      
     |      Returns:
     |          int: Gardener identifier.
     |  
     |  join_date
     |      Return the gardener's join date.
     |      
     |      Returns:
     |          str | None: Join date when present.
     |  
     |  last_name
     |      Return the gardener's last name.
     |      
     |      Returns:
     |          str: Last name value.
     |  
     |  phone
     |      Return the gardener's phone number.
     |      
     |      Returns:
     |          str | None: Phone number when present.
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
    /home/runner/work/database-repository-pattern-example/database-repository-pattern-example/src/garden/models/gardener.py
```
