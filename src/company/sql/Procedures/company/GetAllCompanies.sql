CREATE PROCEDURE [Org].[sp_GetAllCompanies]
AS
BEGIN
    SELECT 
        CompanyID,
        Name,
        CreatedDate
    FROM [Org].[Company]
END
GO
