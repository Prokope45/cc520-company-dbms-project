CREATE OR ALTER PROCEDURE Garden.CreateGardener
   @FirstName NVARCHAR(32),
   @LastName NVARCHAR(32),
   @Phone NVARCHAR(24),
   @Email NVARCHAR(128),
   @JoinDate DATE,
   @GardenerId INT OUTPUT
AS

INSERT Garden.Gardeners(FirstName, LastName, Phone, Email, JoinDate)
VALUES(@FirstName, @LastName, @Phone, @Email, @JoinDate);

SET @GardenerId = SCOPE_IDENTITY();
GO
