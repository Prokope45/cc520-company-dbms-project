// Package rebuild provides database rebuilding functionality
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 10%

package rebuild

import (
	"cc520-company-dbms-project/src/company/db/executor/registry"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/microsoft/go-mssqldb"
)

// Rebuilder handles database rebuilding operations
type Rebuilder struct {
	db       *sql.DB
	baseDir  string
	registry *registry.Registry
}

// NewRebuilder creates a new Rebuilder instance
func NewRebuilder(db *sql.DB, baseDir string) *Rebuilder {
	return &Rebuilder{
		db:       db,
		baseDir:  baseDir,
		registry: registry.NewRegistry(),
	}
}

// Run executes the requested database operation with dependency resolution
func (r *Rebuilder) Run(ctx context.Context, operation string) error {
	var possible_error error = nil

	switch operation {
	case "rebuild":
		possible_error = r.rebuild(ctx)
	case "clear":
		possible_error = r.clearDatabase(ctx)
	case "schema":
		possible_error = r.createSchema(ctx)
	case "procedures":
		possible_error = r.createProcedures(ctx)
	case "tables":
		possible_error = r.createTables(ctx)
	case "seed":
		// Auto-resolve dependencies: create schema & tables before seeding
		fmt.Println("\n[seed] Creating schema (dependency)...")
		if err := r.createSchema(ctx); err != nil {
			return fmt.Errorf("failed to create schema for seed: %w", err)
		}
		fmt.Println("[seed] Creating tables (dependency)...")
		if err := r.createTables(ctx); err != nil {
			return fmt.Errorf("failed to create tables for seed: %w", err)
		}
		possible_error = r.seedData(ctx)
	default:
		r.Close()
		return fmt.Errorf("unknown operation: %s (valid: rebuild, clear, schema, tables, seed)", operation)
	}
	r.Close()
	return possible_error
}

// Rebuild rebuilds the database by clearing it and recreating schemas, tables, and seeding data
func (r *Rebuilder) rebuild(ctx context.Context) error {
	fmt.Println("=== Starting Database Rebuild ===")

	// Step 1: Clear the database
	fmt.Println("\n[Step 1] Clearing database...")
	if err := r.clearDatabase(ctx); err != nil {
		return fmt.Errorf("failed to clear database: %w", err)
	}
	fmt.Println("[x] Database cleared")

	// Step 2: Create schema
	fmt.Println("\n[Step 2] Creating schema...")
	if err := r.createSchema(ctx); err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}
	fmt.Println("[x] Schema created")

	// Step 3: Create tables in dependency order
	fmt.Println("\n[Step 3] Creating tables...")
	if err := r.createTables(ctx); err != nil {
		return fmt.Errorf("failed to create tables: %w", err)
	}
	fmt.Println("[x] Tables created")

	// Step 4: Create stored procedures
	fmt.Println("\n[Step 4] Creating stored procedures...")
	if err := r.createProcedures(ctx); err != nil {
		return fmt.Errorf("failed to create stored procedures: %w", err)
	}
	fmt.Println("[x] Stored procedures created")

	// Step 5: Seed data
	fmt.Println("\n[Step 5] Seeding data...")
	if err := r.seedData(ctx); err != nil {
		return fmt.Errorf("failed to seed data: %w", err)
	}
	fmt.Println("[x] Data seeded")

	fmt.Println("\n=== Database Rebuild Complete ===")
	return nil
}

// clearDatabase drops all tables and schemas in the Org schema
func (r *Rebuilder) clearDatabase(ctx context.Context) error {
	// Drop tables in reverse dependency order to avoid foreign key constraint violations
	dropOrder := []string{
		"DepEmpRole",
		"HourlyEmployee",
		"SalaryEmployee",
		"Salary",
		"Employee",
		"Department",
		"Person",
		"Company",
		"Role",
		"EmployeeType",
		"Address",
	}

	for _, tableName := range dropOrder {
		query := fmt.Sprintf("DROP TABLE IF EXISTS [Org].[%s]", tableName)
		if _, err := r.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to drop table %s: %w", tableName, err)
		}
		fmt.Printf("  Dropped table: %s\n", tableName)
	}

	for proc := range r.registry.GetAll() {
		query := fmt.Sprintf("DROP PROCEDURE IF EXISTS [Org].[%s]", proc)
		if _, err := r.db.ExecContext(ctx, query); err != nil {
			return fmt.Errorf("failed to drop procedure %s: %w", proc, err)
		}
		fmt.Printf("  Dropped procedure: %s\n", proc)
	}

	// Drop the schema
	query := "DROP SCHEMA IF EXISTS [Org]"
	if _, err := r.db.ExecContext(ctx, query); err != nil {
		return fmt.Errorf("failed to drop schema: %w", err)
	}
	fmt.Println("  Dropped schema: Org")

	return nil
}

