CREATE PROCEDURE [sp_UpdateCompany]
    @CompanyID INT,
    @Name NVARCHAR(255)
AS
BEGIN
    UPDATE [Company]
    SET [Name] = @Name
    WHERE CompanyID = @CompanyID
END
GO
