// Package models defines the data structures for the company application
package models

// Company represents a company in the database
type Company struct {
	CompanyID   int64     `json:"company_id"`
	Name        string    `json:"name"`
	CreatedDate string    `json:"created_date"`
}
