# src.db.rebuild

Generated from Python docstrings using `pydoc`.

```text
module src.db.rebuild in src.db

NAME
    src.db.rebuild - Database rebuild/seed command runner.

DESCRIPTION
    Author: Josh Weese

FUNCTIONS
    run_from_cli(argv: list[str] | None = None) -> int
        Parse command-line arguments and run the requested database action.

        Args:
            argv: Optional argument vector to parse instead of ``sys.argv``.

        Returns:
            int: Process-style exit code.

    run_rebuild(
        *,
        server: str = '',
        database: str = '',
        user: str = '',
        password: str = '',
        trusted: bool = True
    ) -> None
        Run schema + table + procedure + seed scripts in order.

        Args:
            server: Optional SQL Server host override.
            database: Optional database name override.
            user: Optional SQL login username.
            password: Optional SQL login password.
            trusted: Whether trusted authentication should be used when credentials are absent.

        Returns:
            None: This function performs database side effects only.

    run_seed_only(
        *,
        server: str = '',
        database: str = '',
        user: str = '',
        password: str = '',
        trusted: bool = True
    ) -> None
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
    DATA_FILES = [PosixPath('/workspaces/database-repository-patte...ample...
    GARDEN_SCHEMA = 'Garden'
    GO_LINE = re.compile('^\\s*GO(?:\\s+\\d+)?\\s*$', re.IGNORECASE)
    Iterable = typing.Iterable
        A generic version of collections.abc.Iterable.

    PROCEDURE_FILES = [PosixPath('/workspaces/database-repository-patte......
    ROOT = PosixPath('/workspaces/database-repository-pattern-example')
    SCHEMA_FILES = [PosixPath('/workspaces/database-repository-pattern-exa...
    SQL_ROOT = PosixPath('/workspaces/database-repository-pattern-example/...
    TABLE_FILES = [PosixPath('/workspaces/database-repository-patte...ple/...

FILE
    /workspaces/database-repository-pattern-example/src/db/rebuild.py
```
