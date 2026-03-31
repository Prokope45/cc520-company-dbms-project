IF OBJECT_ID(N'Garden.Gardeners') IS NULL
BEGIN
   CREATE TABLE Garden.Gardeners
   (
      GardenerId INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
      FirstName NVARCHAR(32) NOT NULL,
      LastName NVARCHAR(32) NOT NULL,
      Phone NVARCHAR(24) NULL,
      Email NVARCHAR(128) NULL,
      JoinDate DATE NULL,
      CreatedOn DATETIMEOFFSET NOT NULL DEFAULT(SYSDATETIMEOFFSET()),
      UpdatedOn DATETIMEOFFSET NOT NULL DEFAULT(SYSDATETIMEOFFSET()),

   );
END;

/****************************
 * Unique Constraints
 ****************************/

IF NOT EXISTS
   (
      SELECT *
      FROM sys.key_constraints kc
      WHERE kc.parent_object_id = OBJECT_ID(N'Garden.Gardeners')
         AND kc.[name] = N'UK_Garden_Gardeners_Email'
   )
BEGIN
   ALTER TABLE Garden.Gardeners
   ADD CONSTRAINT [UK_Garden_Gardeners_Email] UNIQUE NONCLUSTERED
   (
      Email ASC
   )
END;

/****************************
 * Check Constraints
 ****************************/

IF NOT EXISTS
   (
      SELECT *
      FROM sys.check_constraints cc
      WHERE cc.parent_object_id = OBJECT_ID(N'Garden.Gardeners')
         AND cc.[name] = N'CK_Garden_Gardeners_LastName_FirstName'
   )
BEGIN
   ALTER TABLE Garden.Gardeners
   ADD CONSTRAINT [CK_Garden_Gardeners_LastName_FirstName] CHECK
   (
      FirstName > N'' OR LastName > N''
   )
END;
