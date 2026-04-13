CREATE OR ALTER PROCEDURE Garden.GetAvailablePlots
AS

SELECT
    P.PlotId,
    P.PlotTag,
    P.LocationDescription,
    P.SizeSqFt,
    P.IsRaisedBed,
    P.StatusId
FROM
    Garden.Plots P
WHERE
    StatusId = (
        SELECT PS.StatusId
        FROM Garden.PlotStatus PS
        WHERE PS.StatusName = 'Available'
    );
GO
