CREATE PROCEDURE [Org].[sp_CreateEmployeeProfile]
    @CompanyID INT,
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

        -- Insert Address
        INSERT INTO [Org].[Address] (
            AddressLineOne,
            AddressLineTwo,
            CityName,
            [State],
            ZipCode
        )
        VALUES (@Street, @AddressLineTwo, @City, @State, @ZipCode);

        -- Insert Person
        INSERT INTO [Org].[Person] (
            AddressID,
            FirstName,
            LastName,
            Email,
            Phone
        )
        VALUES (
            SCOPE_IDENTITY(),
            @FirstName,
            @LastName,
            @Email,
            @Phone
        );

        DECLARE @PersonID INT;
        SET @PersonID = SCOPE_IDENTITY();

        -- Get EmployeeTypeID
        DECLARE @EmployeeTypeID INT;
        SELECT @EmployeeTypeID = EmployeeTypeID
            FROM [Org].[EmployeeType]
        WHERE [Name] = @StatusType;

        -- Insert Employee
        -- Use provided HireDate or default to GETDATE()
        DECLARE @ActualHireDate DATETIME2 = ISNULL(@HireDate, GETDATE());

        INSERT INTO [Org].[Employee] (
            ManagerID,
            PersonID, 
            EmployeeTypeID, 
            HireDate
        )
        VALUES (
            @ManagerID,
            @PersonID,
            @EmployeeTypeID,
            @ActualHireDate
        );

        DECLARE @EmployeeID INT;
        SET @EmployeeID = SCOPE_IDENTITY();

        -- Insert Pay Structure Details
        IF @StatusType = 'Hourly'
        BEGIN
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
            -- Salary table
            INSERT INTO [Org].[Salary] (
                SalaryID,
                BaseSalary,
                Bonus,
                Deductions,
                EffectiveFrom,
                EffectiveTo
            )
            VALUES (
                @EmployeeID,
                @BaseSalary,
                @Bonus,
                @Deductions,
                ISNULL(@EffectiveFrom, GETDATE()),
                ISNULL(@EffectiveTo, DATEADD(year, 1, GETDATE()))
            );
            
            -- SalaryEmployee table
            INSERT INTO [Org].[SalaryEmployee] (
                EmployeeID,
                PaidTimeOffHours,
                SickHours
            )
            VALUES (
                @EmployeeID,
                @PaidTimeOffHours,
                @SickHours
            );
        END

        -- Insert Role/Department mapping
        DECLARE @RoleID INT;
        SELECT @RoleID = RoleID
            FROM [Org].[Role]
        WHERE [Name] = @RoleTitle;

        IF @RoleID IS NOT NULL AND @DepartmentID IS NOT NULL
        BEGIN
            INSERT INTO [Org].[DepEmpRole] (
                RoleID,
                DepartmentID,
                EmployeeID
            )
            VALUES (@RoleID, @DepartmentID, @EmployeeID);
        END

        COMMIT TRANSACTION;

        -- Return the newly created EmployeeID
        SELECT @EmployeeID AS EmployeeID;

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
