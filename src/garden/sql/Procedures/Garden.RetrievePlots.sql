CREATE OR ALTER PROCEDURE Garden.RetrievePlots
AS

SELECT
    P.PlotId,
    P.PlotTag,
    P.LocationDescription,
    P.SizeSqFt,
    P.IsRaisedBed,
    P.StatusId
FROM
    Garden.Plots P;
GO
