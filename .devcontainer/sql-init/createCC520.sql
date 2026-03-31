IF NOT EXISTS (SELECT * FROM sys.server_principals WHERE name = 'cc520-admin')
BEGIN
    CREATE LOGIN [cc520-admin] WITH PASSWORD = 'database3ssentials!';
END
GO

IF DB_ID('CC520') IS NULL
BEGIN
    CREATE DATABASE CC520;
END
GO

use CC520;
IF NOT EXISTS (SELECT * FROM sys.database_principals WHERE name = 'cc520-admin')
BEGIN
    CREATE USER [cc520-admin] FOR LOGIN [cc520-admin];
END
GO
GO

EXEC sp_addrolemember 'db_owner', 'cc520-admin';
