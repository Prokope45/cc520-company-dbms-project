CREATE PROCEDURE [Org].[sp_GetAllDepartments]
AS
BEGIN
    SELECT 
        DepartmentID,
        CompanyID,
        [Name],
        [Description]
    FROM [Org].[Department]
END
GO
