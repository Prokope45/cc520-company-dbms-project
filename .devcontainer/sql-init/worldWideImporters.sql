IF NOT EXISTS (SELECT *
FROM sys.databases
WHERE name = 'WideWorldImporters')
BEGIN
    RESTORE DATABASE [WideWorldImporters]
    FROM DISK = N'/var/opt/mssql/backup/wwi.bak'
    WITH FILE = 1,
    MOVE N'WWI_Primary' TO N'/var/opt/mssql/data/WideWorldImporters.mdf',
    MOVE N'WWI_UserData' TO N'/var/opt/mssql/data/WideWorldImporters_UserData.ndf',
    MOVE N'WWI_Log' TO N'/var/opt/mssql/data/WideWorldImporters.ldf',
    MOVE N'WWI_InMemory_Data_1' TO N'/var/opt/mssql/data/WideWorldImporters_InMemory_Data_1';
END
GO
-- Create cc520-admin user
USE [WideWorldImporters];
IF NOT EXISTS (SELECT * FROM sys.database_principals WHERE name = 'cc520-admin')
BEGIN
    CREATE USER [cc520-admin] FOR LOGIN [cc520-admin];
END
GO
-- Add cc520-admin user as an owner
EXEC sp_addrolemember 'db_owner', 'cc520-admin';