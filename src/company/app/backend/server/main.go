// Package main provides the entry point for the company backend server
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 90%

package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cc520-company-dbms-project/src/company/app/backend"
	"cc520-company-dbms-project/src/company/app/backend/repositories"
	"cc520-company-dbms-project/src/company/db/executor"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	// Get database connection string from environment
	dbConnStr := os.Getenv("SQL_SERVER_CONNECTION_STRING")
	if dbConnStr == "" {
		log.Fatal("SQL_SERVER_CONNECTION_STRING environment variable is not set")
	}

	// Open database connection
	db, err := sql.Open("sqlserver", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create executor
	executor := executor.NewExecutor(db)

	// Create repositories
	companyRepo := repositories.NewCompanyRepository(executor)
	deptRepo := repositories.NewDepartmentRepository(executor)

	// Create handlers
	companiesHandler := backend.NewCompaniesHandler(companyRepo)
	departmentsHandler := backend.NewDepartmentsHandler(deptRepo)

	// Create router with gorilla/mux
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/companies", companiesHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/departments", departmentsHandler.GetAll).Methods(http.MethodGet)

	// Create server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Print API URL
	log.Printf("Server starting on %s", server.Addr)
	log.Printf("API URL: http://localhost%s", server.Addr)

	// Start server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	}
}
