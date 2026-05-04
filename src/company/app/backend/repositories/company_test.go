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
	"time"

	"cc520-company-dbms-project/src/company/app/backend/models"
	"cc520-company-dbms-project/src/company/db/executor"
)

// mockExecutor implements the executor.Executor interface for testing
type mockExecutor struct {
	// expectedProcedure is the procedure name this mock was set up for
	expectedProcedure string
	// rows to return when the expected procedure is called
	rows []map[string]interface{}
	// rowsAffected to return
	rowsAffected int64
	// lastInsertID to return
	lastInsertID sql.NullInt64
	// err to return
	err error
	// callCount tracks how many times Execute was called
	callCount int
	// lastProcedure stores the last procedure name called
	lastProcedure string
}

func (m *mockExecutor) Execute(_ context.Context, procedureName string, _ map[string]interface{}) executor.Result {
	m.callCount++
	m.lastProcedure = procedureName

	if m.expectedProcedure != "" && m.expectedProcedure != procedureName {
		return executor.Result{
			Error:        errors.New("unexpected procedure: " + procedureName),
			RowsAffected: 0,
			LastInsertID: sql.NullInt64{},
			Rows:         nil,
		}
	}

	return executor.Result{
		Error:        m.err,
		RowsAffected: m.rowsAffected,
		LastInsertID: m.lastInsertID,
		Rows:         m.rows,
	}
}

func TestNewCompanyRepository(t *testing.T) {
	mock := &mockExecutor{}
	repo := NewCompanyRepository(mock)
	if repo == nil {
		t.Fatal("NewCompanyRepository returned nil")
	}
	if repo.executor != mock {
		t.Error("executor was not set correctly")
	}
}

func TestGetAllCompanies_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetAllCompanies",
		rows: []map[string]interface{}{
			{
				"CompanyID":   int64(1),
				"Name":        "Acme Corp",
				"CreatedDate": "2024-01-15T10:30:00",
			},
			{
				"CompanyID":   int64(2),
				"Name":        "Globex Inc",
				"CreatedDate": time.Date(2024, 2, 20, 8, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05"),
			},
		},
	}
	repo := NewCompanyRepository(mock)

	companies, err := repo.GetAllCompanies(context.Background())
	if err != nil {
		t.Fatalf("GetAllCompanies returned error: %v", err)
	}
	if len(companies) != 2 {
		t.Errorf("expected 2 companies, got %d", len(companies))
	}
	if companies[0].CompanyID != 1 {
		t.Errorf("expected first company ID 1, got %d", companies[0].CompanyID)
	}
	if companies[0].Name != "Acme Corp" {
		t.Errorf("expected first company name 'Acme Corp', got '%s'", companies[0].Name)
	}
	if companies[1].Name != "Globex Inc" {
		t.Errorf("expected second company name 'Globex Inc', got '%s'", companies[1].Name)
	}
	if mock.callCount != 1 {
		t.Errorf("expected Execute to be called 1 time, got %d", mock.callCount)
	}
}

func TestGetAllCompanies_EmptyResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetAllCompanies",
		rows:              []map[string]interface{}{},
	}
	repo := NewCompanyRepository(mock)

	companies, err := repo.GetAllCompanies(context.Background())
	if err != nil {
		t.Fatalf("GetAllCompanies returned error: %v", err)
	}
	if len(companies) != 0 {
		t.Errorf("expected 0 companies, got %d", len(companies))
	}
}

func TestGetAllCompanies_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetAllCompanies",
		err:               errors.New("database connection failed"),
	}
	repo := NewCompanyRepository(mock)

	companies, err := repo.GetAllCompanies(context.Background())
	if err == nil {
		t.Fatal("expected error from GetAllCompanies, got nil")
	}
	if companies != nil {
		t.Errorf("expected nil companies on error, got %v", companies)
	}
}

