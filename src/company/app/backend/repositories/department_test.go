// Authors:
//   - Jared Paubel
//   - OpenCode agent - local qwen/qwen3.6-35b-a3b
//
// Percentage written by Agent: 100%
package repositories

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"cc520-company-dbms-project/src/company/app/backend/models"
)

func TestNewDepartmentRepository(t *testing.T) {
	mock := &mockExecutor{}
	repo := NewDepartmentRepository(mock)
	if repo == nil {
		t.Fatal("NewDepartmentRepository returned nil")
	}
	if repo.executor != mock {
		t.Error("executor was not set correctly")
	}
}

func TestGetAllDepartments_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetAllDepartments",
		rows: []map[string]interface{}{
			{
				"DepartmentID": int64(1),
				"CompanyID":    int64(1),
				"Name":         "Engineering",
				"Description":  "Software development",
			},
			{
				"DepartmentID": int64(2),
				"CompanyID":    int64(1),
				"Name":         "Marketing",
				"Description":  "Product marketing",
			},
		},
	}
	repo := NewDepartmentRepository(mock)

	departments, err := repo.GetAllDepartments(context.Background())
	if err != nil {
		t.Fatalf("GetAllDepartments returned error: %v", err)
	}
	if len(departments) != 2 {
		t.Errorf("expected 2 departments, got %d", len(departments))
	}
	if departments[0].DepartmentID != 1 {
		t.Errorf("expected first department ID 1, got %d", departments[0].DepartmentID)
	}
	if departments[0].Name != "Engineering" {
		t.Errorf("expected first department name 'Engineering', got '%s'", departments[0].Name)
	}
	if departments[0].CompanyID != 1 {
		t.Errorf("expected first department company ID 1, got %d", departments[0].CompanyID)
	}
	if departments[1].Name != "Marketing" {
		t.Errorf("expected second department name 'Marketing', got '%s'", departments[1].Name)
	}
	if mock.callCount != 1 {
		t.Errorf("expected Execute to be called 1 time, got %d", mock.callCount)
	}
}

func TestGetAllDepartments_EmptyResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetAllDepartments",
		rows:              []map[string]interface{}{},
	}
	repo := NewDepartmentRepository(mock)

	departments, err := repo.GetAllDepartments(context.Background())
	if err != nil {
		t.Fatalf("GetAllDepartments returned error: %v", err)
	}
	if len(departments) != 0 {
		t.Errorf("expected 0 departments, got %d", len(departments))
	}
}

func TestGetAllDepartments_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetAllDepartments",
		err:               errors.New("database connection failed"),
	}
	repo := NewDepartmentRepository(mock)

	departments, err := repo.GetAllDepartments(context.Background())
	if err == nil {
		t.Fatal("expected error from GetAllDepartments, got nil")
	}
	if departments != nil {
		t.Errorf("expected nil departments on error, got %v", departments)
	}
}

func TestGetDepartmentByID_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetDepartmentByID",
		rows: []map[string]interface{}{
			{
				"DepartmentID": int64(42),
				"CompanyID":    int64(1),
				"Name":         "Sales",
				"Description":  "Revenue generation",
			},
		},
	}
	repo := NewDepartmentRepository(mock)

	department, err := repo.GetDepartmentByID(context.Background(), 42)
	if err != nil {
		t.Fatalf("GetDepartmentByID returned error: %v", err)
	}
	if department == nil {
		t.Fatal("expected non-nil department")
	}
	if department.DepartmentID != 42 {
		t.Errorf("expected department ID 42, got %d", department.DepartmentID)
	}
	if department.Name != "Sales" {
		t.Errorf("expected department name 'Sales', got '%s'", department.Name)
	}
	if department.CompanyID != 1 {
		t.Errorf("expected company ID 1, got %d", department.CompanyID)
	}
	if department.Description != "Revenue generation" {
		t.Errorf("expected description 'Revenue generation', got '%s'", department.Description)
	}
}

func TestGetDepartmentByID_NotFound(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetDepartmentByID",
		rows:              []map[string]interface{}{},
	}
	repo := NewDepartmentRepository(mock)

	department, err := repo.GetDepartmentByID(context.Background(), 999)
	if err != sql.ErrNoRows {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
	if department != nil {
		t.Error("expected nil department for not found")
	}
}

func TestGetDepartmentByID_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetDepartmentByID",
		err:               errors.New("query timeout"),
	}
	repo := NewDepartmentRepository(mock)

	department, err := repo.GetDepartmentByID(context.Background(), 1)
	if err == nil {
		t.Fatal("expected error from GetDepartmentByID, got nil")
	}
	if department != nil {
		t.Error("expected nil department on error")
	}
}

func TestCreateDepartment_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateDepartment",
		rows: []map[string]interface{}{
			{"DepartmentID": int64(100)},
		},
	}
	repo := NewDepartmentRepository(mock)

	department := models.Department{CompanyID: 1, Name: "HR", Description: "Human resources"}
	result, err := repo.CreateDepartment(context.Background(), department)
	if err != nil {
		t.Fatalf("CreateDepartment returned error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil department")
	}
	if result.DepartmentID != 100 {
		t.Errorf("expected department ID 100, got %d", result.DepartmentID)
	}
	if result.Name != "HR" {
		t.Errorf("expected department name 'HR', got '%s'", result.Name)
	}
	if result.CompanyID != 1 {
		t.Errorf("expected company ID 1, got %d", result.CompanyID)
	}
}

