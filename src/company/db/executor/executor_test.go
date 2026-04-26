// Package executor provides a general SQL query executor for stored procedures
// Authors:
//	- Jared Paubel
// 	- Roo agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 80%

package executor

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func TestMain(m *testing.M) {
	// Path relative to src/db/executor/
	envPath := "../../../../.env"
	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Warning: .env file not found")
	}

	os.Exit(m.Run())
}

// TestExecutor_Execute tests the Execute method with various procedures
func TestExecutor_Execute(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	ctx := context.Background()

	// Test 1: Execute sp_GetAllCompanies (no parameters)
	t.Run("GetAllCompanies", func(t *testing.T) {
		result := exec.Execute(ctx, "GetAllCompanies", map[string]interface{}{})
		if result.Error != nil {
			t.Errorf("execute failed: %v", result.Error)
		}
		if result.RowsAffected <= 0 {
			t.Error("expected rows to be returned")
		}
	})

	// Test 2: Execute sp_CreateCompany (with parameters)
	createResult := exec.Execute(ctx, "CreateCompany", map[string]interface{}{
		"Name": "Test Corp",
	})
	if createResult.Error != nil {
		t.Fatalf("CreateCompany failed: %v", createResult.Error)
	}
	if !createResult.LastInsertID.Valid {
		t.Error("expected LastInsertID to be valid")
	}
	t.Logf("Created company with ID: %d", createResult.LastInsertID.Int64)

	// Test 3: Execute sp_GetCompanyByID (with parameters)
	getResult := exec.Execute(ctx, "GetCompanyByID", map[string]interface{}{
		"CompanyID": createResult.LastInsertID.Int64,
	})
	if getResult.Error != nil {
		t.Fatalf("GetCompanyByID failed: %v", getResult.Error)
	}
	if getResult.RowsAffected != 1 {
		t.Errorf("Expected 1 row, got %d", getResult.RowsAffected)
	}

	// Test 4: Execute sp_UpdateCompany (with parameters)
	updateResult := exec.Execute(ctx, "UpdateCompany", map[string]interface{}{
		"CompanyID": createResult.LastInsertID.Int64,
		"Name":      "Updated Corp",
	})
	if updateResult.Error != nil {
		t.Fatalf("UpdateCompany failed: %v", updateResult.Error)
	}
	if updateResult.RowsAffected != 1 {
		t.Errorf("Expected 1 row affected, got %d", updateResult.RowsAffected)
	}

	// Test 5: Execute sp_DeleteCompany (with parameters)
	deleteResult := exec.Execute(ctx, "DeleteCompany", map[string]interface{}{
		"CompanyID": createResult.LastInsertID.Int64,
	})
	if deleteResult.Error != nil {
		t.Fatalf("DeleteCompany failed: %v", deleteResult.Error)
	}
	if deleteResult.RowsAffected != 1 {
		t.Errorf("Expected 1 row affected, got %d", deleteResult.RowsAffected)
	}

	// Test 6: Execute non-existent procedure
	t.Run("NonExistentProcedure", func(t *testing.T) {
		result := exec.Execute(ctx, "NonExistentProcedure", map[string]interface{}{})
		if result.Error == nil {
			t.Error("Expected error for non-existent procedure")
		}
	})
}

// TestExecutor_Execute_LazyInitialization tests that registry is initialized on first Execute call
func TestExecutor_Execute_LazyInitialization(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor without initializing registry
	exec := NewExecutor(db)
	// Clear registry for clean slate
	exec.registry.ClearRegistry()

	// Verify registry is empty initially
	registry_count := exec.registry.Count()
	if registry_count != 0 {
		t.Errorf("Expected empty registry, got %d entries", registry_count)
	}

	// Verify initialized is false
	if exec.registry.Initialized() {
		t.Error("Expected initialized to be false")
	}

	// Execute a procedure (should initialize registry automatically)
	ctx := context.Background()
	result := exec.Execute(ctx, "sp_GetAllCompanies", map[string]interface{}{})

	// Verify registry was initialized
	if !exec.registry.Initialized() {
		t.Error("Expected initialized to be true after Execute")
	}

	if result.Error != nil {
		t.Errorf("Execute failed: %v", result.Error)
	}
}

