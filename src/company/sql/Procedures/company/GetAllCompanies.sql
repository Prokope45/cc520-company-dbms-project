CREATE PROCEDURE [sp_GetAllCompanies]
AS
BEGIN
    SELECT 
        CompanyID,
        Name,
        CreatedDate
    FROM [Company]
END
GO
