IF OBJECT_ID(N'Org.EmployeeType') IS NULL
BEGIN
    CREATE TABLE Org.[EmployeeType]
    (
        [EmployeeTypeID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [Name] NVARCHAR(50) NOT NULL DEFAULT ('Salaried') UNIQUE,
        [Description] NVARCHAR(255),
        CONSTRAINT [CK_EmployeeType_Salaried]
            CHECK ([Name] = 'Salaried' OR [Name] = 'PartTime')
    );
END;