func TestCreateDepartment_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateDepartment",
		err:               errors.New("duplicate name"),
	}
	repo := NewDepartmentRepository(mock)

	department := models.Department{CompanyID: 1, Name: "Dup Dept"}
	result, err := repo.CreateDepartment(context.Background(), department)
	if err == nil {
		t.Fatal("expected error from CreateDepartment, got nil")
	}
	if result != nil {
		t.Error("expected nil result on error")
	}
}

func TestUpdateDepartment_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "UpdateDepartment",
		rows:              []map[string]interface{}{},
	}
	repo := NewDepartmentRepository(mock)

	department := models.Department{DepartmentID: 1, Name: "Updated Dept", Description: "Updated description"}
	result, err := repo.UpdateDepartment(context.Background(), department)
	if err != nil {
		t.Fatalf("UpdateDepartment returned error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil department")
	}
	if result.DepartmentID != 1 {
		t.Errorf("expected department ID 1, got %d", result.DepartmentID)
	}
	if result.Name != "Updated Dept" {
		t.Errorf("expected department name 'Updated Dept', got '%s'", result.Name)
	}
	if result.Description != "Updated description" {
		t.Errorf("expected description 'Updated description', got '%s'", result.Description)
	}
}

func TestUpdateDepartment_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "UpdateDepartment",
		err:               errors.New("department not found"),
	}
	repo := NewDepartmentRepository(mock)

	department := models.Department{DepartmentID: 1, Name: "Nope Dept"}
	result, err := repo.UpdateDepartment(context.Background(), department)
	if err == nil {
		t.Fatal("expected error from UpdateDepartment, got nil")
	}
	if result != nil {
		t.Error("expected nil result on error")
	}
}

func TestDeleteDepartment_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "DeleteDepartment",
		rows:              []map[string]interface{}{},
	}
	repo := NewDepartmentRepository(mock)

	err := repo.DeleteDepartment(context.Background(), 5)
	if err != nil {
		t.Fatalf("DeleteDepartment returned error: %v", err)
	}
}

func TestDeleteDepartment_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "DeleteDepartment",
		err:               errors.New("foreign key constraint"),
	}
	repo := NewDepartmentRepository(mock)

	err := repo.DeleteDepartment(context.Background(), 5)
	if err == nil {
		t.Fatal("expected error from DeleteDepartment, got nil")
	}
}

func TestParseDepartmentID_Int64(t *testing.T) {
	repo := &DepartmentRepository{}
	result := repo.parseDepartmentID(int64(42))
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestParseDepartmentID_Int32(t *testing.T) {
	repo := &DepartmentRepository{}
	result := repo.parseDepartmentID(int32(99))
	if result != 99 {
		t.Errorf("expected 99, got %d", result)
	}
}

func TestParseDepartmentID_String(t *testing.T) {
	repo := &DepartmentRepository{}
	result := repo.parseDepartmentID("12345")
	if result != 12345 {
		t.Errorf("expected 12345, got %d", result)
	}
}

func TestParseDepartmentID_StringInvalid(t *testing.T) {
	repo := &DepartmentRepository{}
	result := repo.parseDepartmentID("not-a-number")
	if result != 0 {
		t.Errorf("expected 0 for invalid string, got %d", result)
	}
}

func TestParseDepartmentID_Default(t *testing.T) {
	repo := &DepartmentRepository{}
	result := repo.parseDepartmentID(nil)
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestDepartmentRepository_ProcedureNamePrefix(t *testing.T) {
	testCases := []struct {
		procedure   string
		name        string
		call        func(*DepartmentRepository, *mockExecutor)
		expectCount int
		rows        []map[string]interface{}
	}{
		{
			procedure: "GetAllDepartments",
			name:      "GetAllDepartments",
			call: func(r *DepartmentRepository, m *mockExecutor) {
				_, _ = r.GetAllDepartments(context.Background())
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"DepartmentID": int64(1), "CompanyID": int64(1), "Name": "Test", "Description": "Test desc"}},
		},
		{
			procedure: "GetDepartmentByID",
			name:      "GetDepartmentByID",
			call: func(r *DepartmentRepository, m *mockExecutor) {
				_, _ = r.GetDepartmentByID(context.Background(), 1)
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"DepartmentID": int64(1), "CompanyID": int64(1), "Name": "Test", "Description": "Test desc"}},
		},
		{
			procedure: "CreateDepartment",
			name:      "CreateDepartment",
			call: func(r *DepartmentRepository, m *mockExecutor) {
				_, _ = r.CreateDepartment(context.Background(), models.Department{})
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"DepartmentID": int64(1)}},
		},
		{
			procedure: "UpdateDepartment",
			name:      "UpdateDepartment",
			call: func(r *DepartmentRepository, m *mockExecutor) {
				_, _ = r.UpdateDepartment(context.Background(), models.Department{})
			},
			expectCount: 1,
			rows:        []map[string]interface{}{},
		},
		{
			procedure: "DeleteDepartment",
			name:      "DeleteDepartment",
			call: func(r *DepartmentRepository, m *mockExecutor) {
				_ = r.DeleteDepartment(context.Background(), 1)
			},
			expectCount: 1,
			rows:        []map[string]interface{}{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockExecutor{
				expectedProcedure: tc.procedure,
				rows:              tc.rows,
			}
			repo := NewDepartmentRepository(mock)

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
