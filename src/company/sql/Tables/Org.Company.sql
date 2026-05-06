IF OBJECT_ID(N'[Org].Company') IS NULL
BEGIN
    CREATE TABLE [Org].Company
    (
        [CompanyID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [Name] NVARCHAR(255) NOT NULL UNIQUE,
        [CreatedDate] DATETIME2 DEFAULT GETDATE(),
        CONSTRAINT [CK_Company_Name_Length] CHECK (LEN([Name]) > 0)
    );
END;
