CREATE PROCEDURE [sp_DeleteCompany]
    @CompanyID INT
AS
BEGIN
    DELETE FROM [Company]
    WHERE CompanyID = @CompanyID
END
GO
