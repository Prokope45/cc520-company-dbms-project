CREATE PROCEDURE [sp_GetCompanyByID]
    @CompanyID INT
AS
BEGIN
    SELECT 
        CompanyID,
        Name,
        CreatedDate
    FROM [Company]
    WHERE CompanyID = @CompanyID
END
GO
