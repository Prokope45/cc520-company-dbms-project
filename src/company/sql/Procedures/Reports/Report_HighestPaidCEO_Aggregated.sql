CREATE PROCEDURE [Org].[sp_Report_HighestPaidCEO_Aggregated]
AS
BEGIN
    SELECT
        C.[Name] AS CompanyName,
        COUNT(*) AS CEOCount,
        MAX(S.BaseSalary) AS HighestCEOSalary
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Person] P ON P.PersonID = E.PersonID
        INNER JOIN [Org].[Salary] S ON S.SalaryID = E.EmployeeID
        INNER JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[Role] R ON R.RoleID = DER.RoleID
        INNER JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
        INNER JOIN [Org].[EmployeeType] ET ON ET.EmployeeTypeID = E.EmployeeTypeID
    WHERE R.[Name] = 'CEO'
    GROUP BY C.[Name];
END
GO
