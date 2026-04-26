CREATE PROCEDURE [Org].[sp_UpdateCompany]
    @CompanyID INT,
    @Name NVARCHAR(255)
AS
BEGIN
    UPDATE [Org].[Company]
    SET [Name] = @Name
    WHERE CompanyID = @CompanyID
END
GO
