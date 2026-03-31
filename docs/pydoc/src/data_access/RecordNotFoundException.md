# src.data_access.RecordNotFoundException

Generated from Python docstrings using `pydoc`.

```text
module src.data_access.RecordNotFoundException in src.data_access

NAME
    src.data_access.RecordNotFoundException

DESCRIPTION
    Custom exception for handling cases where a requested record is not found in the database.
    Author: Josh Weese

CLASSES
    builtins.Exception(builtins.BaseException)
        RecordNotFoundException

    class RecordNotFoundException(builtins.Exception)
     |  RecordNotFoundException(key)
     |
     |  Represent an error raised when a requested database record does not exist.
     |
     |  Method resolution order:
     |      RecordNotFoundException
     |      builtins.Exception
     |      builtins.BaseException
     |      builtins.object
     |
     |  Methods defined here:
     |
     |  __init__(self, key)
     |      Initialize the exception with the missing record key.
     |
     |      Args:
     |          key: Identifier value used to look up the missing record.
     |
     |      Returns:
     |          None: Constructor initializes exception state.
     |
     |  ----------------------------------------------------------------------
     |  Data descriptors defined here:
     |
     |  __weakref__
     |      list of weak references to the object
     |
     |  ----------------------------------------------------------------------
     |  Static methods inherited from builtins.Exception:
     |
     |  __new__(*args, **kwargs) class method of builtins.Exception
     |      Create and return a new object.  See help(type) for accurate signature.
     |
     |  ----------------------------------------------------------------------
     |  Methods inherited from builtins.BaseException:
     |
     |  __reduce__(self, /)
     |      Helper for pickle.
     |
     |  __repr__(self, /)
     |      Return repr(self).
     |
     |  __setstate__(self, state, /)
     |
     |  __str__(self, /)
     |      Return str(self).
     |
     |  add_note(self, note, /)
     |      Add a note to the exception
     |
     |  with_traceback(self, tb, /)
     |      Set self.__traceback__ to tb and return self.
     |
     |  ----------------------------------------------------------------------
     |  Data descriptors inherited from builtins.BaseException:
     |
     |  __cause__
     |
     |  __context__
     |
     |  __dict__
     |
     |  __suppress_context__
     |
     |  __traceback__
     |
     |  args

FILE
    /workspaces/database-repository-pattern-example/src/data_access/RecordNotFoundException.py
```
