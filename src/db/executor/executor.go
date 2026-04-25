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
	"os"
	"path/filepath"
	"strings"
)

// Result contains the result of executing a stored procedure
type Result struct {
	RowsAffected int64
	LastInsertID sql.NullInt64
	Error        error
}

// ProcedureRegistry maps procedure names to their SQL file paths
type ProcedureRegistry map[string]string

// Executor executes stored procedures with parameters
type Executor struct {
	db          *sql.DB
	baseDir     string
	procedures  []string // list of procedure directories to search
	registry    ProcedureRegistry
	initialized bool
}

// NewExecutor creates a new Executor instance
func NewExecutor(db *sql.DB, baseDir string, procedureDirs ...string) *Executor {
	if len(procedureDirs) == 0 {
		procedureDirs = []string{"Procedures/company"} // default to company directory
	}
	return &Executor{
		db:         db,
		baseDir:    baseDir,
		procedures: procedureDirs,
		registry:   make(ProcedureRegistry),
	}
}

// initializeRegistry scans procedure directories and builds the registry
// This should be called after NewExecutor to enable fast procedure lookups
func (e *Executor) initializeRegistry() error {
	for _, procDir := range e.procedures {
		procPath := filepath.Join(e.baseDir, procDir)
		entries, err := os.ReadDir(procPath)
		if err != nil {
			return fmt.Errorf("failed to read procedure directory %s: %w", procPath, err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			// Only process .sql files
			if !strings.HasSuffix(entry.Name(), ".sql") {
				continue
			}

			// Extract procedure name from filename
			// Remove .sql extension
			filename := entry.Name()
			if filename == "Org.sql" {
				continue // Skip the wildcard SELECT file
			}

			// Remove .sql extension
			name := strings.TrimSuffix(filename, ".sql")

			// Remove sp_ prefix if present for registry key
			procName := strings.TrimPrefix(name, "sp_")

			// Store the full path to the SQL file
			sqlPath := filepath.Join(procPath, filename)
			e.registry[procName] = sqlPath
		}
	}
	return nil
}

// Execute runs a stored procedure by name with the given parameters
// Procedure name should include the sp_ prefix (e.g., "sp_CreateCompany")
// Parameters are passed as a map where keys match the procedure parameter names
func (e *Executor) Execute(ctx context.Context, procedureName string, params map[string]interface{}) Result {
	// Initialize the registry if not already done
	if !e.initialized {
		if err := e.initializeRegistry(); err != nil {
			return Result{Error: fmt.Errorf("failed to initialize registry: %w", err)}
		}
		e.initialized = true
	}

	// Look up the SQL file from the registry
	sqlPath, ok := e.registry[procedureName]
	if !ok {
		return Result{Error: fmt.Errorf("procedure not found: %s", procedureName)}
	}

	// Read the SQL file
	sqlContent, err := os.ReadFile(sqlPath)
	if err != nil {
		return Result{Error: fmt.Errorf("failed to read procedure SQL file: %w", err)}
	}

	// Execute the stored procedure with parameters
	return e.executeWithParams(ctx, string(sqlContent), params)
}

// executeWithParams executes the SQL content with the given parameters
func (e *Executor) executeWithParams(ctx context.Context, sqlContent string, params map[string]interface{}) Result {
	// Extract the procedure name from the SQL content
	// Look for CREATE PROCEDURE [sp_...]
	procName := extractProcedureName(sqlContent)

	// Build the parameter string for sp_executesql
	paramStr := e.buildParamString(params)

	// Execute using sp_executesql with parameters
	query := fmt.Sprintf("EXEC %s %s", procName, paramStr)

	result, err := e.db.ExecContext(ctx, query)
	if err != nil {
		return Result{Error: fmt.Errorf("failed to execute stored procedure: %w", err)}
	}

	// Get rows affected
	rowsAffected, _ := result.RowsAffected()

	// Get last insert ID if available
	var lastInsertID sql.NullInt64
	if rowsAffected > 0 {
		lastInsertIDVal, _ := result.LastInsertId()
		lastInsertID = sql.NullInt64{Int64: lastInsertIDVal, Valid: true}
	}

	return Result{
		RowsAffected: rowsAffected,
		LastInsertID: lastInsertID,
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
			paramValueStr = fmt.Sprintf("@%s N'%s'", paramName, escaped)
		case int:
			paramValueStr = fmt.Sprintf("@%s %d", paramName, v)
		case int64:
			paramValueStr = fmt.Sprintf("@%s %d", paramName, v)
		case float64:
			paramValueStr = fmt.Sprintf("@%s %g", paramName, v)
		case bool:
			paramValueStr = fmt.Sprintf("@%s %t", paramName, v)
		default:
			// For other types, convert to string
			paramValueStr = fmt.Sprintf("@%s '%s'", paramName, fmt.Sprintf("%v", v))
		}
		paramParts = append(paramParts, paramValueStr)
	}

	return strings.Join(paramParts, ", ")
}

// extractProcedureName extracts the procedure name from SQL content
func extractProcedureName(sqlContent string) string {
	// Look for CREATE PROCEDURE [sp_...]
	lines := strings.Split(sqlContent, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "CREATE PROCEDURE") || strings.HasPrefix(line, "CREATE OR ALTER PROCEDURE") {
			// Extract the procedure name
			start := strings.Index(line, "[")
			end := strings.Index(line, "]")
			if start != -1 && end != -1 && end > start {
				return line[start+1 : end]
			}
		}
	}
	return ""
}
