// Package repositories provides data access layer for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 100%

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"cc520-company-dbms-project/src/company/app/backend/models"
)

// EmployeeRepository handles data access for Employee entities.
type EmployeeRepository struct {
	executor Executor
}

// NewEmployeeRepository creates a new EmployeeRepository instance
func NewEmployeeRepository(executor Executor) *EmployeeRepository {
	return &EmployeeRepository{
		executor: executor,
	}
}

// GetAllEmployees retrieves all employees
func (r *EmployeeRepository) GetAllEmployees(ctx context.Context) ([]models.Employee, error) {
	result := r.executor.Execute(ctx, "GetEmployeeProfiles", nil)
	if result.Error != nil {
		return nil, result.Error
	}

	var employees []models.Employee
	for _, row := range result.Rows {
		employees = append(employees, r.mapRowToEmployee(row))
	}

	return employees, nil
}

// GetEmployeeByID retrieves an employee by ID
func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int64) (*models.Employee, error) {
	result := r.executor.Execute(ctx, "GetEmployeeProfileByID", map[string]any{
		"EmployeeID": id,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Rows) == 0 {
		return nil, sql.ErrNoRows
	}

	emp := r.mapRowToEmployee(result.Rows[0])
	return &emp, nil
}

// CreateEmployee creates a new employee profile
func (r *EmployeeRepository) CreateEmployee(ctx context.Context, emp models.Employee) (*models.Employee, error) {
	params := map[string]any{
		"CompanyID":      emp.CompanyID,
		"DepartmentID":   emp.DepartmentID,
		"FirstName":      emp.FirstName,
		"LastName":       emp.LastName,
		"Email":          emp.Email,
		"Phone":          emp.Phone,
		"Street":         emp.Street,
		"AddressLineTwo": emp.AddressLineTwo,
		"City":           emp.City,
		"State":          emp.State,
		"ZipCode":        emp.ZipCode,
		"RoleTitle":      emp.RoleTitle,
		"StatusType":     emp.StatusType,
	}

	if emp.ManagerID != nil {
		params["ManagerID"] = *emp.ManagerID
	}
	if emp.HireDate != "" {
		params["HireDate"] = emp.HireDate
	}
	if emp.StatusType == "Hourly" {
		params["HourlyPay"] = emp.HourlyPay
		params["MaxHoursPerWeek"] = emp.MaxHoursPerWeek
	} else if emp.StatusType == "Salary" {
		params["BaseSalary"] = emp.BaseSalary
		params["Bonus"] = emp.Bonus
		params["Deductions"] = emp.Deductions
		params["PaidTimeOffHours"] = emp.PaidTimeOffHours
		params["SickHours"] = emp.SickHours
		if emp.EffectiveFrom != "" {
			params["EffectiveFrom"] = emp.EffectiveFrom
		}
		if emp.EffectiveTo != "" {
			params["EffectiveTo"] = emp.EffectiveTo
		}
	}

	result := r.executor.Execute(ctx, "CreateEmployeeProfile", params)
	if result.Error != nil {
		return nil, result.Error
	}

	if len(result.Rows) > 0 {
		emp.EmployeeID = r.parseEmployeeID(result.Rows[0]["EmployeeID"])
	}

	return &emp, nil
}

// UpdateEmployee updates an existing employee profile
func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, emp models.Employee) (*models.Employee, error) {
	params := map[string]any{
		"EmployeeID":     emp.EmployeeID,
		"DepartmentID":   emp.DepartmentID,
		"FirstName":      emp.FirstName,
		"LastName":       emp.LastName,
		"Email":          emp.Email,
		"Phone":          emp.Phone,
		"Street":         emp.Street,
		"AddressLineTwo": emp.AddressLineTwo,
		"City":           emp.City,
		"State":          emp.State,
		"ZipCode":        emp.ZipCode,
		"RoleTitle":      emp.RoleTitle,
		"StatusType":     emp.StatusType,
	}
	if emp.ManagerID != nil {
		params["ManagerID"] = *emp.ManagerID
	}
	if emp.StatusType == "Hourly" {
		params["HourlyPay"] = emp.HourlyPay
		params["MaxHoursPerWeek"] = emp.MaxHoursPerWeek
	} else if emp.StatusType == "Salary" {
		params["BaseSalary"] = emp.BaseSalary
		params["Bonus"] = emp.Bonus
		params["Deductions"] = emp.Deductions
		params["PaidTimeOffHours"] = emp.PaidTimeOffHours
		params["SickHours"] = emp.SickHours
	}

	result := r.executor.Execute(ctx, "UpdateEmployeeProfile", params)
	if result.Error != nil {
		return nil, result.Error
	}

	return &emp, nil
}

