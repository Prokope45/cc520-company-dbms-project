// Package main provides the database rebuild CLI entry point
package main

import (
	"cc520-company-dbms-project/src/db/rebuild"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// Load environment variables from .env file in the root directory
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	// Get connection string from environment variable
	connStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if connStr == "" {
		log.Fatal("SQL_SERVER_CONNECTION_STRING environment variable is not set")
	}

	// Connect to the database
	dbConn, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer dbConn.Close()

	// Test the connection
	if err := dbConn.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Connected to database successfully")

	// Create the rebuilder
	cwd, _ := os.Getwd()
	sqlPath := filepath.Join(cwd, "src", "company", "sql")
	rebuilder := rebuild.NewRebuilder(dbConn, sqlPath)

	// Parse CLI arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <operation>\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Operations:\n")
		fmt.Fprintf(os.Stderr, "  rebuild   Full rebuild: clear schema, tables, and seed data (default)\n")
		fmt.Fprintf(os.Stderr, "  clear     Drop all tables and schema\n")
		fmt.Fprintf(os.Stderr, "  schema    Create the Org schema only\n")
		fmt.Fprintf(os.Stderr, "  tables    Create all tables (schema must exist)\n")
		fmt.Fprintf(os.Stderr, "  seed      Re-seed data with sample data (creates schema + tables if needed)\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s rebuild\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s seed\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s clear\n", os.Args[0])
	}
	flag.Parse()

	operation := "rebuild"
	args := flag.Args()
	if len(args) > 0 {
		operation = args[0]
	}

	// Run the requested operation
	ctx := context.Background()
	if err := rebuilder.Run(ctx, operation); err != nil {
		flag.Usage()
		log.Fatalf("Failed to run DB operation: %v", err)
	}
}
