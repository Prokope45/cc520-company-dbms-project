CREATE PROCEDURE [Org].[sp_CreateCompany]
    @Name NVARCHAR(255)
AS
BEGIN
    INSERT INTO [Org].[Company] ([Name])
    VALUES (@Name)
    
    SELECT SCOPE_IDENTITY() AS CompanyID
END
GO
