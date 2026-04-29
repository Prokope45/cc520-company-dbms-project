// Package backend provides the HTTP API for the company application
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 90%

package backend

import (
	"encoding/json"
	"net/http"
	"sync"

	"cc520-company-dbms-project/src/company/app/backend/repositories"
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
