IF NOT EXISTS
   (
      SELECT *
      FROM sys.schemas s
      WHERE s.[name] = N'Org'
   )
BEGIN
   EXEC(N'CREATE SCHEMA [Org] AUTHORIZATION [dbo]');
END