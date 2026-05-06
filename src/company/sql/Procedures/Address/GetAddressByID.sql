CREATE PROCEDURE [Org].[sp_GetAddressByID]
    @AddressID INT
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
    WHERE AddressID = @AddressID
END
GO
