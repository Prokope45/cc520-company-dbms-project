// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package repositories

import (
	"context"
	"database/sql"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/db/executor"
)

// DepartmentRepository handles data access for Department entities
type DepartmentRepository struct {
	executor *executor.Executor
}

// NewDepartmentRepository creates a new DepartmentRepository instance
func NewDepartmentRepository(executor *executor.Executor) *DepartmentRepository {
	return &DepartmentRepository{
		executor: executor,
	}
}

// GetAllDepartments retrieves all departments from the database using the stored procedure
func (r *DepartmentRepository) GetAllDepartments(ctx context.Context) ([]models.Department, error) {
	result := r.executor.Execute(ctx, "GetAllDepartments", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var departments []models.Department
	for _, row := range result.Rows {
		department := models.Department{
			DepartmentID: int64(row["DepartmentID"].(int64)),
			CompanyID:    int64(row["CompanyID"].(int64)),
			Name:         row["Name"].(string),
			Description:  row["Description"].(string),
		}
		departments = append(departments, department)
	}

	return departments, nil
}

// GetDepartmentByID retrieves a department by ID
func (r *DepartmentRepository) GetDepartmentByID(ctx context.Context, id int64) (*models.Department, error) {
	result := r.executor.Execute(ctx, "GetDepartmentByID", map[string]any{
		"id": id,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Rows) == 0 {
		return nil, sql.ErrNoRows
	}

	department := models.Department{
		DepartmentID: int64(result.Rows[0]["DepartmentID"].(int64)),
		CompanyID:    int64(result.Rows[0]["CompanyID"].(int64)),
		Name:         result.Rows[0]["Name"].(string),
		Description:  result.Rows[0]["Description"].(string),
	}

	return &department, nil
}

// CreateDepartment creates a new department
func (r *DepartmentRepository) CreateDepartment(ctx context.Context, department models.Department) (*models.Department, error) {
	result := r.executor.Execute(ctx, "CreateDepartment", map[string]any{
		"company_id":  department.CompanyID,
		"name":        department.Name,
		"description": department.Description,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	department.DepartmentID = int64(result.Rows[0]["DepartmentID"].(int64))

	return &department, nil
}

// UpdateDepartment updates an existing department
func (r *DepartmentRepository) UpdateDepartment(ctx context.Context, department models.Department) (*models.Department, error) {
	result := r.executor.Execute(ctx, "UpdateDepartment", map[string]any{
		"id":          department.DepartmentID,
		"name":        department.Name,
		"description": department.Description,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return &department, nil
}

// DeleteDepartment deletes a department by ID
func (r *DepartmentRepository) DeleteDepartment(ctx context.Context, id int64) error {
	result := r.executor.Execute(ctx, "DeleteDepartment", map[string]any{
		"id": id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
