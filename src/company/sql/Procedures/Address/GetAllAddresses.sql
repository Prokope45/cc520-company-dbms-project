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
    FROM [Org].[Address]
END
GO
