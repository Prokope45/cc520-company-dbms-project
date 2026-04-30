// Package backend provides the HTTP API for the company application
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package backend

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/app/backend/repositories"

	"github.com/gorilla/mux"
)

// API holds the HTTP API handlers and repositories
type API struct {
	mu          sync.RWMutex
	companyRepo *repositories.CompanyRepository
	deptRepo    *repositories.DepartmentRepository
}

// NewAPI creates a new API instance
func NewAPI(companyRepo *repositories.CompanyRepository, deptRepo *repositories.DepartmentRepository) *API {
	return &API{
		companyRepo: companyRepo,
		deptRepo:    deptRepo,
	}
}

// CompaniesHandler handles company-related HTTP requests
type CompaniesHandler struct {
	companyRepo *repositories.CompanyRepository
}

// NewCompaniesHandler creates a new CompaniesHandler
func NewCompaniesHandler(companyRepo *repositories.CompanyRepository) *CompaniesHandler {
	return &CompaniesHandler{companyRepo: companyRepo}
}

// GetAll handles GET /companies
func (h *CompaniesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	companies, err := h.companyRepo.GetAllCompanies(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

// GetByID handles GET /companies/{id}
func (h *CompaniesHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

	company, err := h.companyRepo.GetCompanyByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(company)
}

// Create handles POST /companies
func (h *CompaniesHandler) Create(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	newCompany, err := h.companyRepo.CreateCompany(r.Context(), company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCompany)
}

// Update handles PUT /companies/{id}
func (h *CompaniesHandler) Update(w http.ResponseWriter, r *http.Request) {
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

	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	company.CompanyID = id
	updatedCompany, err := h.companyRepo.UpdateCompany(r.Context(), company)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCompany)
}

// Delete handles DELETE /companies/{id}
func (h *CompaniesHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

	err = h.companyRepo.DeleteCompany(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DepartmentsHandler handles department-related HTTP requests
type DepartmentsHandler struct {
	deptRepo *repositories.DepartmentRepository
}

// NewDepartmentsHandler creates a new DepartmentsHandler
func NewDepartmentsHandler(deptRepo *repositories.DepartmentRepository) *DepartmentsHandler {
	return &DepartmentsHandler{deptRepo: deptRepo}
}

// GetAll handles GET /departments
func (h *DepartmentsHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	departments, err := h.deptRepo.GetAllDepartments(r.Context())
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

	department, err := h.deptRepo.GetDepartmentByID(r.Context(), id)
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

	newDepartment, err := h.deptRepo.CreateDepartment(r.Context(), department)
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
	updatedDepartment, err := h.deptRepo.UpdateDepartment(r.Context(), department)
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

	err = h.deptRepo.DeleteDepartment(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
