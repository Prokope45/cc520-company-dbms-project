DECLARE @PlotStatusStaging TABLE
(
   StatusName NVARCHAR(24) NOT NULL PRIMARY KEY
);

/***************************** Modify values here *****************************/

INSERT @PlotStatusStaging(StatusName)
VALUES
   (N'Active'),
   (N'Maintenance'),
   (N'Available');

/******************************************************************************/

MERGE Garden.PlotStatus T
USING @PlotStatusStaging S ON S.StatusName = T.StatusName
WHEN NOT MATCHED THEN
   INSERT(StatusName)
   VALUES(S.StatusName);
