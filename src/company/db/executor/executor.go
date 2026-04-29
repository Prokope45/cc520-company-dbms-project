// Package executor provides a general SQL query executor for stored procedures
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 90%

package executor

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"cc520-company-dbms-project/src/company/db/executor/registry"
)

// Result contains the result of executing a stored procedure
type Result struct {
	RowsAffected int64
	LastInsertID sql.NullInt64
	Rows         []map[string]interface{}
	Error        error
}

// Executor executes stored procedures with parameters
type Executor struct {
	db       *sql.DB
	registry *registry.Registry
}

// NewExecutor creates a new Executor instance
// baseDir should point to the SQL Procedures directory (e.g., .../sql/Procedures)
// procedureDirs are subdirectories within baseDir to scan (e.g., ["company"])
func NewExecutor(db *sql.DB) *Executor {

	return &Executor{
		db:       db,
		registry: registry.NewRegistry(),
	}
}

// Execute runs a stored procedure by name with the given parameters
// Procedure name should include the sp_ prefix (e.g., "sp_CreateCompany")
// Parameters are passed as a map where keys match the procedure parameter names
func (e *Executor) Execute(ctx context.Context, procedureName string, params map[string]interface{}) Result {
	if !strings.Contains(procedureName, "sp_") {
		procedureName = fmt.Sprintf("sp_%s", procedureName)
	}

	// Check for procedure existence
	validatedProcedureName, err := e.registry.Get(procedureName)
	if err != nil {
		return Result{Error: fmt.Errorf("procedure not found: %s", procedureName)}
	}

	// Execute the stored procedure with parameters
	return e.executeWithParams(ctx, validatedProcedureName, params)
}

// executeWithParams executes the SQL content with the given parameters
func (e *Executor) executeWithParams(ctx context.Context, procName string, params map[string]interface{}) Result {
	// Build the parameter string for sp_executesql
	paramStr := e.buildParamString(params)

	// Execute using sp_executesql with parameters
	query := fmt.Sprintf("EXEC [Org].%s %s", procName, paramStr)

	// In go-mssqldb, LastInsertId is not supported. We need to check if this is a procedure
	// that might return an identity (like Create) or just rows (like Get)
	if strings.Contains(procName, "Create") || strings.Contains(procName, "Get") {
		rows, err := e.db.QueryContext(ctx, query)
		if err != nil {
			return Result{Error: fmt.Errorf("failed to execute stored procedure: %w", err)}
		}
		defer rows.Close()

		var rowsAffected int64 = 0
		var lastInsertID sql.NullInt64
		var rowsData []map[string]interface{}

		for rows.Next() {
			cols, err := rows.Columns()
			if err != nil {
				return Result{Error: fmt.Errorf("failed to get columns: %w", err)}
			}

			rowData := make(map[string]interface{}, len(cols))
			values := make([]interface{}, len(cols))
			valuePtrs := make([]interface{}, len(cols))

			for i := range cols {
				valuePtrs[i] = &values[i]
			}

			if err := rows.Scan(valuePtrs...); err != nil {
				return Result{Error: fmt.Errorf("failed to scan row: %w", err)}
			}

			for i, col := range cols {
				if val, ok := values[i].([]byte); ok {
					rowData[col] = string(val)
				} else {
					rowData[col] = values[i]
				}
			}

			rowsData = append(rowsData, rowData)
			rowsAffected++

			// Check for identity column (first column is typically the identity)
			if len(cols) > 0 {
				if val, ok := values[0].(int64); ok {
					lastInsertID = sql.NullInt64{Int64: val, Valid: true}
				} else if val, ok := values[0].(int32); ok {
					lastInsertID = sql.NullInt64{Int64: int64(val), Valid: true}
				} else if val, ok := values[0].([]byte); ok {
					var parsed int64
					fmt.Sscanf(string(val), "%d", &parsed)
					lastInsertID = sql.NullInt64{Int64: parsed, Valid: true}
				}
			}
		}

		if err := rows.Err(); err != nil {
			return Result{Error: fmt.Errorf("procedure execution failed: %w", err)}
		}

		return Result{
			RowsAffected: rowsAffected,
			LastInsertID: lastInsertID,
			Rows:         rowsData,
		}
	}

	// For Update, Delete, etc. use ExecContext to get rows affected
	result, err := e.db.ExecContext(ctx, query)
	if err != nil {
		return Result{Error: fmt.Errorf("failed to execute stored procedure: %w", err)}
	}

	rowsAffected, _ := result.RowsAffected()

	return Result{
		RowsAffected: rowsAffected,
		LastInsertID: sql.NullInt64{},
	}
}

// buildParamString builds the parameter string for sp_executesql
func (e *Executor) buildParamString(params map[string]interface{}) string {
	var paramParts []string

	for paramName, paramValue := range params {
		// Convert parameter value to appropriate SQL type
		var paramValueStr string
		switch v := paramValue.(type) {
		case string:
			// Escape single quotes and wrap in single quotes
			escaped := strings.ReplaceAll(v, "'", "''")
			paramValueStr = fmt.Sprintf("@%s = N'%s'", paramName, escaped)
		case int:
			paramValueStr = fmt.Sprintf("@%s = %d", paramName, v)
		case int64:
			paramValueStr = fmt.Sprintf("@%s = %d", paramName, v)
		case float64:
			paramValueStr = fmt.Sprintf("@%s = %g", paramName, v)
		case bool:
			paramValueStr = fmt.Sprintf("@%s = %t", paramName, v)
		default:
			// For other types, convert to string
			paramValueStr = fmt.Sprintf("@%s = '%s'", paramName, fmt.Sprintf("%v", v))
		}
		paramParts = append(paramParts, paramValueStr)
	}

	return strings.Join(paramParts, ", ")
}
