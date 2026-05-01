// Package models defines the data structures for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package models

// Employee represents a combined view of an employee,
// including their person data, address, role, department, and status.
type Employee struct {
	EmployeeID   int64  `json:"employee_id"`
	CompanyID    int64  `json:"company_id"`
	DepartmentID int64  `json:"department_id"`
	Department   string `json:"department"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`

	// Person Data
	Email string `json:"email"`
	Phone string `json:"phone"`

	// Address data
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`

	// Organization & Role Data
	RoleTitle       string `json:"role_title"`
	ManagerName     string `json:"manager_name"`
	HireDate        string `json:"hire_date"`
	TerminationDate string `json:"termination_date"`

	// Employee Status data
	StatusType string `json:"status_type"` // "Salary" or "Hourly"

	// Hourly Details
	HourlyPay       float64 `json:"hourly_pay"`
	MaxHoursPerWeek int     `json:"max_hours_per_week"`

	// Salary Details
	BaseSalary       float64 `json:"base_salary"`
	Bonus            float64 `json:"bonus"`
	Deductions       float64 `json:"deductions"`
	PaidTimeOffHours int     `json:"paid_time_off_hours"`
	SickHours        int     `json:"sick_hours"`
	EffectiveFrom    string  `json:"effective_from"`
	EffectiveTo      string  `json:"effective_to"`
}
