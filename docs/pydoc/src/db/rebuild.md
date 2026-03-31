# src.db.rebuild

Generated from Python docstrings using `pydoc`.

```text
module src.db.rebuild in src.db

NAME
    src.db.rebuild - Database rebuild/seed command runner.

DESCRIPTION
    Author: Josh Weese

FUNCTIONS
    run_from_cli(argv: 'list[str] | None' = None) -> 'int'
        Parse command-line arguments and run the requested database action.
        
        Args:
            argv: Optional argument vector to parse instead of ``sys.argv``.
        
        Returns:
            int: Process-style exit code.
    
    run_rebuild(*, server: 'str' = '', database: 'str' = '', user: 'str' = '', password: 'str' = '', trusted: 'bool' = True) -> 'None'
        Run schema + table + procedure + seed scripts in order.
        
        Args:
            server: Optional SQL Server host override.
            database: Optional database name override.
            user: Optional SQL login username.
            password: Optional SQL login password.
            trusted: Whether trusted authentication should be used when credentials are absent.
        
        Returns:
            None: This function performs database side effects only.
    
    run_seed_only(*, server: 'str' = '', database: 'str' = '', user: 'str' = '', password: 'str' = '', trusted: 'bool' = True) -> 'None'
        Run only data seed scripts.
        
        Args:
            server: Optional SQL Server host override.
            database: Optional database name override.
            user: Optional SQL login username.
            password: Optional SQL login password.
            trusted: Whether trusted authentication should be used when credentials are absent.
        
        Returns:
            None: This function performs database side effects only.

DATA
    DATA_FILES = [PosixPath('/home/runner/work/ksu-cc520-spring-20...ern-e...
    GARDEN_SCHEMA = 'Garden'
    GO_LINE = re.compile('^\\s*GO(?:\\s+\\d+)?\\s*$', re.IGNORECASE)
    Iterable = typing.Iterable
        A generic version of collections.abc.Iterable.
    
    PROCEDURE_FILES = [PosixPath('/home/runner/work/ksu-cc520-spring-20......
    ROOT = PosixPath('/home/runner/work/ksu-cc520-spring-20...munity-garde...
    SCHEMA_FILES = [PosixPath('/home/runner/work/ksu-cc520-spring-20...ory...
    SQL_ROOT = PosixPath('/home/runner/work/ksu-cc520-spring-20...pp-datab...
    TABLE_FILES = [PosixPath('/home/runner/work/ksu-cc520-spring-20...n-e/...
    annotations = _Feature((3, 7, 0, 'beta', 1), (3, 11, 0, 'alpha', 0), 1...

FILE
    /home/runner/work/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/ksu-cc520-spring-2026-classroom-c6fd68-homework-8-community-garden-app-database-repository-pattern-e/src/db/rebuild.py
```
