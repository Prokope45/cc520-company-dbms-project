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

func TestNewEmployeeRepository(t *testing.T) {
	mock := &mockExecutor{}
	repo := NewEmployeeRepository(mock)
	if repo == nil {
		t.Fatal("NewEmployeeRepository returned nil")
	}
	if repo.executor != mock {
		t.Error("executor was not set correctly")
	}
}

func TestGetAllEmployees_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetEmployeeProfiles",
		rows: []map[string]interface{}{
			{
				"EmployeeID":       int64(1),
				"CompanyID":        int64(1),
				"DepartmentID":     int64(1),
				"Department":       "Engineering",
				"FirstName":        "John",
				"LastName":         "Doe",
				"Email":            "john@example.com",
				"Phone":            "555-1234",
				"Street":           "123 Main St",
				"AddressLineTwo":   "Apt 4",
				"City":             "Springfield",
				"State":            "IL",
				"ZipCode":          "62701",
				"RoleTitle":        "Engineer",
				"ManagerID":        int64(2),
				"ManagerName":      "Jane Smith",
				"HireDate":         "2023-01-15",
				"TerminationDate":  "",
				"StatusType":       "Salary",
				"HourlyPay":        0.0,
				"MaxHoursPerWeek":  0,
				"BaseSalary":       85000.0,
				"Bonus":            5000.0,
				"Deductions":       2000.0,
				"PaidTimeOffHours": 120,
				"SickHours":        80,
				"EffectiveFrom":    "2023-01-01",
				"EffectiveTo":      "",
			},
		},
	}
	repo := NewEmployeeRepository(mock)

	employees, err := repo.GetAllEmployees(context.Background())
	if err != nil {
		t.Fatalf("GetAllEmployees returned error: %v", err)
	}
	if len(employees) != 1 {
		t.Errorf("expected 1 employee, got %d", len(employees))
	}
	emp := employees[0]
	if emp.EmployeeID != 1 {
		t.Errorf("expected employee ID 1, got %d", emp.EmployeeID)
	}
	if emp.FirstName != "John" {
		t.Errorf("expected first name 'John', got '%s'", emp.FirstName)
	}
	if emp.LastName != "Doe" {
		t.Errorf("expected last name 'Doe', got '%s'", emp.LastName)
	}
	if emp.Email != "john@example.com" {
		t.Errorf("expected email 'john@example.com', got '%s'", emp.Email)
	}
	if emp.Department != "Engineering" {
		t.Errorf("expected department 'Engineering', got '%s'", emp.Department)
	}
	if emp.StatusType != "Salary" {
		t.Errorf("expected status 'Salary', got '%s'", emp.StatusType)
	}
	if emp.BaseSalary != 85000.0 {
		t.Errorf("expected base salary 85000, got %f", emp.BaseSalary)
	}
	if emp.ManagerID == nil {
		t.Error("expected ManagerID to be non-nil")
	} else if *emp.ManagerID != 2 {
		t.Errorf("expected ManagerID 2, got %d", *emp.ManagerID)
	}
}

func TestGetAllEmployees_EmptyResult(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetEmployeeProfiles",
		rows:              []map[string]interface{}{},
	}
	repo := NewEmployeeRepository(mock)

	employees, err := repo.GetAllEmployees(context.Background())
	if err != nil {
		t.Fatalf("GetAllEmployees returned error: %v", err)
	}
	if len(employees) != 0 {
		t.Errorf("expected 0 employees, got %d", len(employees))
	}
}

func TestGetAllEmployees_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetEmployeeProfiles",
		err:               errors.New("database connection failed"),
	}
	repo := NewEmployeeRepository(mock)

	employees, err := repo.GetAllEmployees(context.Background())
	if err == nil {
		t.Fatal("expected error from GetAllEmployees, got nil")
	}
	if employees != nil {
		t.Errorf("expected nil employees on error, got %v", employees)
	}
}

