CREATE OR ALTER PROCEDURE Garden.GetPlotStatusById
AS

SELECT
    PS.StatusId,
    PS.StatusName
FROM
    Garden.PlotStatus PS
WHERE PS.StatusId = @StatusId;
GO
