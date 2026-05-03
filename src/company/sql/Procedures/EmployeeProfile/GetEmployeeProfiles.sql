CREATE PROCEDURE [Org].[sp_GetEmployeeProfiles]
AS
BEGIN
    SELECT 
        E.EmployeeID,
        D.CompanyID,
        DER.DepartmentID,
        D.Name AS Department,
        P.FirstName,
        P.LastName,
        P.Email,
        P.Phone,
        A.AddressLineOne AS Street,
        A.AddressLineTwo,
        A.CityName AS City,
        A.State,
        A.ZipCode,
        R.Name AS RoleTitle,
        E.ManagerID,
        CONCAT(Mgr_P.FirstName, ' ', Mgr_P.LastName) AS ManagerName,
        CONVERT(VARCHAR(10), E.HireDate, 120) AS HireDate,
        CONVERT(VARCHAR(10), E.TerminationDate, 120) AS TerminationDate,
        ET.Name AS StatusType,

        -- Hourly Data
        HE.HourlyPay,
        HE.MaxHoursPerWeek,

        -- Salary Data
        S.BaseSalary,
        S.Bonus,
        S.Deductions,
        SE.PaidTimeOffHours,
        SE.SickHours,
        CONVERT(VARCHAR(10), S.EffectiveFrom, 120) AS EffectiveFrom,
        CONVERT(VARCHAR(10), S.EffectiveTo, 120) AS EffectiveTo

    FROM [Org].[Employee] E
        LEFT JOIN [Org].[Person] P ON P.PersonID = E.PersonID
        LEFT JOIN [Org].[Address] A ON A.AddressID = P.AddressID
        LEFT JOIN [Org].[EmployeeType] ET ON ET.EmployeeTypeID = E.EmployeeTypeID
        LEFT JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        LEFT JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        LEFT JOIN [Org].[Role] R ON R.RoleID = DER.RoleID

        -- Self-join for Manager Name
        LEFT JOIN [Org].[Employee] Mgr ON Mgr.EmployeeID = E.ManagerID
        LEFT JOIN [Org].[Person] Mgr_P ON Mgr_P.PersonID = Mgr.PersonID

        -- Pay Data Joins
        LEFT JOIN [Org].[HourlyEmployee] HE ON HE.EmployeeID = E.EmployeeID
        LEFT JOIN [Org].[SalaryEmployee] SE ON SE.EmployeeID = E.EmployeeID
        LEFT JOIN [Org].[Salary] S ON S.SalaryID = SE.EmployeeID
END
GO
