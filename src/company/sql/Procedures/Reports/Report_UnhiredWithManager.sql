CREATE PROCEDURE [Org].[sp_Report_UnhiredWithManager]
AS
BEGIN
    SELECT 
        CONCAT(P.FirstName, ' ', P.LastName) AS EmployeeName,
        CONCAT(MP.FirstName, ' ', MP.LastName) AS ManagerAssigned
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Person] P ON P.PersonID = E.PersonID
        INNER JOIN [Org].[Employee] M ON M.EmployeeID = E.ManagerID
        INNER JOIN [Org].[Person] MP ON MP.PersonID = M.PersonID
    WHERE E.HireDate IS NULL 
      AND E.ManagerID IS NOT NULL;
END
GO
