IF OBJECT_ID(N'Org.Address') IS NULL
BEGIN
    CREATE TABLE [Org].[Address] (
        [AddressID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [AddressLineOne] NVARCHAR(255) NOT NULL,
        [AddressLineTwo] NVARCHAR(255),
        [CityName] NVARCHAR(100) NOT NULL,
        [State] NVARCHAR(100) NOT NULL,
        [ZipCode] NVARCHAR(20) NOT NULL,
        CONSTRAINT [CK_Address_CityName_Length] CHECK (LEN([CityName]) > 0)
    );
END;
