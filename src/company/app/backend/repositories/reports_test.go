// Authors:
//   - Jared Paubel
//   - OpenCode agent - local qwen/qwen3.6-35b-a3b
//
// Percentage written by Agent: 100%
package repositories

import (
	"context"
	"errors"
	"testing"
)

func TestNewReportsRepository(t *testing.T) {
	mock := &mockExecutor{}
	repo := NewReportsRepository(mock)
	if repo == nil {
		t.Fatal("NewReportsRepository returned nil")
	}
	if repo.executor != mock {
		t.Error("executor was not set correctly")
	}
}

func TestGetDepartmentSalaryRanks_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_DepartmentSalaryRanks",
		rows: []map[string]interface{}{
			{
				"DepartmentName": "Engineering",
				"HireMonth":      3,
				"HireYear":       2024,
				"EmployeeName":   "John Doe",
				"EmployeeSalary": float64(85000),
				"SalaryRank":     int64(1),
			},
			{
				"DepartmentName": "Engineering",
				"HireMonth":      6,
				"HireYear":       2024,
				"EmployeeName":   "Jane Smith",
				"EmployeeSalary": float64(75000),
				"SalaryRank":     int64(2),
			},
		},
	}
	repo := NewReportsRepository(mock)

	ranks, err := repo.GetDepartmentSalaryRanks(context.Background(), "2024-01-01")
	if err != nil {
		t.Fatalf("GetDepartmentSalaryRanks returned error: %v", err)
	}
	if len(ranks) != 2 {
		t.Errorf("expected 2 ranks, got %d", len(ranks))
	}
	if ranks[0].DepartmentName != "Engineering" {
		t.Errorf("expected department 'Engineering', got '%s'", ranks[0].DepartmentName)
	}
	if ranks[0].HireMonth != 3 {
		t.Errorf("expected hire month 3, got %d", ranks[0].HireMonth)
	}
	if ranks[0].HireYear != 2024 {
		t.Errorf("expected hire year 2024, got %d", ranks[0].HireYear)
	}
	if ranks[0].EmployeeName != "John Doe" {
		t.Errorf("expected employee 'John Doe', got '%s'", ranks[0].EmployeeName)
	}
	if ranks[0].EmployeeSalary != 85000 {
		t.Errorf("expected salary 85000, got %f", ranks[0].EmployeeSalary)
	}
	if ranks[0].SalaryRank != 1 {
		t.Errorf("expected rank 1, got %d", ranks[0].SalaryRank)
	}
	if mock.callCount != 1 {
		t.Errorf("expected Execute to be called 1 time, got %d", mock.callCount)
	}
}

func TestGetDepartmentSalaryRanks_EmptyResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_DepartmentSalaryRanks",
		rows:              []map[string]interface{}{},
	}
	repo := NewReportsRepository(mock)

	ranks, err := repo.GetDepartmentSalaryRanks(context.Background(), "2024-01-01")
	if err != nil {
		t.Fatalf("GetDepartmentSalaryRanks returned error: %v", err)
	}
	if len(ranks) != 0 {
		t.Errorf("expected 0 ranks, got %d", len(ranks))
	}
}

func TestGetDepartmentSalaryRanks_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_DepartmentSalaryRanks",
		err:               errors.New("query timeout"),
	}
	repo := NewReportsRepository(mock)

	ranks, err := repo.GetDepartmentSalaryRanks(context.Background(), "2024-01-01")
	if err == nil {
		t.Fatal("expected error from GetDepartmentSalaryRanks, got nil")
	}
	if ranks != nil {
		t.Errorf("expected nil ranks on error, got %v", ranks)
	}
}

