IF OBJECT_ID(N'Garden.PlotAssignments') IS NULL
BEGIN
   CREATE TABLE Garden.PlotAssignments
   (
      AssignmentId INT NOT NULL IDENTITY(1, 1) PRIMARY KEY,
      PlotId INT NOT NULL,
      GardenerId INT NOT NULL,
      StartDate DATE NOT NULL,
      EndDate DATE NULL,
      Notes NVARCHAR(256) NULL,
      CreatedOn DATETIMEOFFSET NOT NULL
         CONSTRAINT DF_Garden_PlotAssignments_CreatedOn DEFAULT(SYSDATETIMEOFFSET()),
      UpdatedOn DATETIMEOFFSET NOT NULL
         CONSTRAINT DF_Garden_PlotAssignments_UpdatedOn DEFAULT(SYSDATETIMEOFFSET()),


      CONSTRAINT FK_Garden_PlotAssignments_Garden_Plots FOREIGN KEY(PlotId)
      REFERENCES Garden.Plots(PlotId),

      CONSTRAINT FK_Garden_PlotAssignments_Garden_Gardeners FOREIGN KEY(GardenerId)
      REFERENCES Garden.Gardeners(GardenerId)
   );
END

/****************************
 * Unique Constraints
 ****************************/

IF NOT EXISTS
   (
      SELECT *
      FROM sys.key_constraints kc
      WHERE kc.parent_object_id = OBJECT_ID(N'Garden.PlotAssignments')
         AND kc.[name] = N'UK_Garden_PlotAssignments_PlotId_GardenerId_StartDate'
   )
BEGIN
   ALTER TABLE Garden.PlotAssignments
   ADD CONSTRAINT [UK_Garden_PlotAssignments_PlotId_GardenerId_StartDate] UNIQUE NONCLUSTERED
   (
      PlotId,
      GardenerId,
      StartDate
   )
END;

