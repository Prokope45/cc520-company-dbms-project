// Package models defines the data structures for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package models

// Department represents a department in the database
type Department struct {
	DepartmentID int64  `json:"department_id"`
	CompanyID    int64  `json:"company_id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
}
