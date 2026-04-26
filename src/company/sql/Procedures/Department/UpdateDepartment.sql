CREATE PROCEDURE [Org].[sp_UpdateDepartment]
    @DepartmentID INT,
    @CompanyID INT,
    @Name NVARCHAR(255),
    @Description NVARCHAR(255) = NULL
AS
BEGIN
    UPDATE [Org].[Department]
    SET [CompanyID] = @CompanyID,
        [Name] = @Name,
        [Description] = @Description
    WHERE DepartmentID = @DepartmentID
END
GO
