CREATE PROCEDURE [Org].[sp_CreateDepartment]
    @CompanyID INT,
    @Name NVARCHAR(255),
    @Description NVARCHAR(255) = NULL
AS
BEGIN
    INSERT INTO [Org].[Department] (
        [CompanyID],
        [Name],
        [Description]
    )
    VALUES (
        @CompanyID,
        @Name,
        @Description
    )

    SELECT SCOPE_IDENTITY() AS DepartmentID
END
GO