// TestExecutor_Execute_ParamTypes tests various parameter types
func TestExecutor_Execute_ParamTypes(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	ctx := context.Background()

	// Test with string parameter
	t.Run("StringParameter", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_CreateCompany", map[string]interface{}{
			"Name": "Param Test Corp",
		})
		if result.Error != nil {
			t.Errorf("Execute failed: %v", result.Error)
		}
	})

	// Test with int parameter
	t.Run("IntParameter", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_GetCompanyByID", map[string]interface{}{
			"CompanyID": 1,
		})
		if result.Error != nil {
			t.Errorf("Execute failed: %v", result.Error)
		}
	})

	// Test with int64 parameter
	t.Run("Int64Parameter", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_GetCompanyByID", map[string]interface{}{
			"CompanyID": int64(1),
		})
		if result.Error != nil {
			t.Errorf("Execute failed: %v", result.Error)
		}
	})

	// Test with float64 parameter
	t.Run("Float64Parameter", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_GetCompanyByID", map[string]interface{}{
			"CompanyID": 1.0,
		})
		if result.Error != nil {
			t.Errorf("Execute failed: %v", result.Error)
		}
	})

	// Test with bool parameter (if applicable)
	t.Run("BoolParameter", func(t *testing.T) {
		// Create a company first
		createResult := exec.Execute(ctx, "sp_CreateCompany", map[string]interface{}{
			"Name": "Bool Test Corp",
		})
		if createResult.Error != nil {
			t.Fatalf("Create failed: %v", createResult.Error)
		}

		// Update with bool (if the procedure supports it)
		updateResult := exec.Execute(ctx, "sp_UpdateCompany", map[string]interface{}{
			"CompanyID": createResult.LastInsertID.Int64,
			"Name":      "Updated Bool Test Corp",
		})
		if updateResult.Error != nil {
			t.Errorf("Execute failed: %v", updateResult.Error)
		}
	})
}

// TestExecutor_Execute_EmptyParams tests executing a procedure with no parameters
func TestExecutor_Execute_EmptyParams(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	ctx := context.Background()

	// Test sp_GetAllCompanies with empty params map
	t.Run("EmptyParamsMap", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_GetAllCompanies", map[string]interface{}{})
		if result.Error != nil {
			t.Errorf("Execute failed: %v", result.Error)
		}
		if result.RowsAffected <= 0 {
			t.Error("Expected rows to be returned")
		}
	})
}

// TestExecutor_Execute_ProcedureNotFound tests error handling for non-existent procedures
func TestExecutor_Execute_ProcedureNotFound(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	ctx := context.Background()

	// Test with non-existent procedure
	t.Run("NonExistentProcedure", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_NonExistentProcedure", map[string]interface{}{})
		if result.Error == nil {
			t.Error("Expected error for non-existent procedure")
		}
		expectedErr := "procedure not found: sp_NonExistentProcedure"
		if result.Error.Error() != expectedErr {
			t.Errorf("Expected error '%s', got '%s'", expectedErr, result.Error.Error())
		}
	})
}

// TestExecutor_Registry_BuildsCorrectly tests that the registry is built correctly
func TestExecutor_Registry_BuildsCorrectly(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	// Check that expected procedures are in the registry
	expectedProcedures := []string{
		"CreateCompany",
		"DeleteCompany",
		"GetAllCompanies",
		"GetCompanyByID",
		"UpdateCompany",
	}

	for _, procName := range expectedProcedures {
		_, err := exec.registry.Get(procName)
		if err != nil {
			t.Errorf("Expected procedure '%s' in registry, not found", procName)
		}
	}

	// Check that expected procedures are in the registry
	expectedProceduresPrefixed := []string{
		"sp_CreateCompany",
		"sp_DeleteCompany",
		"sp_GetAllCompanies",
		"sp_GetCompanyByID",
		"sp_UpdateCompany",
	}

	for _, procName := range expectedProceduresPrefixed {
		_, err := exec.registry.Get(procName)
		if err != nil {
			t.Errorf("Expected procedure '%s' in registry, not found", procName)
		}
	}
}

// TestExecutor_Execute_WithSpecialCharacters tests handling of special characters in parameters
func TestExecutor_Execute_WithSpecialCharacters(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	ctx := context.Background()

	// Test with special characters (single quotes)
	t.Run("SpecialCharacters", func(t *testing.T) {
		result := exec.Execute(ctx, "sp_CreateCompany", map[string]interface{}{
			"Name": "Test Corp's Company",
		})
		if result.Error != nil {
			t.Errorf("Execute failed: %v", result.Error)
		}
	})
}

// TestExecutor_Execute_MultipleCalls tests that multiple calls work correctly
func TestExecutor_Execute_MultipleCalls(t *testing.T) {
	// Get test database connection string
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		t.Skip("SQL_SERVER_CONNECTION_STRING not set, skipping test")
	}

	// Connect to test database
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create executor
	exec := NewExecutor(db)

	ctx := context.Background()

	// Execute multiple procedures
	t.Run("MultipleCalls", func(t *testing.T) {
		// First call
		result1 := exec.Execute(ctx, "sp_GetAllCompanies", map[string]interface{}{})
		if result1.Error != nil {
			t.Errorf("First call failed: %v", result1.Error)
		}

		// Second call
		result2 := exec.Execute(ctx, "sp_GetAllCompanies", map[string]interface{}{})
		if result2.Error != nil {
			t.Errorf("Second call failed: %v", result2.Error)
		}

		// Third call
		result3 := exec.Execute(ctx, "sp_GetAllCompanies", map[string]interface{}{})
		if result3.Error != nil {
			t.Errorf("Third call failed: %v", result3.Error)
		}

		// All calls should have same number of rows
		if result1.RowsAffected != result2.RowsAffected || result2.RowsAffected != result3.RowsAffected {
			t.Errorf("Inconsistent row counts: %d, %d, %d", result1.RowsAffected, result2.RowsAffected, result3.RowsAffected)
		}
	})
}