// createSchema creates the Org schema
func (r *Rebuilder) createSchema(ctx context.Context) error {
	// Read and execute the schema creation SQL file
	schemaSQLPath := filepath.Join(r.baseDir, "Schemas", "Org.sql")
	schemaSQL, err := os.ReadFile(schemaSQLPath)
	if err != nil {
		return fmt.Errorf("failed to read schema SQL file: %w", err)
	}

	// Execute the schema creation script
	if _, err := r.db.ExecContext(ctx, string(schemaSQL)); err != nil {
		return fmt.Errorf("failed to execute schema creation: %w", err)
	}

	return nil
}

// createTables creates all tables in the correct dependency order
func (r *Rebuilder) createTables(ctx context.Context) error {
	// Define tables in dependency order (tables with no dependencies first)
	tableOrder := []string{
		"Address",
		"EmployeeType",
		"Role",
		"Company",
		"Person",         // Depends on Address
		"Department",     // Depends on Company
		"Employee",       // Depends on Person, EmployeeType, HourlyEmployee
		"Salary",         // Depends on Salary
		"SalaryEmployee", // Depends on Employee, Salary
		"HourlyEmployee", // Depends on Employee
		"DepEmpRole",     // Depends on Role, Department, Employee
	}

	for _, tableName := range tableOrder {
		tableSQLPath := filepath.Join(r.baseDir, "Tables", fmt.Sprintf("Org.%s.sql", tableName))
		tableSQL, err := os.ReadFile(tableSQLPath)
		if err != nil {
			return fmt.Errorf("failed to read table SQL file for %s: %w", tableName, err)
		}

		if _, err := r.db.ExecContext(ctx, string(tableSQL)); err != nil {
			return fmt.Errorf("failed to create table %s: %w", tableName, err)
		}
		fmt.Printf("  Created table: %s\n", tableName)
	}

	return nil
}

// createProcedures creates all stored procedures
func (r *Rebuilder) createProcedures(ctx context.Context) error {
	proceduresDir := filepath.Join(r.baseDir, "Procedures", "company")
	entries, err := os.ReadDir(proceduresDir)
	if err != nil {
		// If directory doesn't exist, just skip
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read procedures directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		procPath := filepath.Join(proceduresDir, entry.Name())
		procSQL, err := os.ReadFile(procPath)
		if err != nil {
			return fmt.Errorf("failed to read procedure SQL file %s: %w", entry.Name(), err)
		}

		// Split by GO since go-mssqldb doesn't support it natively
		batches := strings.Split(string(procSQL), "GO")
		for _, batch := range batches {
			batch = strings.TrimSpace(batch)
			if batch != "" {
				if _, err := r.db.ExecContext(ctx, batch); err != nil {
					return fmt.Errorf("failed to create procedure %s: %w", entry.Name(), err)
				}
			}
		}
		fmt.Printf("  Created procedure: %s\n", entry.Name())
	}
	return nil
}

// seedData seeds all tables with data
func (r *Rebuilder) seedData(ctx context.Context) error {
	// Define data files in the correct order (respecting foreign key dependencies)
	dataOrder := []string{
		"Address",
		"EmployeeType",
		"Role",
		"Company",
		"Person",         // Depends on Address
		"Department",     // Depends on Company
		"Employee",       // Depends on Person, EmployeeType
		"Salary",         // Depends on Salary
		"SalaryEmployee", // Depends on Employee, Salary
		"HourlyEmployee", // Depends on Employee
		"DepEmpRole",     // Depends on Role, Department, Employee
	}

	for _, dataFile := range dataOrder {
		dataSQLPath := filepath.Join(r.baseDir, "Data", fmt.Sprintf("Org.%s.sql", dataFile))
		dataSQL, err := os.ReadFile(dataSQLPath)
		if err != nil {
			return fmt.Errorf("failed to read data SQL file for %s: %w", dataFile, err)
		}

		// Split the SQL into batches by "GO" statements
		batches := strings.Split(string(dataSQL), "GO")
		for _, batch := range batches {
			batch = strings.TrimSpace(batch)
			if batch != "" {
				if _, err := r.db.ExecContext(ctx, batch); err != nil {
					return fmt.Errorf("failed to seed data for %s: %w", dataFile, err)
				}
			}
		}
		fmt.Printf("  Seeded data for: %s\n", dataFile)
	}

	return nil
}

// Close closes the database connection
func (r *Rebuilder) Close() error {
	if r.db != nil {
		return r.db.Close()
	}
	return nil
}
