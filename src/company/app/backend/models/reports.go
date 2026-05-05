// Package models defines the data structures for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - Gemini 3 Pro Preview
// Percentage written by Agent: 95%

package models

// DepartmentSalaryRank represents the rank of a salaried employee within their department
type DepartmentSalaryRank struct {
	CompanyName    string  `json:"company_name"`
	DepartmentName string  `json:"department_name"`
	HireMonth      int     `json:"hire_month"`
	HireYear       int     `json:"hire_year"`
	EmployeeName   string  `json:"employee_name"`
	EmployeeSalary float64 `json:"employee_salary"`
	SalaryRank     int64   `json:"salary_rank"`
}

// TopTerminatedHourly represents a high-earning hourly employee who recently left
type TopTerminatedHourly struct {
	DateTerminated string  `json:"date_terminated"`
	HourlyPay      float64 `json:"hourly_pay"`
	EmployeeName   string  `json:"employee_name"`
}

// UnhiredWithManager represents an employee record that has a manager but no hire date
type UnhiredWithManager struct {
	EmployeeName    string `json:"employee_name"`
	ManagerAssigned string `json:"manager_assigned"`
}

// HighestPaidCEO represents the company and name of the highest paid CEO
type HighestPaidCEO struct {
	CompanyName string  `json:"company_name"`
	CEOName     string  `json:"ceo_name"`
	CEOSalary   float64 `json:"ceo_salary"`
}
