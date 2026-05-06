CREATE PROCEDURE [Org].[sp_UpdateAddress]
    @AddressID INT,
    @AddressLineOne NVARCHAR(255),
    @AddressLineTwo NVARCHAR(255) = NULL,
    @CityName NVARCHAR(100),
    @State NVARCHAR(100),
    @ZipCode NVARCHAR(20)
AS
BEGIN
    UPDATE [Org].[Address]
    SET [AddressLineOne] = @AddressLineOne,
        [AddressLineTwo] = @AddressLineTwo,
        [CityName] = @CityName,
        [State] = @State,
        [ZipCode] = @ZipCode
    WHERE AddressID = @AddressID
END
GO
