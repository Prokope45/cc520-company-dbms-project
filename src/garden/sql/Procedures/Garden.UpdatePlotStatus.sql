CREATE OR ALTER PROCEDURE Garden.UpdatePlotStatus
    @PlotId INT,
    @NewStatusId INT
AS

UPDATE Garden.Plots
SET StatusId = @NewStatusId
FROM Garden.Plots P
WHERE P.PlotId = @PlotId;
GO
