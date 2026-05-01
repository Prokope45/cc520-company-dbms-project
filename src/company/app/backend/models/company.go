// Package models defines the data structures for the company application
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 95%

package models

// Company represents a company in the database
type Company struct {
	CompanyID   int64  `json:"company_id"`
	Name        string `json:"name"`
	CreatedDate string `json:"created_date"`
}