func TestGetTopTerminatedHourly_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_TopTerminatedHourly",
		rows: []map[string]interface{}{
			{
				"DateTerminated": "2024-03-15",
				"HourlyPay":      float64(55.0),
				"EmployeeName":   "Bob Williams",
			},
			{
				"DateTerminated": "2024-02-28",
				"HourlyPay":      float64(48.0),
				"EmployeeName":   "Carol Davis",
			},
		},
	}
	repo := NewReportsRepository(mock)

	reports, err := repo.GetTopTerminatedHourly(context.Background(), "2024-01-01")
	if err != nil {
		t.Fatalf("GetTopTerminatedHourly returned error: %v", err)
	}
	if len(reports) != 2 {
		t.Errorf("expected 2 reports, got %d", len(reports))
	}
	if reports[0].EmployeeName != "Bob Williams" {
		t.Errorf("expected employee 'Bob Williams', got '%s'", reports[0].EmployeeName)
	}
	if reports[0].HourlyPay != 55.0 {
		t.Errorf("expected hourly pay 55, got %f", reports[0].HourlyPay)
	}
	if reports[0].DateTerminated != "2024-03-15" {
		t.Errorf("expected date '2024-03-15', got '%s'", reports[0].DateTerminated)
	}
	if mock.callCount != 1 {
		t.Errorf("expected Execute to be called 1 time, got %d", mock.callCount)
	}
}

func TestGetTopTerminatedHourly_EmptyResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_TopTerminatedHourly",
		rows:              []map[string]interface{}{},
	}
	repo := NewReportsRepository(mock)

	reports, err := repo.GetTopTerminatedHourly(context.Background(), "2024-01-01")
	if err != nil {
		t.Fatalf("GetTopTerminatedHourly returned error: %v", err)
	}
	if len(reports) != 0 {
		t.Errorf("expected 0 reports, got %d", len(reports))
	}
}

func TestGetTopTerminatedHourly_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_TopTerminatedHourly",
		err:               errors.New("database error"),
	}
	repo := NewReportsRepository(mock)

	reports, err := repo.GetTopTerminatedHourly(context.Background(), "2024-01-01")
	if err == nil {
		t.Fatal("expected error from GetTopTerminatedHourly, got nil")
	}
	if reports != nil {
		t.Errorf("expected nil reports on error, got %v", reports)
	}
}

func TestGetUnhiredWithManager_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_UnhiredWithManager",
		rows: []map[string]interface{}{
			{
				"EmployeeName":    "Alice Johnson",
				"ManagerAssigned": "Bob Smith",
			},
			{
				"EmployeeName":    "Charlie Brown",
				"ManagerAssigned": "Diana Prince",
			},
		},
	}
	repo := NewReportsRepository(mock)

	reports, err := repo.GetUnhiredWithManager(context.Background())
	if err != nil {
		t.Fatalf("GetUnhiredWithManager returned error: %v", err)
	}
	if len(reports) != 2 {
		t.Errorf("expected 2 reports, got %d", len(reports))
	}
	if reports[0].EmployeeName != "Alice Johnson" {
		t.Errorf("expected employee 'Alice Johnson', got '%s'", reports[0].EmployeeName)
	}
	if reports[0].ManagerAssigned != "Bob Smith" {
		t.Errorf("expected manager 'Bob Smith', got '%s'", reports[0].ManagerAssigned)
	}
	if reports[1].EmployeeName != "Charlie Brown" {
		t.Errorf("expected employee 'Charlie Brown', got '%s'", reports[1].EmployeeName)
	}
	if mock.callCount != 1 {
		t.Errorf("expected Execute to be called 1 time, got %d", mock.callCount)
	}
}

func TestGetUnhiredWithManager_EmptyResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_UnhiredWithManager",
		rows:              []map[string]interface{}{},
	}
	repo := NewReportsRepository(mock)

	reports, err := repo.GetUnhiredWithManager(context.Background())
	if err != nil {
		t.Fatalf("GetUnhiredWithManager returned error: %v", err)
	}
	if len(reports) != 0 {
		t.Errorf("expected 0 reports, got %d", len(reports))
	}
}

func TestGetUnhiredWithManager_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_UnhiredWithManager",
		err:               errors.New("query failed"),
	}
	repo := NewReportsRepository(mock)

	reports, err := repo.GetUnhiredWithManager(context.Background())
	if err == nil {
		t.Fatal("expected error from GetUnhiredWithManager, got nil")
	}
	if reports != nil {
		t.Errorf("expected nil reports on error, got %v", reports)
	}
}

