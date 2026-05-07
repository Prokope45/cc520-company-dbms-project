CREATE PROCEDURE [Org].[sp_Report_HighestPaidExecutive]
AS
BEGIN
    SELECT TOP(1)
        C.[Name] AS CompanyName,
        CONCAT(P.FirstName, ' ', P.LastName) AS ExecutiveName,
        S.BaseSalary AS ExecutiveSalary,
        R.[Name] AS RoleName
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Person] P ON E.PersonID = P.PersonID
        INNER JOIN [Org].[Salary] S ON S.SalaryID = E.EmployeeID
        INNER JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[Role] R ON R.RoleID = DER.RoleID
        INNER JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
        INNER JOIN [Org].[EmployeeType] ET ON ET.EmployeeTypeID = E.EmployeeTypeID
    WHERE R.[Name] IN ('CEO', 'CFO', 'CTO', 'CHRO')
    ORDER BY S.BaseSalary DESC;
END
GO
