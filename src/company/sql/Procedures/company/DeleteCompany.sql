CREATE PROCEDURE [Org].[sp_DeleteCompany]
    @CompanyID INT
AS
BEGIN
    DELETE FROM [Org].[Company]
    WHERE CompanyID = @CompanyID
END
GO
