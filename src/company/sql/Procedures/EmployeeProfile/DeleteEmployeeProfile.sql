CREATE PROCEDURE [Org].[sp_DeleteEmployeeProfile]
    @EmployeeID INT
AS
BEGIN
    BEGIN TRY
        BEGIN TRANSACTION;

        -- Get PersonID and AddressID before deletion
        DECLARE @PersonID INT;
        SELECT @PersonID = PersonID
            FROM [Org].[Employee]
        WHERE EmployeeID = @EmployeeID;

        DECLARE @AddressID INT;
        IF @PersonID IS NOT NULL
            SELECT @AddressID = AddressID
                FROM [Org].[Person]
            WHERE PersonID = @PersonID;

        -- Remove DER mapping
        DELETE FROM [Org].[DepEmpRole]
        WHERE EmployeeID = @EmployeeID;

        -- Remove from Pay structures
        DELETE FROM [Org].[HourlyEmployee]
        WHERE EmployeeID = @EmployeeID;

        DELETE FROM [Org].[SalaryEmployee]
        WHERE EmployeeID = @EmployeeID;

        DELETE FROM [Org].[Salary]
        WHERE SalaryID = @EmployeeID;

        -- Unassign this employee as a manager to others
        UPDATE [Org].[Employee]
        SET ManagerID = NULL
        WHERE ManagerID = @EmployeeID;

        -- Delete Employee
        DELETE FROM [Org].[Employee]
        WHERE EmployeeID = @EmployeeID;

        -- Delete Person
        IF @PersonID IS NOT NULL
            DELETE FROM [Org].[Person]
            WHERE PersonID = @PersonID;

        -- Delete Address
        IF @AddressID IS NOT NULL
            DELETE FROM [Org].[Address]
            WHERE AddressID = @AddressID;

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;

        DECLARE @ErrorMessage NVARCHAR(4000) = ERROR_MESSAGE();
        DECLARE @ErrorSeverity INT = ERROR_SEVERITY();
        DECLARE @ErrorState INT = ERROR_STATE();

        RAISERROR (@ErrorMessage, @ErrorSeverity, @ErrorState);
    END CATCH
END
GO
