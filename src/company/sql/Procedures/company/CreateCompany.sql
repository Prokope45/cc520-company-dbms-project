CREATE PROCEDURE [sp_CreateCompany]
    @Name NVARCHAR(255)
AS
BEGIN
    INSERT INTO [Company] ([Name])
    VALUES (@Name)
    
    SELECT SCOPE_IDENTITY() AS CompanyID
END
GO
