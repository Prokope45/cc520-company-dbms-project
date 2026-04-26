CREATE PROCEDURE [Org].[sp_GetDepartmentByID]
    @DepartmentID INT
AS
BEGIN
    SELECT 
        DepartmentID,
        CompanyID,
        [Name],
        [Description]
    FROM [Org].[Department]
    WHERE DepartmentID = @DepartmentID
END
GO
