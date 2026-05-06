CREATE PROCEDURE [Org].[sp_Report_UnhiredWithManager_Aggregated]
AS
BEGIN
    SELECT
        C.[Name] AS CompanyName,
        D.[Name] AS DepartmentName,
        COUNT(*) AS UnhiredCount,
        -- Used STRING_AGG to aggregate on concatonated names
        -- https://learn.microsoft.com/en-us/sql/t-sql/functions/string-agg-transact-sql?view=sql-server-ver17
        STRING_AGG(CONCAT(MP.FirstName, ' ', MP.LastName), ', ') AS ManagerNames
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Employee] M ON M.EmployeeID = E.ManagerID
        INNER JOIN [Org].[Person] MP ON MP.PersonID = M.PersonID
        INNER JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
    WHERE E.HireDate IS NULL 
      AND E.ManagerID IS NOT NULL
    GROUP BY C.[Name], D.[Name];
END
GO
