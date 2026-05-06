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

// DepartmentSalary_Aggregated represents aggregated department salary data
type DepartmentSalary_Aggregated struct {
	CompanyName    string  `json:"company_name"`
	DepartmentName string  `json:"department_name"`
	EmployeeCount  int     `json:"employee_count"`
	AverageSalary  float64 `json:"average_salary"`
	HighestSalary  float64 `json:"highest_salary"`
	LowestSalary   float64 `json:"lowest_salary"`
}

// TopTerminatedHourly represents a high-earning hourly employee who recently left
type TopTerminatedHourly struct {
	CompanyName    string  `json:"company_name"`
	DepartmentName string  `json:"department_name"`
	DateTerminated string  `json:"date_terminated"`
	HourlyPay      float64 `json:"hourly_pay"`
	EmployeeName   string  `json:"employee_name"`
}

// TopTerminatedHourly_Aggregated represents aggregated terminated hourly employee data by department
type TopTerminatedHourly_Aggregated struct {
	CompanyName            string  `json:"company_name"`
	DepartmentName         string  `json:"department_name"`
	TerminatedCount        int     `json:"terminated_count"`
	AvgTerminatedHourlyPay float64 `json:"avg_terminated_hourly_pay"`
	LatestTerminationDate  string  `json:"latest_termination_date"`
}

// UnhiredWithManager represents an employee record that has a manager but no hire date
type UnhiredWithManager struct {
	CompanyName     string `json:"company_name"`
	DepartmentName  string `json:"department_name"`
	EmployeeName    string `json:"employee_name"`
	ManagerAssigned string `json:"manager_assigned"`
}

// UnhiredWithManager_Aggregated represents aggregated unhired employee data by department
type UnhiredWithManager_Aggregated struct {
	CompanyName    string `json:"company_name"`
	DepartmentName string `json:"department_name"`
	UnhiredCount   int    `json:"unhired_count"`
	ManagerNames   string `json:"manager_names"`
}

// HighestPaidCEO represents the company and name of the highest paid CEO
type HighestPaidCEO struct {
	CompanyName string  `json:"company_name"`
	CEOName     string  `json:"ceo_name"`
	CEOSalary   float64 `json:"ceo_salary"`
}

// HighestPaidCEO_Aggregated represents aggregated CEO salary data per company
type HighestPaidCEO_Aggregated struct {
	CompanyName      string  `json:"company_name"`
	CEOCount         int     `json:"ceo_count"`
	HighestCEOSalary float64 `json:"highest_ceo_salary"`
}
