// Package repositories provides data access layer for the company application
package repositories

import (
	"context"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/db/executor"
)

// CompanyRepository handles data access for Company entities
type CompanyRepository struct {
	executor *executor.Executor
}

// NewCompanyRepository creates a new CompanyRepository instance
func NewCompanyRepository(executor *executor.Executor) *CompanyRepository {
	return &CompanyRepository{
		executor: executor,
	}
}

// GetAllCompanies retrieves all companies from the database using the stored procedure
func (r *CompanyRepository) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	result := r.executor.Execute(ctx, "GetAllCompanies", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var companies []models.Company
	for _, row := range result.Rows {
		company := models.Company{
			CompanyID:   int64(row["CompanyID"].(int64)),
			Name:        row["Name"].(string),
			CreatedDate: row["CreatedDate"].(string),
		}
		companies = append(companies, company)
	}

	return companies, nil
}
