CREATE PROCEDURE [Org].[sp_CreateAddress]
    @AddressLineOne NVARCHAR(255),
    @AddressLineTwo NVARCHAR(255) = NULL,
    @CityName NVARCHAR(100),
    @State NVARCHAR(100),
    @ZipCode NVARCHAR(20)
AS
BEGIN
    INSERT INTO [Org].[Address] (
        [AddressLineOne],
        [AddressLineTwo],
        [CityName],
        [State],
        [ZipCode]
    )
    VALUES (
        @AddressLineOne,
        @AddressLineTwo,
        @CityName,
        @State,
        @ZipCode
    )

    SELECT SCOPE_IDENTITY() AS AddressID
END
GO