func TestGetEmployeeByID_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetEmployeeProfileByID",
		rows: []map[string]interface{}{
			{
				"EmployeeID":       int64(42),
				"CompanyID":        int64(1),
				"DepartmentID":     int64(1),
				"Department":       "Sales",
				"FirstName":        "Alice",
				"LastName":         "Johnson",
				"Email":            "alice@example.com",
				"Phone":            "555-5678",
				"Street":           "456 Oak Ave",
				"AddressLineTwo":   "",
				"City":             "Portland",
				"State":            "OR",
				"ZipCode":          "97201",
				"RoleTitle":        "Sales Rep",
				"ManagerID":        nil,
				"ManagerName":      "",
				"HireDate":         "2022-06-01",
				"TerminationDate":  "",
				"StatusType":       "Hourly",
				"HourlyPay":        35.0,
				"MaxHoursPerWeek":  40,
				"BaseSalary":       0.0,
				"Bonus":            0.0,
				"Deductions":       1500.0,
				"PaidTimeOffHours": 80,
				"SickHours":        60,
				"EffectiveFrom":    "",
				"EffectiveTo":      "",
			},
		},
	}
	repo := NewEmployeeRepository(mock)

	emp, err := repo.GetEmployeeByID(context.Background(), 42)
	if err != nil {
		t.Fatalf("GetEmployeeByID returned error: %v", err)
	}
	if emp == nil {
		t.Fatal("expected non-nil employee")
	}
	if emp.EmployeeID != 42 {
		t.Errorf("expected employee ID 42, got %d", emp.EmployeeID)
	}
	if emp.FirstName != "Alice" {
		t.Errorf("expected first name 'Alice', got '%s'", emp.FirstName)
	}
	if emp.LastName != "Johnson" {
		t.Errorf("expected last name 'Johnson', got '%s'", emp.LastName)
	}
	if emp.Email != "alice@example.com" {
		t.Errorf("expected email 'alice@example.com', got '%s'", emp.Email)
	}
	if emp.StatusType != "Hourly" {
		t.Errorf("expected status 'Hourly', got '%s'", emp.StatusType)
	}
	if emp.HourlyPay != 35.0 {
		t.Errorf("expected hourly pay 35, got %f", emp.HourlyPay)
	}
	if emp.ManagerID != nil {
		t.Error("expected ManagerID to be nil")
	}
}

func TestGetEmployeeByID_NotFound(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetEmployeeProfileByID",
		rows:              []map[string]interface{}{},
	}
	repo := NewEmployeeRepository(mock)

	emp, err := repo.GetEmployeeByID(context.Background(), 999)
	if err != sql.ErrNoRows {
		t.Fatalf("expected sql.ErrNoRows, got %v", err)
	}
	if emp != nil {
		t.Error("expected nil employee for not found")
	}
}

func TestGetEmployeeByID_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "GetEmployeeProfileByID",
		err:               errors.New("query timeout"),
	}
	repo := NewEmployeeRepository(mock)

	emp, err := repo.GetEmployeeByID(context.Background(), 1)
	if err == nil {
		t.Fatal("expected error from GetEmployeeByID, got nil")
	}
	if emp != nil {
		t.Error("expected nil employee on error")
	}
}

func TestCreateEmployee_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateEmployeeProfile",
		rows: []map[string]interface{}{
			{"EmployeeID": int64(200)},
		},
	}
	repo := NewEmployeeRepository(mock)

	managerID := int64(5)
	emp := models.Employee{
		CompanyID:    1,
		DepartmentID: 1,
		FirstName:    "Bob",
		LastName:     "Williams",
		Email:        "bob@example.com",
		Phone:        "555-9999",
		Street:       "789 Pine Rd",
		City:         "Denver",
		State:        "CO",
		ZipCode:      "80201",
		RoleTitle:    "Developer",
		StatusType:   "Salary",
		ManagerID:    &managerID,
		HireDate:     "2024-03-01",
		BaseSalary:   90000.0,
		Bonus:        8000.0,
	}
	result, err := repo.CreateEmployee(context.Background(), emp)
	if err != nil {
		t.Fatalf("CreateEmployee returned error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil employee")
	}
	if result.EmployeeID != 200 {
		t.Errorf("expected employee ID 200, got %d", result.EmployeeID)
	}
	if result.FirstName != "Bob" {
		t.Errorf("expected first name 'Bob', got '%s'", result.FirstName)
	}
}

