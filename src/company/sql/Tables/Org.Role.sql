IF OBJECT_ID(N'Org.Role') IS NULL
BEGIN
    CREATE TABLE Org.[Role] (
        [RoleID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [Name] NVARCHAR(100) NOT NULL UNIQUE,
        [Description] NVARCHAR(255),
        CONSTRAINT [CK_Role_Name_Length] CHECK (LEN([Name]) > 0),
        CONSTRAINT [CK_Role_Types_Available] CHECK (
            ([Name] = 'Member') OR
            ([Name] = 'Manager') OR
            ([Name] = 'Director') OR
            ([Name] = 'CEO')
        )
    );
END;
