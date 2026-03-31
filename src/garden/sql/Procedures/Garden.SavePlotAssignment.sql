CREATE OR ALTER PROCEDURE Garden.SavePlotAssignment
   @PlotId INT,
   @GardenerId INT,
   @StartDate DATE,
   @EndDate DATE,
   @Notes NVARCHAR(256)
AS

MERGE Garden.PlotAssignments T
USING
      (
         VALUES(@PlotId, @GardenerId, @StartDate, @EndDate, @Notes)
      ) S(PlotId, GardenerId, StartDate, EndDate, Notes)
   ON S.PlotId = T.PlotId
      AND S.GardenerId = T.GardenerId
      AND S.StartDate = T.StartDate
WHEN MATCHED AND NOT EXISTS
      (
         SELECT S.EndDate, S.Notes
         INTERSECT
         SELECT  T.EndDate, T.Notes
      ) THEN
   UPDATE
   SET
      EndDate = S.EndDate,
      Notes = S.Notes,
      UpdatedOn = SYSDATETIMEOFFSET()
WHEN NOT MATCHED THEN
   INSERT(PlotId, GardenerId, StartDate, EndDate, Notes)
   VALUES(S.PlotId, S.GardenerId, S.StartDate, S.EndDate, S.Notes);
GO
