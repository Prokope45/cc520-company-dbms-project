CREATE OR ALTER PROCEDURE Garden.FetchGardener
   @GardenerId INT
AS

SELECT G.GardenerId, G.FirstName, G.LastName, G.Phone, G.Email, G.JoinDate
FROM Garden.Gardeners G
WHERE G.GardenerId = @GardenerId;
GO
