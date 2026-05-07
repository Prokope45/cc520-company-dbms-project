// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - Gemini 3 Pro Preview
// Percentage written by Agent: 95%

package repositories

import (
	"context"

	"cc520-company-dbms-project/src/company/app/backend/models"
)

// ReportsRepository handles data access for aggregate reporting queries
type ReportsRepository struct {
	executor Executor
}

// NewReportsRepository creates a new ReportsRepository instance
func NewReportsRepository(executor Executor) *ReportsRepository {
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
			CompanyName:    safeString(row["CompanyName"]),
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

// GetDepartmentSalary_Aggregated gets aggregated salaries by department since a given hire date
func (r *ReportsRepository) GetDepartmentSalary_Aggregated(ctx context.Context, hireDate string) ([]models.DepartmentSalary_Aggregated, error) {
	result := r.executor.Execute(ctx, "Report_DepartmentSalary_Aggregated", map[string]any{
		"HireDate": hireDate,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.DepartmentSalary_Aggregated
	for _, row := range result.Rows {
		report := models.DepartmentSalary_Aggregated{
			CompanyName:    safeString(row["CompanyName"]),
			DepartmentName: safeString(row["DepartmentName"]),
			EmployeeCount:  safeInt(row["EmployeeCount"]),
			AverageSalary:  safeFloat64(row["AverageSalary"]),
			HighestSalary:  safeFloat64(row["HighestSalary"]),
			LowestSalary:   safeFloat64(row["LowestSalary"]),
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
			CompanyName:    safeString(row["CompanyName"]),
			DepartmentName: safeString(row["DepartmentName"]),
			DateTerminated: safeString(row["DateTerminated"]),
			HourlyPay:      safeFloat64(row["HourlyPay"]),
			EmployeeName:   safeString(row["EmployeeName"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// GetTopTerminatedHourly_Aggregated gets aggregated terminated hourly employee data by department
func (r *ReportsRepository) GetTopTerminatedHourly_Aggregated(ctx context.Context, terminationDate string) ([]models.TopTerminatedHourly_Aggregated, error) {
	result := r.executor.Execute(ctx, "Report_TopTerminatedHourly_Aggregated", map[string]any{
		"TerminationDate": terminationDate,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.TopTerminatedHourly_Aggregated
	for _, row := range result.Rows {
		report := models.TopTerminatedHourly_Aggregated{
			CompanyName:            safeString(row["CompanyName"]),
			DepartmentName:         safeString(row["DepartmentName"]),
			TerminatedCount:        safeInt(row["TerminatedCount"]),
			AvgTerminatedHourlyPay: safeFloat64(row["AvgTerminatedHourlyPay"]),
			LatestTerminationDate:  safeString(row["LatestTerminationDate"]),
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
			CompanyName:     safeString(row["CompanyName"]),
			DepartmentName:  safeString(row["DepartmentName"]),
			EmployeeName:    safeString(row["EmployeeName"]),
			ManagerAssigned: safeString(row["ManagerAssigned"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// GetUnhiredWithManager_Aggregated gets aggregated unhired employee data by department
func (r *ReportsRepository) GetUnhiredWithManager_Aggregated(ctx context.Context) ([]models.UnhiredWithManager_Aggregated, error) {
	result := r.executor.Execute(ctx, "Report_UnhiredWithManager_Aggregated", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.UnhiredWithManager_Aggregated
	for _, row := range result.Rows {
		report := models.UnhiredWithManager_Aggregated{
			CompanyName:    safeString(row["CompanyName"]),
			DepartmentName: safeString(row["DepartmentName"]),
			UnhiredCount:   safeInt(row["UnhiredCount"]),
			ManagerNames:   safeString(row["ManagerNames"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}

// GetHighestPaidExecutive gets the company with the highest paid executive
func (r *ReportsRepository) GetHighestPaidExecutive(ctx context.Context) (*models.HighestPaidExecutive, error) {
	result := r.executor.Execute(ctx, "Report_HighestPaidExecutive", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Rows) == 0 {
		return nil, nil // No executive found
	}

	row := result.Rows[0]
	report := models.HighestPaidExecutive{
		CompanyName:     safeString(row["CompanyName"]),
		ExecutiveName:   safeString(row["ExecutiveName"]),
		ExecutiveSalary: safeFloat64(row["ExecutiveSalary"]),
		RoleName:        safeString(row["RoleName"]),
	}

	return &report, nil
}

// GetHighestPaidExecutive_Aggregated gets aggregated executive salary data per company
func (r *ReportsRepository) GetHighestPaidExecutive_Aggregated(ctx context.Context) ([]models.HighestPaidExecutive_Aggregated, error) {
	result := r.executor.Execute(ctx, "Report_HighestPaidExecutive_Aggregated", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var reports []models.HighestPaidExecutive_Aggregated
	for _, row := range result.Rows {
		report := models.HighestPaidExecutive_Aggregated{
			CompanyName:          safeString(row["CompanyName"]),
			ExecutiveCount:       safeInt(row["ExecutiveCount"]),
			HighestExecutiveSalary: safeFloat64(row["HighestExecutiveSalary"]),
		}
		reports = append(reports, report)
	}

	return reports, nil
}
