// Package main provides the database rebuild entry point
package main

import (
	"cc520-company-dbms-project/src/db/rebuild"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

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
	rebuilder := rebuild.NewRebuilder(dbConn, "src/company/sql")

	// Rebuild the database
	ctx := context.Background()
	if err := rebuilder.Rebuild(ctx); err != nil {
		log.Fatalf("Failed to rebuild database: %v", err)
	}

	fmt.Println("Database rebuild completed successfully")
}