// DeleteEmployee deletes an employee profile by ID
func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id int64) error {
	result := r.executor.Execute(ctx, "DeleteEmployeeProfile", map[string]any{
		"EmployeeID": id,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Helper methods to safely parse SQL nulls
func safeString(val interface{}) string {
	if val == nil {
		return ""
	}
	if s, ok := val.(string); ok {
		return s
	}
	return ""
}

func safeInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int64:
		return v
	case int32:
		return int64(v)
	case int:
		return int64(v)
	case float64:
		return int64(v)
	}
	return 0
}

func safeInt64Ptr(val interface{}) *int64 {
	if val == nil {
		return nil
	}
	v := safeInt64(val)
	return &v
}

func safeInt(val interface{}) int {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int64:
		return int(v)
	case int32:
		return int(v)
	case int:
		return v
	case float64:
		return int(v)
	}
	return 0
}

func safeFloat64(val interface{}) float64 {
	if val == nil {
		return 0.0
	}
	switch v := val.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	// Some DB drivers return decimals as byte arrays or strings
	case string:
		var f float64
		// Best effort parsing if we hit a string for decimal
		fmt.Sscanf(v, "%f", &f)
		return f
	case []byte:
		var f float64
		fmt.Sscanf(string(v), "%f", &f)
		return f
	}
	return 0.0
}

func (r *EmployeeRepository) mapRowToEmployee(row map[string]interface{}) models.Employee {
	return models.Employee{
		EmployeeID:       safeInt64(row["EmployeeID"]),
		CompanyID:        safeInt64(row["CompanyID"]),
		DepartmentID:     safeInt64(row["DepartmentID"]),
		Department:       safeString(row["Department"]),
		FirstName:        safeString(row["FirstName"]),
		LastName:         safeString(row["LastName"]),
		Email:            safeString(row["Email"]),
		Phone:            safeString(row["Phone"]),
		Street:           safeString(row["Street"]),
		AddressLineTwo:   safeString(row["AddressLineTwo"]),
		City:             safeString(row["City"]),
		State:            safeString(row["State"]),
		ZipCode:          safeString(row["ZipCode"]),
		RoleTitle:        safeString(row["RoleTitle"]),
		ManagerID:        safeInt64Ptr(row["ManagerID"]),
		ManagerName:      safeString(row["ManagerName"]),
		HireDate:         safeString(row["HireDate"]),
		TerminationDate:  safeString(row["TerminationDate"]),
		StatusType:       safeString(row["StatusType"]),
		HourlyPay:        safeFloat64(row["HourlyPay"]),
		MaxHoursPerWeek:  safeInt(row["MaxHoursPerWeek"]),
		BaseSalary:       safeFloat64(row["BaseSalary"]),
		Bonus:            safeFloat64(row["Bonus"]),
		Deductions:       safeFloat64(row["Deductions"]),
		PaidTimeOffHours: safeInt(row["PaidTimeOffHours"]),
		SickHours:        safeInt(row["SickHours"]),
		EffectiveFrom:    safeString(row["EffectiveFrom"]),
		EffectiveTo:      safeString(row["EffectiveTo"]),
	}
}

// parseEmployeeID handles both int64 and string types from SCOPE_IDENTITY()
func (r *EmployeeRepository) parseEmployeeID(val interface{}) int64 {
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
