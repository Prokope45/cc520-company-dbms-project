CREATE PROCEDURE [Org].[sp_Report_UnhiredWithManager]
AS
BEGIN
    SELECT 
        C.[Name] AS CompanyName,
        D.[Name] AS DepartmentName,
        CONCAT(P.FirstName, ' ', P.LastName) AS EmployeeName,
        CONCAT(MP.FirstName, ' ', MP.LastName) AS ManagerAssigned
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Person] P ON P.PersonID = E.PersonID
        INNER JOIN [Org].[Employee] M ON M.EmployeeID = E.ManagerID
        INNER JOIN [Org].[Person] MP ON MP.PersonID = M.PersonID
        INNER JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
    WHERE E.HireDate IS NULL 
      AND E.ManagerID IS NOT NULL;
END
GO
