CREATE PROCEDURE [Org].[sp_GetCompanyByID]
    @CompanyID INT
AS
BEGIN
    SELECT 
        CompanyID,
        Name,
        CreatedDate
    FROM [Org].[Company]
    WHERE CompanyID = @CompanyID
END
GO