func TestCreateEmployee_Hourly(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateEmployeeProfile",
		rows: []map[string]interface{}{
			{"EmployeeID": int64(201)},
		},
	}
	repo := NewEmployeeRepository(mock)

	emp := models.Employee{
		CompanyID:       1,
		DepartmentID:    1,
		FirstName:       "Carol",
		LastName:        "Davis",
		Email:           "carol@example.com",
		StatusType:      "Hourly",
		HourlyPay:       40.0,
		MaxHoursPerWeek: 35,
	}
	result, err := repo.CreateEmployee(context.Background(), emp)
	if err != nil {
		t.Fatalf("CreateEmployee (hourly) returned error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil employee")
	}
	if result.EmployeeID != 201 {
		t.Errorf("expected employee ID 201, got %d", result.EmployeeID)
	}
}

func TestCreateEmployee_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "CreateEmployeeProfile",
		err:               errors.New("duplicate email"),
	}
	repo := NewEmployeeRepository(mock)

	emp := models.Employee{FirstName: "Eve", LastName: "Wilson", Email: "eve@example.com"}
	result, err := repo.CreateEmployee(context.Background(), emp)
	if err == nil {
		t.Fatal("expected error from CreateEmployee, got nil")
	}
	if result != nil {
		t.Error("expected nil result on error")
	}
}

func TestUpdateEmployee_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "UpdateEmployeeProfile",
		rows:              []map[string]interface{}{},
	}
	repo := NewEmployeeRepository(mock)

	managerID := int64(10)
	emp := models.Employee{
		EmployeeID:   1,
		DepartmentID: 2,
		FirstName:    "Frank",
		LastName:     "Miller",
		Email:        "frank@example.com",
		StatusType:   "Salary",
		BaseSalary:   95000.0,
		Bonus:        10000.0,
		ManagerID:    &managerID,
	}
	result, err := repo.UpdateEmployee(context.Background(), emp)
	if err != nil {
		t.Fatalf("UpdateEmployee returned error: %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil employee")
	}
	if result.EmployeeID != 1 {
		t.Errorf("expected employee ID 1, got %d", result.EmployeeID)
	}
	if result.FirstName != "Frank" {
		t.Errorf("expected first name 'Frank', got '%s'", result.FirstName)
	}
}

func TestUpdateEmployee_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "UpdateEmployeeProfile",
		err:               errors.New("employee not found"),
	}
	repo := NewEmployeeRepository(mock)

	emp := models.Employee{EmployeeID: 1, FirstName: "Nope"}
	result, err := repo.UpdateEmployee(context.Background(), emp)
	if err == nil {
		t.Fatal("expected error from UpdateEmployee, got nil")
	}
	if result != nil {
		t.Error("expected nil result on error")
	}
}

func TestDeleteEmployee_Success(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "DeleteEmployeeProfile",
		rows:              []map[string]interface{}{},
	}
	repo := NewEmployeeRepository(mock)

	err := repo.DeleteEmployee(context.Background(), 5)
	if err != nil {
		t.Fatalf("DeleteEmployee returned error: %v", err)
	}
}

func TestDeleteEmployee_Error(t *testing.T) {
	mock := &mockExecutor{
		expectedProcedure: "DeleteEmployeeProfile",
		err:               errors.New("foreign key constraint"),
	}
	repo := NewEmployeeRepository(mock)

	err := repo.DeleteEmployee(context.Background(), 5)
	if err == nil {
		t.Fatal("expected error from DeleteEmployee, got nil")
	}
}