func TestGetHighestPaidCEO_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_HighestPaidCEO",
		rows: []map[string]interface{}{
			{
				"CompanyName": "TechCorp",
				"CEOName":     "John CEO",
				"CEOSalary":   float64(500000),
			},
		},
	}
	repo := NewReportsRepository(mock)

	report, err := repo.GetHighestPaidCEO(context.Background())
	if err != nil {
		t.Fatalf("GetHighestPaidCEO returned error: %v", err)
	}
	if report == nil {
		t.Fatal("expected non-nil report")
	}
	if report.CompanyName != "TechCorp" {
		t.Errorf("expected company 'TechCorp', got '%s'", report.CompanyName)
	}
	if report.CEOName != "John CEO" {
		t.Errorf("expected CEO 'John CEO', got '%s'", report.CEOName)
	}
	if report.CEOSalary != 500000 {
		t.Errorf("expected salary 500000, got %f", report.CEOSalary)
	}
}

func TestGetHighestPaidCEO_NoResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_HighestPaidCEO",
		rows:              []map[string]interface{}{},
	}
	repo := NewReportsRepository(mock)

	report, err := repo.GetHighestPaidCEO(context.Background())
	if err != nil {
		t.Fatalf("GetHighestPaidCEO returned error: %v", err)
	}
	if report != nil {
		t.Error("expected nil report for no result")
	}
}

func TestGetHighestPaidCEO_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "Report_HighestPaidCEO",
		err:               errors.New("database error"),
	}
	repo := NewReportsRepository(mock)

	report, err := repo.GetHighestPaidCEO(context.Background())
	if err == nil {
		t.Fatal("expected error from GetHighestPaidCEO, got nil")
	}
	if report != nil {
		t.Error("expected nil report on error")
	}
}

func TestReportsRepository_ProcedureNamePrefix(t *testing.T) {
	testCases := []struct {
		procedure   string
		name        string
		call        func(*ReportsRepository, *mockExecutor)
		expectCount int
		rows        []map[string]interface{}
	}{
		{
			procedure: "Report_DepartmentSalaryRanks",
			name:      "GetDepartmentSalaryRanks",
			call: func(r *ReportsRepository, m *mockExecutor) {
				_, _ = r.GetDepartmentSalaryRanks(context.Background(), "2024-01-01")
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"DepartmentName": "Engineering", "HireMonth": 3, "HireYear": 2024, "EmployeeName": "Test", "EmployeeSalary": 85000.0, "SalaryRank": int64(1)}},
		},
		{
			procedure: "Report_TopTerminatedHourly",
			name:      "GetTopTerminatedHourly",
			call: func(r *ReportsRepository, m *mockExecutor) {
				_, _ = r.GetTopTerminatedHourly(context.Background(), "2024-01-01")
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"DateTerminated": "2024-01-01", "HourlyPay": 50.0, "EmployeeName": "Test"}},
		},
		{
			procedure: "Report_UnhiredWithManager",
			name:      "GetUnhiredWithManager",
			call: func(r *ReportsRepository, m *mockExecutor) {
				_, _ = r.GetUnhiredWithManager(context.Background())
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"EmployeeName": "Test", "ManagerAssigned": "Manager"}},
		},
		{
			procedure: "Report_HighestPaidCEO",
			name:      "GetHighestPaidCEO",
			call: func(r *ReportsRepository, m *mockExecutor) {
				_, _ = r.GetHighestPaidCEO(context.Background())
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"CompanyName": "Test", "CEOName": "CEO", "CEOSalary": 100000.0}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockExecutor{
				expectedProcedure: tc.procedure,
				rows:              tc.rows,
			}
			repo := NewReportsRepository(mock)

			tc.call(repo, mock)

			if mock.callCount != tc.expectCount {
				t.Errorf("expected %d calls, got %d", tc.expectCount, mock.callCount)
			}
			if mock.lastProcedure != tc.procedure {
				t.Errorf("expected last procedure '%s', got '%s'", tc.procedure, mock.lastProcedure)
			}
		})
	}
}