func TestGetCompanyByID_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetCompanyByID",
		rows: []map[string]interface{}{
			{
				"CompanyID":   int64(42),
				"Name":        "Initech",
				"CreatedDate": "2023-06-01T00:00:00",
			},
		},
	}
	repo := NewCompanyRepository(mock)

	company, err := repo.GetCompanyByID(context.Background(), 42)
	if err != nil {
		t.Fatalf("GetCompanyByID returned error: %v", err)
	}
	if company == nil {
		t.Fatal("expected non-nil company")
	}
	if company.CompanyID != 42 {
		t.Errorf("expected company ID 42, got %d", company.CompanyID)
	}
	if company.Name != "Initech" {
		t.Errorf("expected company name 'Initech', got '%s'", company.Name)
	}
}

func TestGetCompanyByID_NotFound(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetCompanyByID",
		rows:              []map[string]interface{}{},
	}
	repo := NewCompanyRepository(mock)

	company, err := repo.GetCompanyByID(context.Background(), 999)
	if err != sql.ErrNoRows {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
	if company != nil {
		t.Error("expected nil company for not found")
	}
}

func TestGetCompanyByID_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetCompanyByID",
		err:               errors.New("query timeout"),
	}
	repo := NewCompanyRepository(mock)

	company, err := repo.GetCompanyByID(context.Background(), 1)
	if err == nil {
		t.Fatal("expected error from GetCompanyByID, got nil")
	}
	if company != nil {
		t.Error("expected nil company on error")
	}
}

func TestCreateCompany_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateCompany",
		rows: []map[string]interface{}{
			{"CompanyID": int64(100)},
		},
	}
	repo := NewCompanyRepository(mock)

	company, err := repo.CreateCompany(context.Background(), models.Company{Name: "Umbrella Corp"})
	if err != nil {
		t.Fatalf("CreateCompany returned error: %v", err)
	}
	if company == nil {
		t.Fatal("expected non-nil company")
	}
	if company.CompanyID != 100 {
		t.Errorf("expected company ID 100, got %d", company.CompanyID)
	}
	if company.Name != "Umbrella Corp" {
		t.Errorf("expected company name 'Umbrella Corp', got '%s'", company.Name)
	}
}

func TestCreateCompany_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateCompany",
		err:               errors.New("duplicate name"),
	}
	repo := NewCompanyRepository(mock)

	company, err := repo.CreateCompany(context.Background(), models.Company{Name: "Dup Corp"})
	if err == nil {
		t.Fatal("expected error from CreateCompany, got nil")
	}
	if company != nil {
		t.Error("expected nil company on error")
	}
}

func TestUpdateCompany_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "UpdateCompany",
		rows:              []map[string]interface{}{},
	}
	repo := NewCompanyRepository(mock)

	company := models.Company{CompanyID: 1, Name: "Updated Corp"}
	result, err := repo.UpdateCompany(context.Background(), company)
	if err != nil {
		t.Fatalf("UpdateCompany returned error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil company")
	}
	if result.CompanyID != 1 {
		t.Errorf("expected company ID 1, got %d", result.CompanyID)
	}
	if result.Name != "Updated Corp" {
		t.Errorf("expected company name 'Updated Corp', got '%s'", result.Name)
	}
}

func TestUpdateCompany_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "UpdateCompany",
		err:               errors.New("company not found"),
	}
	repo := NewCompanyRepository(mock)

	company := models.Company{CompanyID: 1, Name: "Nope Corp"}
	result, err := repo.UpdateCompany(context.Background(), company)
	if err == nil {
		t.Fatal("expected error from UpdateCompany, got nil")
	}
	if result != nil {
		t.Error("expected nil result on error")
	}
}

func TestDeleteCompany_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "DeleteCompany",
		rows:              []map[string]interface{}{},
	}
	repo := NewCompanyRepository(mock)

	err := repo.DeleteCompany(context.Background(), 5)
	if err != nil {
		t.Fatalf("DeleteCompany returned error: %v", err)
	}
}