func TestMapRowToEmployee_AllFields(t *testing.T) {
	repo := &EmployeeRepository{}
	row := map[string]interface{}{
		"EmployeeID":       int64(1),
		"CompanyID":        int64(1),
		"DepartmentID":     int64(1),
		"Department":       "Engineering",
		"FirstName":        "John",
		"LastName":         "Doe",
		"Email":            "john@example.com",
		"Phone":            "555-1234",
		"Street":           "123 Main St",
		"AddressLineTwo":   "Apt 4",
		"City":             "Springfield",
		"State":            "IL",
		"ZipCode":          "62701",
		"RoleTitle":        "Engineer",
		"ManagerID":        int64(2),
		"ManagerName":      "Jane Smith",
		"HireDate":         "2023-01-15",
		"TerminationDate":  "",
		"StatusType":       "Salary",
		"HourlyPay":        0.0,
		"MaxHoursPerWeek":  0,
		"BaseSalary":       85000.0,
		"Bonus":            5000.0,
		"Deductions":       2000.0,
		"PaidTimeOffHours": 120,
		"SickHours":        80,
		"EffectiveFrom":    "2023-01-01",
		"EffectiveTo":      "",
	}

	emp := repo.mapRowToEmployee(row)
	if emp.EmployeeID != 1 {
		t.Errorf("expected EmployeeID 1, got %d", emp.EmployeeID)
	}
	if emp.CompanyID != 1 {
		t.Errorf("expected CompanyID 1, got %d", emp.CompanyID)
	}
	if emp.DepartmentID != 1 {
		t.Errorf("expected DepartmentID 1, got %d", emp.DepartmentID)
	}
	if emp.Department != "Engineering" {
		t.Errorf("expected Department 'Engineering', got '%s'", emp.Department)
	}
	if emp.FirstName != "John" {
		t.Errorf("expected FirstName 'John', got '%s'", emp.FirstName)
	}
	if emp.LastName != "Doe" {
		t.Errorf("expected LastName 'Doe', got '%s'", emp.LastName)
	}
	if emp.Email != "john@example.com" {
		t.Errorf("expected Email 'john@example.com', got '%s'", emp.Email)
	}
	if emp.Phone != "555-1234" {
		t.Errorf("expected Phone '555-1234', got '%s'", emp.Phone)
	}
	if emp.Street != "123 Main St" {
		t.Errorf("expected Street '123 Main St', got '%s'", emp.Street)
	}
	if emp.AddressLineTwo != "Apt 4" {
		t.Errorf("expected AddressLineTwo 'Apt 4', got '%s'", emp.AddressLineTwo)
	}
	if emp.City != "Springfield" {
		t.Errorf("expected City 'Springfield', got '%s'", emp.City)
	}
	if emp.State != "IL" {
		t.Errorf("expected State 'IL', got '%s'", emp.State)
	}
	if emp.ZipCode != "62701" {
		t.Errorf("expected ZipCode '62701', got '%s'", emp.ZipCode)
	}
	if emp.RoleTitle != "Engineer" {
		t.Errorf("expected RoleTitle 'Engineer', got '%s'", emp.RoleTitle)
	}
	if emp.ManagerName != "Jane Smith" {
		t.Errorf("expected ManagerName 'Jane Smith', got '%s'", emp.ManagerName)
	}
	if emp.HireDate != "2023-01-15" {
		t.Errorf("expected HireDate '2023-01-15', got '%s'", emp.HireDate)
	}
	if emp.StatusType != "Salary" {
		t.Errorf("expected StatusType 'Salary', got '%s'", emp.StatusType)
	}
	if emp.BaseSalary != 85000.0 {
		t.Errorf("expected BaseSalary 85000, got %f", emp.BaseSalary)
	}
	if emp.Bonus != 5000.0 {
		t.Errorf("expected Bonus 5000, got %f", emp.Bonus)
	}
	if emp.Deductions != 2000.0 {
		t.Errorf("expected Deductions 2000, got %f", emp.Deductions)
	}
	if emp.PaidTimeOffHours != 120 {
		t.Errorf("expected PaidTimeOffHours 120, got %d", emp.PaidTimeOffHours)
	}
	if emp.SickHours != 80 {
		t.Errorf("expected SickHours 80, got %d", emp.SickHours)
	}
	if emp.EffectiveFrom != "2023-01-01" {
		t.Errorf("expected EffectiveFrom '2023-01-01', got '%s'", emp.EffectiveFrom)
	}
	if emp.EffectiveTo != "" {
		t.Errorf("expected EffectiveTo empty, got '%s'", emp.EffectiveTo)
	}
}

