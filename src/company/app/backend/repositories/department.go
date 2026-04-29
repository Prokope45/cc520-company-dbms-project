// Package repositories provides data access layer for the company application
package repositories

import (
	"context"

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
