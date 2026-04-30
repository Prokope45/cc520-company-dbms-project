IF OBJECT_ID(N'Org.Person') IS NULL
BEGIN
    CREATE TABLE [Org].[Person] (
        [PersonID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [AddressID] INT NOT NULL,
        [FirstName] NVARCHAR(100) NOT NULL,
        [LastName] NVARCHAR(100) NOT NULL,
        [Email] NVARCHAR(255) NOT NULL UNIQUE,
        [Phone] NVARCHAR(50) Null,
        CONSTRAINT [FK_Person_Address] FOREIGN KEY ([AddressID])
            REFERENCES Org.[Address]([AddressID])
            ON DELETE CASCADE,
        CONSTRAINT [CK_Person_Email_Length] CHECK (LEN([Email]) > 0)
    );
END;
