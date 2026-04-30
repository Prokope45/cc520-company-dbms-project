CREATE PROCEDURE [Org].[sp_DeleteAddress]
    @AddressID INT
AS
BEGIN
    DELETE FROM [Address]
    WHERE AddressID = @AddressID
END
GO
