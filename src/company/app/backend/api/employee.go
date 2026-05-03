// Package api provides HTTP API handlers for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/app/backend/repositories"

	"github.com/gorilla/mux"
)

// EmployeesHandler handles employee-related HTTP requests
type EmployeesHandler struct {
	repo *repositories.EmployeeRepository
}

// NewEmployeesHandler creates a new EmployeesHandler
func NewEmployeesHandler(repo *repositories.EmployeeRepository) *EmployeesHandler {
	return &EmployeesHandler{repo: repo}
}

// GetAll handles GET /employees
func (h *EmployeesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := h.repo.GetAllEmployees(r.Context())
	if err != nil {
		http.Error(w, FormatDBError(err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

// GetByID handles GET /employees/{id}
func (h *EmployeesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee, err := h.repo.GetEmployeeByID(r.Context(), id)
	if err != nil {
		http.Error(w, FormatDBError(err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

// Create handles POST /employees
func (h *EmployeesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newEmployee, err := h.repo.CreateEmployee(r.Context(), employee)
	if err != nil {
		http.Error(w, FormatDBError(err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEmployee)
}

// Update handles PUT /employees/{id}
func (h *EmployeesHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var employee models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	employee.EmployeeID = id
	updatedEmployee, err := h.repo.UpdateEmployee(r.Context(), employee)
	if err != nil {
		http.Error(w, FormatDBError(err), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEmployee)
}

// Delete handles DELETE /employees/{id}
func (h *EmployeesHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.repo.DeleteEmployee(r.Context(), id)
	if err != nil {
		http.Error(w, FormatDBError(err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
