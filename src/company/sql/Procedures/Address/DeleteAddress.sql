CREATE PROCEDURE [Org].[sp_DeleteAddress]
    @AddressID INT
AS
BEGIN
    DELETE FROM [Org].[Address]
    WHERE AddressID = @AddressID
END
GO