func TestMapRowToEmployee_NullManagerID(t *testing.T) {
	repo := &EmployeeRepository{}
	row := map[string]interface{}{
		"EmployeeID":   int64(1),
		"CompanyID":    int64(1),
		"DepartmentID": int64(1),
		"FirstName":    "Test",
		"LastName":     "User",
		"Email":        "test@example.com",
		"ManagerID":    nil,
	}

	emp := repo.mapRowToEmployee(row)
	if emp.ManagerID != nil {
		t.Error("expected ManagerID to be nil")
	}
}

func TestMapRowToEmployee_NullValues(t *testing.T) {
	repo := &EmployeeRepository{}
	row := map[string]interface{}{
		"EmployeeID":       int64(1),
		"CompanyID":        int64(1),
		"DepartmentID":     int64(1),
		"FirstName":        nil,
		"LastName":         nil,
		"Email":            nil,
		"Phone":            nil,
		"Street":           nil,
		"AddressLineTwo":   nil,
		"City":             nil,
		"State":            nil,
		"ZipCode":          nil,
		"RoleTitle":        nil,
		"Department":       nil,
		"ManagerName":      nil,
		"HireDate":         nil,
		"TerminationDate":  nil,
		"StatusType":       nil,
		"HourlyPay":        nil,
		"MaxHoursPerWeek":  nil,
		"BaseSalary":       nil,
		"Bonus":            nil,
		"Deductions":       nil,
		"PaidTimeOffHours": nil,
		"SickHours":        nil,
		"EffectiveFrom":    nil,
		"EffectiveTo":      nil,
	}

	emp := repo.mapRowToEmployee(row)
	if emp.FirstName != "" {
		t.Errorf("expected empty FirstName, got '%s'", emp.FirstName)
	}
	if emp.LastName != "" {
		t.Errorf("expected empty LastName, got '%s'", emp.LastName)
	}
	if emp.Email != "" {
		t.Errorf("expected empty Email, got '%s'", emp.Email)
	}
	if emp.HourlyPay != 0.0 {
		t.Errorf("expected HourlyPay 0, got %f", emp.HourlyPay)
	}
	if emp.BaseSalary != 0.0 {
		t.Errorf("expected BaseSalary 0, got %f", emp.BaseSalary)
	}
	if emp.ManagerID != nil {
		t.Error("expected nil ManagerID for null value")
	}
}

func TestMapRowToEmployee_FloatFields(t *testing.T) {
	repo := &EmployeeRepository{}
	row := map[string]interface{}{
		"EmployeeID":       int64(1),
		"CompanyID":        int64(1),
		"DepartmentID":     int64(1),
		"FirstName":        "Test",
		"LastName":         "User",
		"Email":            "test@example.com",
		"StatusType":       "Hourly",
		"HourlyPay":        float64(35.50),
		"MaxHoursPerWeek":  int(40),
		"BaseSalary":       float64(75000.0),
		"Bonus":            float64(5000.0),
		"Deductions":       float64(2000.0),
		"PaidTimeOffHours": int(80),
		"SickHours":        int(60),
	}

	emp := repo.mapRowToEmployee(row)
	if emp.HourlyPay != 35.50 {
		t.Errorf("expected HourlyPay 35.50, got %f", emp.HourlyPay)
	}
	if emp.BaseSalary != 75000.0 {
		t.Errorf("expected BaseSalary 75000, got %f", emp.BaseSalary)
	}
	if emp.Bonus != 5000.0 {
		t.Errorf("expected Bonus 5000, got %f", emp.Bonus)
	}
	if emp.Deductions != 2000.0 {
		t.Errorf("expected Deductions 2000, got %f", emp.Deductions)
	}
}

