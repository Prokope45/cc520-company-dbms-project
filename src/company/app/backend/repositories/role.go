// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.6-35b-a3b
// Percentage written by Agent: 100%

package repositories

import (
	"context"

	"cc520-company-dbms-project/src/company/app/backend/models"
)

// RoleRepository handles data access for Role entities
type RoleRepository struct {
	executor Executor
}

// NewRoleRepository creates a new RoleRepository instance
func NewRoleRepository(executor Executor) *RoleRepository {
	return &RoleRepository{
		executor: executor,
	}
}

// GetAllRoles retrieves all roles from the database using the stored procedure
func (r *RoleRepository) GetAllRoles(ctx context.Context) ([]models.Role, error) {
	result := r.executor.Execute(ctx, "GetAllRoles", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var roles []models.Role
	for _, row := range result.Rows {
		role := models.Role{
			Name: row["Name"].(string),
		}
		roles = append(roles, role)
	}

	return roles, nil
}
