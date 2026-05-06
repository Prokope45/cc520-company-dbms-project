// Package api provides HTTP handlers for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - Gemini 3 Pro Preview
// Percentage written by Agent: 100%

package api

import (
	"encoding/json"
	"net/http"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/app/backend/repositories"
)

// ReportsHandler handles HTTP requests for aggregate reports
type ReportsHandler struct {
	repo *repositories.ReportsRepository
}

// NewReportsHandler creates a new ReportsHandler instance
func NewReportsHandler(repo *repositories.ReportsRepository) *ReportsHandler {
	return &ReportsHandler{
		repo: repo,
	}
}

// GetDepartmentSalaryRanks handles GET /reports/department-salary-ranks
// Query params: hireDate (required)
func (h *ReportsHandler) GetDepartmentSalaryRanks(w http.ResponseWriter, r *http.Request) {
	hireDate := r.URL.Query().Get("hireDate")
	if hireDate == "" {
		http.Error(w, "hireDate query parameter is required", http.StatusBadRequest)
		return
	}

	reports, err := h.repo.GetDepartmentSalaryRanks(r.Context(), hireDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.DepartmentSalaryRank, 0) // Return empty array instead of null
	}
	json.NewEncoder(w).Encode(reports)
}

// GetDepartmentSalary_Aggregated handles GET /reports/department-salary-ranks-aggregated
// Query params: hireDate (required)
func (h *ReportsHandler) GetDepartmentSalary_Aggregated(w http.ResponseWriter, r *http.Request) {
	hireDate := r.URL.Query().Get("hireDate")
	if hireDate == "" {
		http.Error(w, "hireDate query parameter is required", http.StatusBadRequest)
		return
	}

	reports, err := h.repo.GetDepartmentSalary_Aggregated(r.Context(), hireDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.DepartmentSalary_Aggregated, 0)
	}
	json.NewEncoder(w).Encode(reports)
}

// GetTopTerminatedHourly handles GET /reports/top-terminated-hourly
// Query params: terminationDate (required)
func (h *ReportsHandler) GetTopTerminatedHourly(w http.ResponseWriter, r *http.Request) {
	terminationDate := r.URL.Query().Get("terminationDate")
	if terminationDate == "" {
		http.Error(w, "terminationDate query parameter is required", http.StatusBadRequest)
		return
	}

	reports, err := h.repo.GetTopTerminatedHourly(r.Context(), terminationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.TopTerminatedHourly, 0)
	}
	json.NewEncoder(w).Encode(reports)
}

// GetTopTerminatedHourly_Aggregated handles GET /reports/top-terminated-hourly-aggregated
// Query params: terminationDate (required)
func (h *ReportsHandler) GetTopTerminatedHourly_Aggregated(w http.ResponseWriter, r *http.Request) {
	terminationDate := r.URL.Query().Get("terminationDate")
	if terminationDate == "" {
		http.Error(w, "terminationDate query parameter is required", http.StatusBadRequest)
		return
	}

	reports, err := h.repo.GetTopTerminatedHourly_Aggregated(r.Context(), terminationDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.TopTerminatedHourly_Aggregated, 0)
	}
	json.NewEncoder(w).Encode(reports)
}

// GetUnhiredWithManager handles GET /reports/unhired-with-manager
func (h *ReportsHandler) GetUnhiredWithManager(w http.ResponseWriter, r *http.Request) {
	reports, err := h.repo.GetUnhiredWithManager(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.UnhiredWithManager, 0)
	}
	json.NewEncoder(w).Encode(reports)
}

// GetUnhiredWithManager_Aggregated handles GET /reports/unhired-with-manager-aggregated
func (h *ReportsHandler) GetUnhiredWithManager_Aggregated(w http.ResponseWriter, r *http.Request) {
	reports, err := h.repo.GetUnhiredWithManager_Aggregated(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.UnhiredWithManager_Aggregated, 0)
	}
	json.NewEncoder(w).Encode(reports)
}

// GetHighestPaidCEO handles GET /reports/highest-paid-ceo
func (h *ReportsHandler) GetHighestPaidCEO(w http.ResponseWriter, r *http.Request) {
	report, err := h.repo.GetHighestPaidCEO(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if report == nil {
		http.Error(w, "No CEO found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// GetHighestPaidCEO_Aggregated handles GET /reports/highest-paid-ceo-aggregated
func (h *ReportsHandler) GetHighestPaidCEO_Aggregated(w http.ResponseWriter, r *http.Request) {
	reports, err := h.repo.GetHighestPaidCEO_Aggregated(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if reports == nil {
		reports = make([]models.HighestPaidCEO_Aggregated, 0)
	}
	json.NewEncoder(w).Encode(reports)
}