func TestMapRowToEmployee_IntFromFloat64(t *testing.T) {
	repo := &EmployeeRepository{}
	row := map[string]interface{}{
		"EmployeeID":       float64(1),
		"CompanyID":        float64(2),
		"DepartmentID":     float64(3),
		"FirstName":        "Test",
		"LastName":         "User",
		"Email":            "test@example.com",
		"MaxHoursPerWeek":  float64(40),
		"PaidTimeOffHours": float64(120),
		"SickHours":        float64(80),
	}

	emp := repo.mapRowToEmployee(row)
	if emp.EmployeeID != 1 {
		t.Errorf("expected EmployeeID 1, got %d", emp.EmployeeID)
	}
	if emp.CompanyID != 2 {
		t.Errorf("expected CompanyID 2, got %d", emp.CompanyID)
	}
	if emp.MaxHoursPerWeek != 40 {
		t.Errorf("expected MaxHoursPerWeek 40, got %d", emp.MaxHoursPerWeek)
	}
	if emp.PaidTimeOffHours != 120 {
		t.Errorf("expected PaidTimeOffHours 120, got %d", emp.PaidTimeOffHours)
	}
	if emp.SickHours != 80 {
		t.Errorf("expected SickHours 80, got %d", emp.SickHours)
	}
}

func TestSafeString_Nil(t *testing.T) {
	result := safeString(nil)
	if result != "" {
		t.Errorf("expected empty string, got '%s'", result)
	}
}

func TestSafeString_String(t *testing.T) {
	result := safeString("hello")
	if result != "hello" {
		t.Errorf("expected 'hello', got '%s'", result)
	}
}

func TestSafeString_NonString(t *testing.T) {
	result := safeString(123)
	if result != "123" {
		t.Errorf("expected '123', got '%s'", result)
	}
}

