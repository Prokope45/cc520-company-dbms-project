DECLARE @PlotStaging TABLE
(
   PlotTag NVARCHAR(8) NOT NULL PRIMARY KEY,
   LocationDescription NVARCHAR(128) NOT NULL,
   SizeSqFt INT NOT NULL,
   IsRaisedBed BIT NOT NULL,
   StatusName NVARCHAR(24) NOT NULL
);

/***************************** Modify values here *****************************/
-- The following data was gnerated by Claude Opus 4.6
INSERT @PlotStaging(PlotTag, LocationDescription, SizeSqFt, IsRaisedBed, StatusName)
VALUES
   (N'A1', N'Northwest Corner', 120, 1, N'Active'),
   (N'A2', N'Northwest Corner', 100, 0, N'Active'),
   (N'A3', N'Northwest Corner', 110, 1, N'Active'),
   (N'A4', N'Northwest Corner', 95, 0, N'Available'),
   (N'A5', N'Northwest Corner', 130, 1, N'Active'),
   (N'B1', N'East Fence', 90, 1, N'Maintenance'),
   (N'B2', N'East Fence', 100, 0, N'Active'),
   (N'B3', N'East Fence', 85, 1, N'Active'),
   (N'B4', N'East Fence', 110, 0, N'Available'),
   (N'B5', N'East Fence', 95, 1, N'Active'),
   (N'B6', N'East Fence', 105, 0, N'Maintenance'),
   (N'C1', N'South Gate', 140, 0, N'Active'),
   (N'C2', N'South Gate', 120, 1, N'Active'),
   (N'C3', N'South Gate', 140, 0, N'Available'),
   (N'C4', N'South Gate', 100, 1, N'Active'),
   (N'C5', N'South Gate', 115, 0, N'Maintenance'),
   (N'C6', N'South Gate', 130, 1, N'Active'),
   (N'D1', N'Central Meadow', 150, 0, N'Active'),
   (N'D2', N'Central Meadow', 160, 0, N'Active'),
   (N'D3', N'Central Meadow', 145, 1, N'Available'),
   (N'D4', N'Central Meadow', 135, 0, N'Active'),
   (N'D5', N'Central Meadow', 155, 1, N'Active'),
   (N'D6', N'Central Meadow', 140, 0, N'Maintenance'),
   (N'E1', N'West Terrace', 80, 1, N'Active'),
   (N'E2', N'West Terrace', 85, 1, N'Active'),
   (N'E3', N'West Terrace', 90, 0, N'Available'),
   (N'E4', N'West Terrace', 75, 1, N'Active'),
   (N'E5', N'West Terrace', 95, 0, N'Active'),
   (N'E6', N'West Terrace', 80, 1, N'Maintenance'),
   (N'F1', N'Greenhouse Row', 60, 1, N'Active'),
   (N'F2', N'Greenhouse Row', 65, 1, N'Active'),
   (N'F3', N'Greenhouse Row', 55, 1, N'Active'),
   (N'F4', N'Greenhouse Row', 70, 1, N'Available'),
   (N'F5', N'Greenhouse Row', 60, 1, N'Active'),
   (N'F6', N'Greenhouse Row', 50, 1, N'Maintenance'),
   (N'G1', N'North Hill', 200, 0, N'Active'),
   (N'G2', N'North Hill', 180, 0, N'Active'),
   (N'G3', N'North Hill', 190, 0, N'Available'),
   (N'G4', N'North Hill', 210, 0, N'Active'),
   (N'G5', N'North Hill', 175, 1, N'Active'),
   (N'G6', N'North Hill', 195, 0, N'Maintenance'),
   (N'H1', N'Pond Side', 100, 0, N'Active'),
   (N'H2', N'Pond Side', 110, 1, N'Active'),
   (N'H3', N'Pond Side', 95, 0, N'Available'),
   (N'H4', N'Pond Side', 105, 1, N'Active'),
   (N'H5', N'Pond Side', 90, 0, N'Maintenance'),
   (N'J1', N'Entrance Garden', 70, 1, N'Active'),
   (N'J2', N'Entrance Garden', 75, 0, N'Active'),
   (N'J3', N'Entrance Garden', 65, 1, N'Available'),
   (N'J4', N'Entrance Garden', 80, 0, N'Active');

/******************************************************************************/

MERGE Garden.Plots T
USING
(
   SELECT S.PlotTag,
          S.LocationDescription,
          S.SizeSqFt,
          S.IsRaisedBed,
          PS.StatusId
   FROM @PlotStaging S
   INNER JOIN Garden.PlotStatus PS ON PS.StatusName = S.StatusName
) S ON S.PlotTag = T.PlotTag
WHEN MATCHED AND
   (
      T.LocationDescription <> S.LocationDescription
      OR T.SizeSqFt <> S.SizeSqFt
      OR T.IsRaisedBed <> S.IsRaisedBed
      OR T.StatusId <> S.StatusId
   ) THEN
   UPDATE
   SET LocationDescription = S.LocationDescription,
       SizeSqFt = S.SizeSqFt,
       IsRaisedBed = S.IsRaisedBed,
       StatusId = S.StatusId
WHEN NOT MATCHED THEN
   INSERT(PlotTag, LocationDescription, SizeSqFt, IsRaisedBed, StatusId)
   VALUES(S.PlotTag, S.LocationDescription, S.SizeSqFt, S.IsRaisedBed, S.StatusId);
