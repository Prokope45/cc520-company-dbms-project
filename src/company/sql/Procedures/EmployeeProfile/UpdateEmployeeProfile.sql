CREATE PROCEDURE [Org].[sp_UpdateEmployeeProfile]
    @EmployeeID INT,
    @DepartmentID INT,
    @FirstName NVARCHAR(100),
    @LastName NVARCHAR(100),
    @Email NVARCHAR(255),
    @Phone NVARCHAR(50) = NULL,
    @Street NVARCHAR(255),
    @AddressLineTwo NVARCHAR(255) = NULL,
    @City NVARCHAR(100),
    @State NVARCHAR(100),
    @ZipCode NVARCHAR(20),
    @RoleTitle NVARCHAR(100),
    @ManagerID INT = NULL,
    @HireDate DATETIME2 = NULL,
    @TerminationDate DATETIME2 = NULL,
    @StatusType NVARCHAR(50),
    
    -- Hourly Data (Nullable)
    @HourlyPay DECIMAL(18,2) = NULL,
    @MaxHoursPerWeek INT = NULL,
    
    -- Salary Data (Nullable)
    @BaseSalary DECIMAL(18,2) = NULL,
    @Bonus DECIMAL(18,2) = 0,
    @Deductions DECIMAL(18,2) = 0,
    @PaidTimeOffHours INT = 0,
    @SickHours INT = 0,
    @EffectiveFrom DATETIME2 = NULL,
    @EffectiveTo DATETIME2 = NULL
AS
BEGIN
    BEGIN TRY
        BEGIN TRANSACTION;

        -- Get related IDs
        DECLARE @PersonID INT = (
            SELECT PersonID
                FROM [Org].[Employee]
            WHERE EmployeeID = @EmployeeID
        );

        DECLARE @AddressID INT = (
            SELECT AddressID
                FROM [Org].[Person]
            WHERE PersonID = @PersonID
        );

        -- Update Address
        UPDATE [Org].[Address]
        SET AddressLineOne = @Street,
            AddressLineTwo = @AddressLineTwo,
            CityName = @City,
            [State] = @State,
            ZipCode = @ZipCode
        WHERE AddressID = @AddressID;

        -- Update Person
        UPDATE [Org].[Person]
        SET FirstName = @FirstName,
            LastName = @LastName,
            Email = @Email,
            Phone = @Phone
        WHERE PersonID = @PersonID;

        -- Get new EmployeeTypeID
        DECLARE @NewEmployeeTypeID INT = (
            SELECT EmployeeTypeID
                FROM [Org].[EmployeeType]
            WHERE [Name] = @StatusType
        );

        -- Update Employee (Manager & Type)
        UPDATE [Org].[Employee]
        SET ManagerID = @ManagerID,
            EmployeeTypeID = @NewEmployeeTypeID,
            HireDate = @HireDate,
            TerminationDate = @TerminationDate
        WHERE EmployeeID = @EmployeeID;

        -- Handle Pay Structure
        IF @StatusType = 'Hourly'
        BEGIN
            -- Clean up Salary data if it exists for employee
            DELETE FROM [Org].[SalaryEmployee]
            WHERE EmployeeID = @EmployeeID;

            DELETE FROM [Org].[Salary]
            WHERE SalaryID = @EmployeeID;

            -- Update or Insert Hourly data
            IF EXISTS (
                SELECT EmployeeID
                    FROM [Org].[HourlyEmployee]
                WHERE EmployeeID = @EmployeeID
            )
                UPDATE [Org].[HourlyEmployee]
                SET HourlyPay = @HourlyPay,
                    MaxHoursPerWeek = @MaxHoursPerWeek
                WHERE EmployeeID = @EmployeeID;
            ELSE
                INSERT INTO [Org].[HourlyEmployee] (
                    EmployeeID,
                    HourlyPay,
                    MaxHoursPerWeek
                )
                VALUES (
                    @EmployeeID,
                    @HourlyPay,
                    @MaxHoursPerWeek
                );
        END
        ELSE IF @StatusType = 'Salary'
        BEGIN
            -- Clean up Hourly data if it exists
            DELETE FROM [Org].[HourlyEmployee]
            WHERE EmployeeID = @EmployeeID;

            -- Update or Insert Salary data
            IF EXISTS (
                SELECT SalaryID
                    FROM [Org].[Salary]
                WHERE SalaryID = @EmployeeID
            )
                UPDATE [Org].[Salary]
                SET BaseSalary = ISNULL(@BaseSalary, BaseSalary),
                    Bonus = ISNULL(@Bonus, Bonus),
                    Deductions = ISNULL(@Deductions, Deductions),
                    EffectiveFrom = ISNULL(@EffectiveFrom, EffectiveFrom),
                    EffectiveTo = ISNULL(@EffectiveTo, EffectiveTo)
                WHERE SalaryID = @EmployeeID;
            ELSE
                INSERT INTO [Org].[Salary] (
                    SalaryID,
                    BaseSalary,
                    Bonus,
                    Deductions,
                    EffectiveFrom,
                    EffectiveTo)
                VALUES (
                    @EmployeeID,
                    @BaseSalary,
                    @Bonus,
                    @Deductions,
                    ISNULL(@EffectiveFrom, GETDATE()),
                    ISNULL(@EffectiveTo, DATEADD(year, 1, GETDATE()))
                );

            -- Update or Insert SalaryEmployee data
            IF EXISTS (
                SELECT EmployeeID
                    FROM [Org].[SalaryEmployee]
                WHERE EmployeeID = @EmployeeID
            )
                UPDATE [Org].[SalaryEmployee]
                SET PaidTimeOffHours = @PaidTimeOffHours,
                    SickHours = @SickHours
                WHERE EmployeeID = @EmployeeID;
            ELSE
                INSERT INTO [Org].[SalaryEmployee] (
                    EmployeeID,
                    SalaryID,
                    PaidTimeOffHours,
                    SickHours
                )
                VALUES (@EmployeeID, @EmployeeID, @PaidTimeOffHours, @SickHours);
        END

        -- Update Role/Department mapping
        DECLARE @RoleID INT;
        SELECT @RoleID = RoleID
            FROM [Org].[Role]
        WHERE [Name] = @RoleTitle;

        IF @RoleID IS NOT NULL AND @DepartmentID IS NOT NULL
        BEGIN
            -- Remove old DER mappings for this employee
            DELETE FROM [Org].[DepEmpRole]
            WHERE EmployeeID = @EmployeeID;

            -- Insert new mapping
            INSERT INTO [Org].[DepEmpRole] (
                RoleID,
                DepartmentID,
                EmployeeID
            )
            VALUES (@RoleID, @DepartmentID, @EmployeeID);
        END

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
