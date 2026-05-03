// Package api provides HTTP API handlers for the company application
// Errors provide helper functions to transform the error message into a readable format.
// Authors:
// 	- Jared Paubel
//  - OpenCode agent - Gemini 3 Pro Preview
// Percentage written by Agent: 100%

package api

import (
	"strings"
)

// Map constraint names to user-friendly messages
var constraintMessages = map[string]string{
	"CK_HourlyPay_Positive":                 "Hourly wage must be greater than 0.",
	"CK_Employee_HireDate_NotNull_Salaried": "Hire date must be provided.",
	"CK_Employee_TerminationDate_NotNull":   "Termination date must be after hire date.",
	"CK_Employee_Manager_NotSelf":           "An employee cannot be their own manager.",
	"CK_Salary_BaseSalary_Positive":         "Base salary cannot be negative.",
	"CK_Salary_Bonus_Positive":              "Bonus cannot be negative.",
	"CK_Salary_Deductions_Positive":         "Deductions cannot be negative.",
	"CK_Salary_EffectiveDates":              "Effective 'From' date must be before 'To' date.",
	"CK_Person_Email_Length":                "Email address cannot be empty.",
	"CK_Company_Name_Length":                "Company name cannot be empty.",
	"CK_Role_Name_Length":                   "Role name cannot be empty.",
}

// FormatDBError converts raw database errors into user-friendly messages
func FormatDBError(err error) string {
	errStr := err.Error()

	// Check for known constraints
	for constraint, msg := range constraintMessages {
		if strings.Contains(errStr, constraint) {
			return msg
		}
	}

	// Check for unique constraints
	if strings.Contains(errStr, "Violation of UNIQUE KEY constraint") {
		if strings.Contains(errStr, "UK_Employee_PersonID") {
			return "This person is already registered as an employee."
		}
		return "A record with this information already exists."
	}

	// Check for foreign key constraints
	if strings.Contains(errStr, "The DELETE statement conflicted with the REFERENCE constraint") {
		return "Cannot delete this record because it is currently being referenced by other records."
	}

	// Return original error if not recognized
	return errStr
}
