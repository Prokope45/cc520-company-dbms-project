// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package repositories

import (
	"context"
	"database/sql"
	"strconv"
	"time"

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
			CompanyID:   r.parseCompanyID(row["CompanyID"]),
			Name:        row["Name"].(string),
			CreatedDate: r.parseDateTime(row["CreatedDate"]),
		}
		companies = append(companies, company)
	}

	return companies, nil
}

// GetCompanyByID retrieves a company by ID
func (r *CompanyRepository) GetCompanyByID(ctx context.Context, id int64) (*models.Company, error) {
	result := r.executor.Execute(ctx, "GetCompanyByID", map[string]any{
		"id": id,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Rows) == 0 {
		return nil, sql.ErrNoRows
	}

	company := models.Company{
		CompanyID:   r.parseCompanyID(result.Rows[0]["CompanyID"]),
		Name:        result.Rows[0]["Name"].(string),
		CreatedDate: r.parseDateTime(result.Rows[0]["CreatedDate"]),
	}

	return &company, nil
}

// CreateCompany creates a new company
func (r *CompanyRepository) CreateCompany(ctx context.Context, company models.Company) (*models.Company, error) {
	result := r.executor.Execute(ctx, "CreateCompany", map[string]any{
		"Name": company.Name,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	company.CompanyID = r.parseCompanyID(result.Rows[0]["CompanyID"])

	return &company, nil
}

// UpdateCompany updates an existing company
func (r *CompanyRepository) UpdateCompany(ctx context.Context, company models.Company) (*models.Company, error) {
	result := r.executor.Execute(ctx, "UpdateCompany", map[string]any{
		"CompanyID": company.CompanyID,
		"Name":      company.Name,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return &company, nil
}

// DeleteCompany deletes a company by ID
func (r *CompanyRepository) DeleteCompany(ctx context.Context, id int64) error {
	result := r.executor.Execute(ctx, "DeleteCompany", map[string]any{
		"CompanyID": id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// parseDateTime handles both string and time.Time types from the database
func (r *CompanyRepository) parseDateTime(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case time.Time:
		return v.Format("2006-01-02T15:04:05")
	case int64:
		return time.Unix(v, 0).Format("2006-01-02T15:04:05")
	case int32:
		return time.Unix(int64(v), 0).Format("2006-01-02T15:04:05")
	default:
		return ""
	}
}

// parseCompanyID handles both int64 and string types from SCOPE_IDENTITY()
func (r *CompanyRepository) parseCompanyID(val interface{}) int64 {
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
