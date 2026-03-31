IF NOT EXISTS
   (
      SELECT *
      FROM sys.schemas s
      WHERE s.[name] = N'Garden'
   )
BEGIN
   EXEC(N'CREATE SCHEMA [Garden] AUTHORIZATION [dbo]');
END