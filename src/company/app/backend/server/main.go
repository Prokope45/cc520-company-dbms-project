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

	"cc520-company-dbms-project/src/company/app/backend/api"
	"cc520-company-dbms-project/src/company/app/backend/repositories"

	"cc520-company-dbms-project/src/company/db/executor"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

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
	employeeRepo := repositories.NewEmployeeRepository(executor)
	reportsRepo := repositories.NewReportsRepository(executor)
	roleRepo := repositories.NewRoleRepository(executor)

	// Create handlers
	companiesHandler := api.NewCompanyHandler(companyRepo)
	departmentsHandler := api.NewDepartmentsHandler(deptRepo)
	employeesHandler := api.NewEmployeesHandler(employeeRepo)
	reportsHandler := api.NewReportsHandler(reportsRepo)
	rolesHandler := api.NewRoleHandler(roleRepo)

	// Create router with gorilla/mux
	r := mux.NewRouter()

	// API routes - Companies
	r.HandleFunc("/companies", companiesHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/companies", companiesHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/companies/{id}", companiesHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/companies/{id}", companiesHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/companies/{id}", companiesHandler.Delete).Methods(http.MethodDelete)

	// API routes - Roles
	r.HandleFunc("/roles", rolesHandler.GetAll).Methods(http.MethodGet)

	// API routes - Departments
	r.HandleFunc("/departments", departmentsHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/departments", departmentsHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/departments/{id}", departmentsHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/departments/{id}", departmentsHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/departments/{id}", departmentsHandler.Delete).Methods(http.MethodDelete)

	// API routes - Employees
	r.HandleFunc("/employees", employeesHandler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/employees", employeesHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/employees/{id}", employeesHandler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/employees/{id}", employeesHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/employees/{id}", employeesHandler.Delete).Methods(http.MethodDelete)

	// API routes - Reports
	r.HandleFunc("/reports/department-salary-ranks", reportsHandler.GetDepartmentSalaryRanks).Methods(http.MethodGet)
	r.HandleFunc("/reports/top-terminated-hourly", reportsHandler.GetTopTerminatedHourly).Methods(http.MethodGet)
	r.HandleFunc("/reports/unhired-with-manager", reportsHandler.GetUnhiredWithManager).Methods(http.MethodGet)
	r.HandleFunc("/reports/highest-paid-ceo", reportsHandler.GetHighestPaidCEO).Methods(http.MethodGet)

	// API routes - Aggregated Reports
	r.HandleFunc("/reports/department-salary-ranks-aggregated", reportsHandler.GetDepartmentSalary_Aggregated).Methods(http.MethodGet)
	r.HandleFunc("/reports/highest-paid-ceo-aggregated", reportsHandler.GetHighestPaidCEO_Aggregated).Methods(http.MethodGet)
	r.HandleFunc("/reports/top-terminated-hourly-aggregated", reportsHandler.GetTopTerminatedHourly_Aggregated).Methods(http.MethodGet)
	r.HandleFunc("/reports/unhired-with-manager-aggregated", reportsHandler.GetUnhiredWithManager_Aggregated).Methods(http.MethodGet)

	// Setup CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Create server with CORS handler
	server := &http.Server{
		Addr:         ":8080",
		Handler:      c.Handler(r),
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
