CREATE PROCEDURE [Org].[sp_GetAllAddresses]
AS
BEGIN
    SELECT
        AddressID,
        AddressLineOne,
        AddressLineTwo,
        CityName,
        [State],
        ZipCode
    FROM [Address]
END
GO