func TestDeleteCompany_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "DeleteCompany",
		err:               errors.New("foreign key constraint"),
	}
	repo := NewCompanyRepository(mock)

	err := repo.DeleteCompany(context.Background(), 5)
	if err == nil {
		t.Fatal("expected error from DeleteCompany, got nil")
	}
}

func TestParseDateTime_String(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseDateTime("2024-01-15T10:30:00")
	if result != "2024-01-15T10:30:00" {
		t.Errorf("expected '2024-01-15T10:30:00', got '%s'", result)
	}
}

func TestParseDateTime_Time(t *testing.T) {
	repo := &CompanyRepository{}
	expected := time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC)
	result := repo.parseDateTime(expected)
	if result != "2024-06-15T12:30:45" {
		t.Errorf("expected '2024-06-15T12:30:45', got '%s'", result)
	}
}

func TestParseDateTime_Int64(t *testing.T) {
	repo := &CompanyRepository{}
	// Unix timestamp for 2024-01-01T00:00:00 UTC
	ts := int64(1704067200)
	result := repo.parseDateTime(ts)
	if result != "2024-01-01T00:00:00" {
		t.Errorf("expected '2024-01-01T00:00:00', got '%s'", result)
	}
}

func TestParseDateTime_Int32(t *testing.T) {
	repo := &CompanyRepository{}
	ts := int32(1704067200)
	result := repo.parseDateTime(ts)
	if result != "2024-01-01T00:00:00" {
		t.Errorf("expected '2024-01-01T00:00:00', got '%s'", result)
	}
}

func TestParseDateTime_Default(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseDateTime(123.456)
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

func TestParseCompanyID_Int64(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseCompanyID(int64(42))
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestParseCompanyID_Int32(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseCompanyID(int32(99))
	if result != 99 {
		t.Errorf("expected 99, got %d", result)
	}
}

func TestParseCompanyID_String(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseCompanyID("12345")
	if result != 12345 {
		t.Errorf("expected 12345, got %d", result)
	}
}

func TestParseCompanyID_StringInvalid(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseCompanyID("not-a-number")
	if result != 0 {
		t.Errorf("expected 0 for invalid string, got %d", result)
	}
}

func TestParseCompanyID_Default(t *testing.T) {
	repo := &CompanyRepository{}
	result := repo.parseCompanyID(nil)
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestCompanyRepository_ProcedureNamePrefix(t *testing.T) {
	testCases := []struct {
		procedure   string
		name        string
		call        func(*CompanyRepository, *mockExecutor)
		expectCount int
		rows        []map[string]interface{}
	}{
		{
			procedure: "GetAllCompanies",
			name:      "GetAllCompanies",
			call: func(r *CompanyRepository, m *mockExecutor) {
				_, _ = r.GetAllCompanies(context.Background())
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"CompanyID": int64(1), "Name": "Test", "CreatedDate": "2024-01-01T00:00:00"}},
		},
		{
			procedure: "GetCompanyByID",
			name:      "GetCompanyByID",
			call: func(r *CompanyRepository, m *mockExecutor) {
				_, _ = r.GetCompanyByID(context.Background(), 1)
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"CompanyID": int64(1), "Name": "Test", "CreatedDate": "2024-01-01T00:00:00"}},
		},
		{
			procedure: "CreateCompany",
			name:      "CreateCompany",
			call: func(r *CompanyRepository, m *mockExecutor) {
				_, _ = r.CreateCompany(context.Background(), models.Company{})
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"CompanyID": int64(1)}},
		},
		{
			procedure: "UpdateCompany",
			name:      "UpdateCompany",
			call: func(r *CompanyRepository, m *mockExecutor) {
				_, _ = r.UpdateCompany(context.Background(), models.Company{})
			},
			expectCount: 1,
			rows:        []map[string]interface{}{},
		},
		{
			procedure: "DeleteCompany",
			name:      "DeleteCompany",
			call: func(r *CompanyRepository, m *mockExecutor) {
				_ = r.DeleteCompany(context.Background(), 1)
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
			repo := NewCompanyRepository(mock)

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
