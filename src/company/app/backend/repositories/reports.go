// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - Gemini 3 Pro Preview
// Percentage written by Agent: 100%

package repositories

import (
	"context"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/db/executor"
)

// ReportsRepository handles data access for aggregate reporting queries
type ReportsRepository struct {
	executor *executor.Executor
}

// NewReportsRepository creates a new ReportsRepository instance
func NewReportsRepository(executor *executor.Executor) *ReportsRepository {
	return &ReportsRepository{
		executor: executor,
	}
}

// GetDepartmentSalaryRanks gets ranked salaries by department since a given hire date
func (r *ReportsRepository) GetDepartmentSalaryRanks(ctx context.Context, hireDate string) ([]models.DepartmentSalaryRank, error) {
	result := r.executor.Execute(ctx, "Report_DepartmentSalaryRanks", map[string]any{
		"HireDate": hireDate,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.DepartmentSalaryRank
	for _, row := range result.Rows {
		report := models.DepartmentSalaryRank{
			DepartmentName: safeString(row["DepartmentName"]),
			HireMonth:      safeInt(row["HireMonth"]),
			HireYear:       safeInt(row["HireYear"]),
			EmployeeName:   safeString(row["EmployeeName"]),
			EmployeeSalary: safeFloat64(row["EmployeeSalary"]),
			SalaryRank:     safeInt64(row["SalaryRank"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// GetTopTerminatedHourly gets the top 5 highest paid hourly employees who left since a date
func (r *ReportsRepository) GetTopTerminatedHourly(ctx context.Context, terminationDate string) ([]models.TopTerminatedHourly, error) {
	result := r.executor.Execute(ctx, "Report_TopTerminatedHourly", map[string]any{
		"TerminationDate": terminationDate,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.TopTerminatedHourly
	for _, row := range result.Rows {
		report := models.TopTerminatedHourly{
			DateTerminated: safeString(row["DateTerminated"]),
			HourlyPay:      safeFloat64(row["HourlyPay"]),
			EmployeeName:   safeString(row["EmployeeName"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// GetUnhiredWithManager gets employees with an assigned manager but no hire date
func (r *ReportsRepository) GetUnhiredWithManager(ctx context.Context) ([]models.UnhiredWithManager, error) {
	result := r.executor.Execute(ctx, "Report_UnhiredWithManager", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.UnhiredWithManager
	for _, row := range result.Rows {
		report := models.UnhiredWithManager{
			EmployeeName:    safeString(row["EmployeeName"]),
			ManagerAssigned: safeString(row["ManagerAssigned"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// GetHighestPaidCEO gets the company with the highest paid CEO
func (r *ReportsRepository) GetHighestPaidCEO(ctx context.Context) (*models.HighestPaidCEO, error) {
	result := r.executor.Execute(ctx, "Report_HighestPaidCEO", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Rows) == 0 {
		return nil, nil // No CEO found
	}

	row := result.Rows[0]
	report := models.HighestPaidCEO{
		CompanyName: safeString(row["CompanyName"]),
		CEOName:     safeString(row["CEOName"]),
		CEOSalary:   safeFloat64(row["CEOSalary"]),
	}

	return &report, nil
}
