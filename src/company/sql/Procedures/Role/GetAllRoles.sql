CREATE PROCEDURE [Org].[sp_GetAllRoles]
AS
BEGIN
    SELECT [Name]
    FROM [Org].[Role]
    ORDER BY [Name];
END
GO
