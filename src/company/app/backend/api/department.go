// Package api provides HTTP API handlers for the company application
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
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

// DepartmentsHandler handles department-related HTTP requests
type DepartmentsHandler struct {
	repo *repositories.DepartmentRepository
}

// NewDepartmentsHandler creates a new DepartmentsHandler
func NewDepartmentsHandler(repo *repositories.DepartmentRepository) *DepartmentsHandler {
	return &DepartmentsHandler{repo: repo}
}

// GetAll handles GET /departments
func (h *DepartmentsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	departments, err := h.repo.GetAllDepartments(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departments)
}

// GetByID handles GET /departments/{id}
func (h *DepartmentsHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	department, err := h.repo.GetDepartmentByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(department)
}

// Create handles POST /departments
func (h *DepartmentsHandler) Create(w http.ResponseWriter, r *http.Request) {
	var department models.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newDepartment, err := h.repo.CreateDepartment(r.Context(), department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newDepartment)
}

// Update handles PUT /departments/{id}
func (h *DepartmentsHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var department models.Department
	if err := json.NewDecoder(r.Body).Decode(&department); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	department.DepartmentID = id
	updatedDepartment, err := h.repo.UpdateDepartment(r.Context(), department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedDepartment)
}

// Delete handles DELETE /departments/{id}
func (h *DepartmentsHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.repo.DeleteDepartment(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
