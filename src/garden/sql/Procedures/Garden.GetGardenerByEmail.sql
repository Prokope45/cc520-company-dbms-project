CREATE OR ALTER PROCEDURE Garden.GetGardenerByEmail
   @Email NVARCHAR(128)
AS

SELECT G.GardenerId, G.FirstName, G.LastName, G.Phone, G.Email, G.JoinDate
FROM Garden.Gardeners G
WHERE G.Email = @Email;
GO
