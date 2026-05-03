CREATE PROCEDURE [Org].[sp_DeleteDepartment]
    @DepartmentID INT
AS
BEGIN
    -- Before deleting department, delete DepEmpRole records for employees in department.
    -- This leaves employees as "floating" without a department or role.
    DELETE FROM [Org].[DepEmpRole]
    WHERE DepartmentID = @DepartmentID;

    DELETE FROM [Org].[Department]
    WHERE DepartmentID = @DepartmentID;
END
GO
