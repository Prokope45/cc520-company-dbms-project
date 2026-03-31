"""Run project commands.

Usage:
    python -m src
    python -m src rebuild-db [rebuild|seed-only] [--server ...] [--database ...]
"""

import sys

from src.db.rebuild import run_from_cli
from src.web.app import create_app


app = create_app()


if __name__ == "__main__":
    if len(sys.argv) > 1 and sys.argv[1] == "rebuild-db":
        raise SystemExit(run_from_cli(sys.argv[2:]))
    app.run(debug=True)
