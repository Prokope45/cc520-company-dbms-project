CREATE OR ALTER PROCEDURE Garden.RetrieveAssignmentsForGardener
   @GardenerId INT
AS

SELECT A.AssignmentId,
       A.PlotId,
       A.GardenerId,
       A.StartDate,
       A.EndDate,
       A.Notes,
       P.PlotTag,
       P.LocationDescription,
       P.SizeSqFt,
       P.IsRaisedBed,
       PS.StatusName AS [Status]
FROM Garden.PlotAssignments A
INNER JOIN Garden.Plots P ON P.PlotId = A.PlotId
LEFT JOIN Garden.PlotStatus PS ON PS.StatusId = P.StatusId
WHERE A.GardenerId = @GardenerId;
GO
