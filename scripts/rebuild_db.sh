#!/bin/sh

# Authors:
#	- Jared Paubel
# 	- Roo agent - local qwen/qwen3.5-9b
# Percentage written by Agent: 95%

# Database rebuild script - wraps src/company/db/main.go
# Usage: ./scripts/rebuild_db.sh [operation]
#
# Operations:
#   rebuild   Full rebuild: clear schema, tables, and seed data (default)
#   clear     Drop all tables and schema
#   schema    Create the Org schema only
#   tables    Create all tables (schema must exist)
#   procedures Create all stored procedures
#   seed      Re-seed data with sample data (creates schema + tables if needed)
#
# Examples:
#   ./scripts/rebuild_db.sh          # Full rebuild (default)
#   ./scripts/rebuild_db.sh seed     # Seed data only
#   ./scripts/rebuild_db.sh clear    # Clear database

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

OPERATION="${1:-rebuild}"

echo "=== Database Operation: $OPERATION ==="

cd "$PROJECT_ROOT"
go run src/company/db/main.go "$OPERATION"
