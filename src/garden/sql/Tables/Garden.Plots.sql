IF OBJECT_ID(N'Garden.Plots') IS NULL
BEGIN
   CREATE TABLE Garden.Plots
   (
      PlotId INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
      PlotTag NVARCHAR(8) NOT NULL,
      LocationDescription NVARCHAR(128) NULL,
      SizeSqFt INT NULL,
      IsRaisedBed BIT NOT NULL CONSTRAINT DF_Garden_Plots_IsRaisedBed DEFAULT(0),
      StatusId INT NULL,

   );
END;

/****************************
 * Unique Constraints
 ****************************/

IF NOT EXISTS
   (
      SELECT *
      FROM sys.key_constraints kc
      WHERE kc.parent_object_id = OBJECT_ID(N'Garden.Plots')
         AND kc.[name] = N'UK_Garden_Plots_PlotTag'
   )
BEGIN
   ALTER TABLE Garden.Plots
   ADD CONSTRAINT [UK_Garden_Plots_PlotTag] UNIQUE NONCLUSTERED
   (
      [PlotTag] ASC
   )
END;

/****************************
 * Foreign Key Constraints
 ****************************/

IF NOT EXISTS
   (
      SELECT *
      FROM sys.foreign_keys fk
      WHERE fk.parent_object_id = OBJECT_ID(N'Garden.Plots')
         AND fk.[name] = N'FK_Garden_Plots_StatusId'
   )
BEGIN
   ALTER TABLE Garden.Plots
   ADD CONSTRAINT [FK_Garden_Plots_StatusId] FOREIGN KEY
   (
      StatusId
   )
   REFERENCES Garden.PlotStatus(StatusId)
END;
