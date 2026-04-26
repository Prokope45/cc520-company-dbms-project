CREATE PROCEDURE [Org].[sp_DeleteDepartment]
    @DepartmentID INT
AS
BEGIN
    DELETE FROM [Org].[Department]
    WHERE DepartmentID = @DepartmentID
END
GO
