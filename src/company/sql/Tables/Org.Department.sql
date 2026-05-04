IF OBJECT_ID(N'Org.Department') IS NULL
BEGIN
    CREATE TABLE [Org].Department
    (
        [DepartmentID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [CompanyID] INT NOT NULL,
        [Name] NVARCHAR(255) NOT NULL,
        [Description] NVARCHAR(255),
        CONSTRAINT [FK_Department_Company] FOREIGN KEY ([CompanyID])
            REFERENCES Org.[Company] ([CompanyID])
            ON DELETE CASCADE,
        CONSTRAINT [UK_Department_CompanyName] UNIQUE ([CompanyID], [Name])
    );
END;
