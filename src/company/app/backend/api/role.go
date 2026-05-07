// Package api provides HTTP API handlers for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.6-35b-a3b
// Percentage written by Agent: 100%

package api

import (
	"encoding/json"
	"net/http"

	"cc520-company-dbms-project/src/company/app/backend/repositories"
)

// RoleHandler handles role-related HTTP requests
type RoleHandler struct {
	repo *repositories.RoleRepository
}

// NewRoleHandler creates a new RoleHandler
func NewRoleHandler(repo *repositories.RoleRepository) *RoleHandler {
	return &RoleHandler{repo: repo}
}

// GetAll handles GET /roles
func (h *RoleHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := h.repo.GetAllRoles(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}
