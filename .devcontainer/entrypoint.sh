#!/bin/bash

# Start SQL Server in the background
/opt/mssql/bin/sqlservr &

# Wait for SQL Server to start
sleep 30s

# Run the SQL scripts
/opt/mssql-tools18/bin/sqlcmd -C -S localhost -U sa -P sup3rs3cur3P@ssword -i /home/mssql/sql-init/createCC520.sql
/opt/mssql-tools18/bin/sqlcmd -C -S localhost -U sa -P sup3rs3cur3P@ssword -i /home/mssql/sql-init/worldWideImporters.sql

# Set VS Code to use this venv by default (for devcontainers, this is done via symlink or settings)
mkdir -p /workspaces/.vscode

# Keep the container running
wait