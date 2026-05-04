// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package repositories

import (
	"context"
	"strconv"
	"database/sql"

	"cc520-company-dbms-project/src/company/app/backend/models"
)

// DepartmentRepository handles data access for Department entities
type DepartmentRepository struct {
	executor Executor
}

// NewDepartmentRepository creates a new DepartmentRepository instance
func NewDepartmentRepository(executor Executor) *DepartmentRepository {
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
			DepartmentID: r.parseDepartmentID(row["DepartmentID"]),
			CompanyID:    r.parseDepartmentID(row["CompanyID"]),
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
		DepartmentID: r.parseDepartmentID(result.Rows[0]["DepartmentID"]),
		CompanyID:    r.parseDepartmentID(result.Rows[0]["CompanyID"]),
		Name:         result.Rows[0]["Name"].(string),
		Description:  result.Rows[0]["Description"].(string),
	}

	return &department, nil
}

// CreateDepartment creates a new department
func (r *DepartmentRepository) CreateDepartment(ctx context.Context, department models.Department) (*models.Department, error) {
	result := r.executor.Execute(ctx, "CreateDepartment", map[string]any{
		"CompanyID":   department.CompanyID,
		"Name":        department.Name,
		"Description": department.Description,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	department.DepartmentID = r.parseDepartmentID(result.Rows[0]["DepartmentID"])

	return &department, nil
}

// UpdateDepartment updates an existing department
func (r *DepartmentRepository) UpdateDepartment(ctx context.Context, department models.Department) (*models.Department, error) {
	result := r.executor.Execute(ctx, "UpdateDepartment", map[string]any{
		"DepartmentID": department.DepartmentID,
		"Name":         department.Name,
		"Description":  department.Description,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return &department, nil
}

// DeleteDepartment deletes a department by ID
func (r *DepartmentRepository) DeleteDepartment(ctx context.Context, id int64) error {
	result := r.executor.Execute(ctx, "DeleteDepartment", map[string]any{
		"DepartmentID": id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}


// parseDepartmentID handles both int64 and string types from SCOPE_IDENTITY()
func (r *DepartmentRepository) parseDepartmentID(val interface{}) int64 {
	switch v := val.(type) {
	case int64:
		return v
	case int32:
		return int64(v)
	case string:
		id, _ := strconv.ParseInt(v, 10, 64)
		return id
	default:
		return 0
	}
}
