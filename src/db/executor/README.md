# SQL Query Executor

A general-purpose SQL query executor for executing stored procedures with parameters.

## Overview

This package provides a simple interface to execute stored procedures by name, passing parameters as a map. The executor reads the SQL file for each procedure and executes it using `sp_executesql` with parameter binding.

## Usage

### Basic Usage

```go
package main

import (
    "context"
    "database/sql"
    "cc520-company-dbms-project/src/db/executor"
)

func main() {
    // Connect to database
    db, err := sql.Open("sqlserver", connStr)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create executor
    exec := executor.NewExecutor(db, "src/company/sql", "company")

    // Execute a stored procedure (registry is initialized automatically on first call)
    ctx := context.Background()
    result := exec.Execute(ctx, "sp_CreateCompany", map[string]interface{}{
        "Name": "Acme Corp",
    })

    if result.Error != nil {
        panic(result.Error)
    }

    println("Rows affected:", result.RowsAffected)
}
```

### Supported Procedure Names

- Full name with `sp_` prefix: `"sp_CreateCompany"`
- The registry uses the name without the prefix (e.g., `"CreateCompany"`)

### Parameter Types

The executor supports the following parameter types:

| Type | SQL Type | Example |
|------|----------|---------|
| `string` | `NVARCHAR` or `VARCHAR` | `"Acme Corp"` |
| `int` | `INT` | `123` |
| `int64` | `BIGINT` | `1234567890123456789` |
| `float64` | `FLOAT` or `DECIMAL` | `99.99` |
| `bool` | `BIT` | `true` |

### Return Values

The `Result` struct contains:

- `RowsAffected int64`: Number of rows affected by the operation
- `LastInsertID sql.NullInt64`: Last inserted identity value (if available)
- `Error error`: Error if the execution failed

## Procedure File Location

The executor searches for procedure SQL files in the specified procedure directories. By default, it searches in the `company` directory.

File naming conventions supported:
- `CreateCompany.sql`
- `sp_CreateCompany.sql`
- `Org.CreateCompany.sql`
- `Org.sp_CreateCompany.sql`

## Example: Create Company

```go
result := exec.Execute(ctx, "sp_CreateCompany", map[string]interface{}{
    "Name": "Acme Corp",
})

if result.Error != nil {
    log.Fatal(result.Error)
}

println("Created company with ID:", result.LastInsertID.Int64)
```

## Example: Get All Companies

```go
result := exec.Execute(ctx, "sp_GetAllCompanies", map[string]interface{}{})

if result.Error != nil {
    log.Fatal(result.Error)
}

// Result will contain rows affected (number of rows returned)
println("Found", result.RowsAffected, "companies")
```

## Example: Update Company

```go
result := exec.Execute(ctx, "sp_UpdateCompany", map[string]interface{}{
    "CompanyID": 1,
    "Name": "Updated Corp",
})

if result.Error != nil {
    log.Fatal(result.Error)
}

println("Updated", result.RowsAffected, "row(s)")
```

## Example: Delete Company

```go
result := exec.Execute(ctx, "sp_DeleteCompany", map[string]interface{}{
    "CompanyID": 1,
})

if result.Error != nil {
    log.Fatal(result.Error)
}

println("Deleted", result.RowsAffected, "row(s)")
```

## Error Handling

Always check the `Error` field in the `Result` struct:

```go
result := exec.Execute(ctx, "sp_CreateCompany", map[string]interface{}{
    "Name": "Acme Corp",
})

if result.Error != nil {
    log.Printf("Failed to execute procedure: %v", result.Error)
    return
}
```

## Notes

- The executor uses `sp_executesql` to execute stored procedures with parameters
- Parameters are automatically bound based on their types
- String parameters are escaped to prevent SQL injection
- The executor uses a procedure registry for fast O(1) lookups
- The registry is lazily initialized on the first `Execute()` call
- The registry is built by scanning all `.sql` files in procedure directories