func TestSafeInt64_Nil(t *testing.T) {
	result := safeInt64(nil)
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestSafeInt64_Int64(t *testing.T) {
	result := safeInt64(int64(42))
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestSafeInt64_Int32(t *testing.T) {
	result := safeInt64(int32(99))
	if result != 99 {
		t.Errorf("expected 99, got %d", result)
	}
}

func TestSafeInt64_Int(t *testing.T) {
	result := safeInt64(int(77))
	if result != 77 {
		t.Errorf("expected 77, got %d", result)
	}
}

func TestSafeInt64_Float64(t *testing.T) {
	result := safeInt64(float64(55.9))
	if result != 55 {
		t.Errorf("expected 55, got %d", result)
	}
}

func TestSafeInt64_NonMatching(t *testing.T) {
	result := safeInt64("not-a-number")
	if result != 0 {
		t.Errorf("expected 0 for non-matching type, got %d", result)
	}
}

func TestSafeInt64Ptr_Nil(t *testing.T) {
	result := safeInt64Ptr(nil)
	if result != nil {
		t.Error("expected nil pointer")
	}
}

func TestSafeInt64Ptr_Value(t *testing.T) {
	result := safeInt64Ptr(int64(123))
	if result == nil {
		t.Fatal("expected non-nil pointer")
	}
	if *result != 123 {
		t.Errorf("expected 123, got %d", *result)
	}
}

func TestSafeInt_Nil(t *testing.T) {
	result := safeInt(nil)
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestSafeInt_Int64(t *testing.T) {
	result := safeInt(int64(42))
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestSafeInt_Int32(t *testing.T) {
	result := safeInt(int32(99))
	if result != 99 {
		t.Errorf("expected 99, got %d", result)
	}
}

func TestSafeInt_Int(t *testing.T) {
	result := safeInt(int(77))
	if result != 77 {
		t.Errorf("expected 77, got %d", result)
	}
}

func TestSafeInt_Float64(t *testing.T) {
	result := safeInt(float64(55.9))
	if result != 55 {
		t.Errorf("expected 55, got %d", result)
	}
}

func TestSafeInt_NonMatching(t *testing.T) {
	result := safeInt("not-a-number")
	if result != 0 {
		t.Errorf("expected 0 for non-matching type, got %d", result)
	}
}

func TestSafeFloat64_Nil(t *testing.T) {
	result := safeFloat64(nil)
	if result != 0.0 {
		t.Errorf("expected 0.0, got %f", result)
	}
}

func TestSafeFloat64_Float64(t *testing.T) {
	result := safeFloat64(float64(35.5))
	if result != 35.5 {
		t.Errorf("expected 35.5, got %f", result)
	}
}

func TestSafeFloat64_Float32(t *testing.T) {
	result := safeFloat64(float32(42.25))
	if result != 42.25 {
		t.Errorf("expected 42.25, got %f", result)
	}
}

func TestSafeFloat64_String(t *testing.T) {
	result := safeFloat64("123.45")
	if result != 123.45 {
		t.Errorf("expected 123.45, got %f", result)
	}
}

func TestSafeFloat64_ByteSlice(t *testing.T) {
	result := safeFloat64([]byte("99.99"))
	if result != 99.99 {
		t.Errorf("expected 99.99, got %f", result)
	}
}

func TestSafeFloat64_NonMatching(t *testing.T) {
	result := safeFloat64(true)
	if result != 0.0 {
		t.Errorf("expected 0.0 for non-matching type, got %f", result)
	}
}

func TestEmployeeRepository_ProcedureNamePrefix(t *testing.T) {
	testCases := []struct {
		procedure   string
		name        string
		call        func(*EmployeeRepository, *mockExecutor)
		expectCount int
		rows        []map[string]interface{}
	}{
		{
			procedure: "GetEmployeeProfiles",
			name:      "GetAllEmployees",
			call: func(r *EmployeeRepository, m *mockExecutor) {
				_, _ = r.GetAllEmployees(context.Background())
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"EmployeeID": int64(1), "CompanyID": int64(1), "DepartmentID": int64(1), "FirstName": "Test", "LastName": "User", "Email": "test@example.com", "StatusType": "Salary", "HourlyPay": 0.0, "MaxHoursPerWeek": 0, "BaseSalary": 0.0, "Bonus": 0.0, "Deductions": 0.0, "PaidTimeOffHours": 0, "SickHours": 0}},
		},
		{
			procedure: "GetEmployeeProfileByID",
			name:      "GetEmployeeByID",
			call: func(r *EmployeeRepository, m *mockExecutor) {
				_, _ = r.GetEmployeeByID(context.Background(), 1)
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"EmployeeID": int64(1), "CompanyID": int64(1), "DepartmentID": int64(1), "FirstName": "Test", "LastName": "User", "Email": "test@example.com", "StatusType": "Salary", "HourlyPay": 0.0, "MaxHoursPerWeek": 0, "BaseSalary": 0.0, "Bonus": 0.0, "Deductions": 0.0, "PaidTimeOffHours": 0, "SickHours": 0}},
		},
		{
			procedure: "CreateEmployeeProfile",
			name:      "CreateEmployee",
			call: func(r *EmployeeRepository, m *mockExecutor) {
				_, _ = r.CreateEmployee(context.Background(), models.Employee{})
			},
			expectCount: 1,
			rows:        []map[string]interface{}{{"EmployeeID": int64(1)}},
		},
		{
			procedure: "UpdateEmployeeProfile",
			name:      "UpdateEmployee",
			call: func(r *EmployeeRepository, m *mockExecutor) {
				_, _ = r.UpdateEmployee(context.Background(), models.Employee{})
			},
			expectCount: 1,
			rows:        []map[string]interface{}{},
		},
		{
			procedure: "DeleteEmployeeProfile",
			name:      "DeleteEmployee",
			call: func(r *EmployeeRepository, m *mockExecutor) {
				_ = r.DeleteEmployee(context.Background(), 1)
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
			repo := NewEmployeeRepository(mock)

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
