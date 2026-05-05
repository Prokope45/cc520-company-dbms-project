CREATE PROCEDURE [Org].[sp_Report_DepartmentSalaryRanks]
    @HireDate DATETIME2
AS
BEGIN
    SELECT
        C.[Name] AS CompanyName,
        D.[Name] AS DepartmentName,
        MONTH(E.HireDate) AS HireMonth,
        YEAR(E.HireDate) AS HireYear,
        CONCAT(P.FirstName, ' ', P.LastName) AS EmployeeName,
        S.BaseSalary AS EmployeeSalary,
        RANK() OVER (PARTITION BY D.DepartmentID ORDER BY S.BaseSalary DESC) AS SalaryRank
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Person] P ON E.PersonID = P.PersonID
        INNER JOIN [Org].[Salary] S ON E.EmployeeID = S.SalaryID
        INNER JOIN [Org].[DepEmpRole] DER ON E.EmployeeID = DER.EmployeeID
        INNER JOIN [Org].[Department] D ON DER.DepartmentID = D.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
    WHERE E.HireDate IS NOT NULL 
        AND E.HireDate >= @HireDate
    ORDER BY D.[Name] ASC, SalaryRank ASC;
END
GO
