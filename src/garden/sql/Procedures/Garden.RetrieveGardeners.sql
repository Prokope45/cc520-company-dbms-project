CREATE OR ALTER PROCEDURE Garden.RetrieveGardeners
AS

SELECT G.GardenerId, G.FirstName, G.LastName, G.Phone, G.Email, G.JoinDate
FROM Garden.Gardeners G;
GO
