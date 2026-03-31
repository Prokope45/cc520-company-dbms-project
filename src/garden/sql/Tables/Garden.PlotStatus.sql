IF OBJECT_ID(N'Garden.PlotStatus') IS NULL
BEGIN
   CREATE TABLE Garden.PlotStatus
   (
      StatusId INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
      StatusName NVARCHAR(24) NOT NULL,
   );
END;

/****************************
 * Unique Constraints
 ****************************/

IF NOT EXISTS
   (
      SELECT *
      FROM sys.key_constraints kc
      WHERE kc.parent_object_id = OBJECT_ID(N'Garden.PlotStatus')
         AND kc.[name] = N'UK_Garden_PlotStatus_StatusName'
   )
BEGIN
   ALTER TABLE Garden.PlotStatus
   ADD CONSTRAINT [UK_Garden_PlotStatus_StatusName] UNIQUE NONCLUSTERED
   (
      [StatusName] ASC
   )
END;
