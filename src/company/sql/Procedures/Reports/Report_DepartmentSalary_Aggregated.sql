/*
Provides summary statistics of the number of employees, and the upper
and lower bounds and average salary of employees in each department.
*/

CREATE PROCEDURE [Org].[sp_Report_DepartmentSalary_Aggregated]
    @HireDate DATETIME2
AS
BEGIN
    SELECT
        C.[Name] AS CompanyName,
        D.[Name] AS DepartmentName,
        COUNT(E.EmployeeID) AS EmployeeCount,
        AVG(S.BaseSalary) AS AverageSalary,
        MAX(S.BaseSalary) AS HighestSalary,
        MIN(S.BaseSalary) AS LowestSalary
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Salary] S ON S.SalaryID = E.EmployeeID
        INNER JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
    WHERE E.HireDate IS NOT NULL 
        AND E.HireDate >= @HireDate
    GROUP BY 
        C.[Name], D.[Name]
    ORDER BY 
        C.[Name] ASC, D.[Name] ASC;
END
GO
