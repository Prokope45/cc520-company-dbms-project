CREATE OR ALTER PROCEDURE Garden.UpdatePlotStatus
    @PlotId INT,
    @NewStatusId INT
AS

UPDATE Garden.Plots
SET StatusId = @NewStatusId
WHERE PlotId = @PlotId;
GO
