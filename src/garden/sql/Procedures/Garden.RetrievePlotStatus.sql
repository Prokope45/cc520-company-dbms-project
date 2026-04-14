CREATE OR ALTER PROCEDURE Garden.RetrievePlotStatus
    @PlotId INT
AS

SELECT PS.StatusId, PS.StatusName
FROM Garden.PlotStatus PS
WHERE PS.StatusId = (
    Select P.StatusId
    FROM Garden.Plots P
    WHERE P.PlotId = @PlotId
)
GO
