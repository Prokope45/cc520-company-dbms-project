DECLARE @PlotAssignmentStaging TABLE
(
   Email NVARCHAR(128) NOT NULL,
   PlotTag NVARCHAR(8) NOT NULL,
   StartDate DATE NOT NULL,
   EndDate DATE NULL,
   Notes NVARCHAR(256) NULL,

   PRIMARY KEY(Email, PlotTag, StartDate)
);

/***************************** Modify values here *****************************/
-- The following data was gnerated by Claude Opus 4.6
INSERT @PlotAssignmentStaging(Email, PlotTag, StartDate, EndDate, Notes)
VALUES
   (N'ada@example.com', N'A1', '2026-04-01', NULL, N'Overgrown, needs work'),
   (N'ada@example.com', N'F1', '2026-06-01', NULL, N'Herb garden project'),
   (N'grace@example.com', N'A2', '2026-04-10', NULL, N'Prefers tomatoes and beans'),
   (N'grace@example.com', N'D1', '2026-05-20', NULL, N'Sunflower experiment'),
   (N'alan@example.com', N'B1', '2026-05-01', '2026-10-30', N'Seasonal assignment'),
   (N'alan@example.com', N'E1', '2026-06-15', NULL, N'Root vegetables'),
   (N'katherine@example.com', N'C3', '2026-05-15', NULL, N'Running pollinator section'),
   (N'katherine@example.com', N'G1', '2026-07-01', NULL, N'Wildflower meadow'),
   (N'margaret@example.com', N'A3', '2026-04-20', NULL, N'Squash and pumpkins'),
   (N'linus@example.com', N'B2', '2026-06-01', NULL, N'Kernel corn rows'),
   (N'linus@example.com', N'C1', '2026-06-01', NULL, N'Peppers and eggplant'),
   (N'barbara@example.com', N'B3', '2026-06-10', NULL, N'Lettuce varieties'),
   (N'barbara@example.com', N'H1', '2026-07-15', NULL, N'Water-loving plants'),
   (N'donald@example.com', N'A5', '2026-06-22', NULL, N'Algorithmic planting pattern'),
   (N'hedy@example.com', N'C2', '2026-07-04', NULL, N'Frequency-hopping sprinklers'),
   (N'hedy@example.com', N'F2', '2026-07-04', NULL, N'Greenhouse herbs'),
   (N'tim@example.com', N'D2', '2026-07-15', NULL, N'Interconnected vine trellis'),
   (N'tim@example.com', N'J1', '2026-08-01', NULL, N'Welcome garden display'),
   (N'radia@example.com', N'B5', '2026-07-28', NULL, N'Spanning tree of beans'),
   (N'dennis@example.com', N'C4', '2026-08-05', NULL, N'C-style rows'),
   (N'dennis@example.com', N'E2', '2026-08-05', NULL, N'Pointer herbs'),
   (N'frances@example.com', N'D4', '2026-08-18', NULL, N'Optimized crop rotation'),
   (N'john@example.com', N'C6', '2026-08-30', NULL, N'Lisp-shaped garden beds'),
   (N'john@example.com', N'G2', '2026-09-01', NULL, N'AI-planned layout'),
   (N'anita@example.com', N'D5', '2026-09-10', NULL, N'Community outreach plot'),
   (N'anita@example.com', N'H2', '2026-09-10', NULL, N'Medicinal herbs'),
   (N'vint@example.com', N'E4', '2026-09-22', NULL, N'Networked irrigation test'),
   (N'vint@example.com', N'J2', '2026-09-22', NULL, N'Entrance path flowers'),
   (N'sister@example.com', N'E5', '2026-10-01', NULL, N'Prayer garden herbs'),
   (N'ken@example.com', N'F3', '2026-10-14', NULL, N'Unix thyme and sage'),
   (N'ken@example.com', N'G5', '2026-10-14', NULL, N'Hillside terracing'),
   (N'feifei@example.com', N'F5', '2026-10-28', NULL, N'Vision garden with color patterns'),
   (N'edsger@example.com', N'G4', '2026-11-05', NULL, N'Shortest-path walkways'),
   (N'edsger@example.com', N'H4', '2026-11-05', NULL, N'Structured beds'),
   (N'jean@example.com', N'J4', '2026-11-18', NULL, N'Historic flower varieties'),
   (N'guido@example.com', N'A5', '2025-01-01', '2026-06-21', N'Previous season assignment'),
   (N'lynn@example.com', N'B2', '2025-01-01', '2026-05-31', N'Chip-design garden layout'),
   (N'james@example.com', N'D2', '2025-02-01', '2026-07-14', N'Java beans—literally'),
   (N'adele@example.com', N'E1', '2025-03-01', '2026-06-14', N'Smalltalk herb circle'),
   (N'bjarne@example.com', N'C1', '2025-02-15', '2026-05-31', N'Plus-plus yield experiment'),
   (N'sophie@example.com', N'F2', '2025-04-01', '2026-07-03', N'ARM-length raised beds'),
   (N'brian@example.com', N'G2', '2025-03-15', '2026-08-31', N'Hello world garden'),
   (N'shafi@example.com', N'H1', '2025-05-01', '2026-07-14', N'Zero-knowledge compost'),
   (N'anders@example.com', N'J1', '2025-04-10', '2026-07-31', N'Type-safe tulips');

/******************************************************************************/

MERGE Garden.PlotAssignments T
USING
(
   SELECT G.GardenerId,
          P.PlotId,
          S.StartDate,
          S.EndDate,
          S.Notes
   FROM @PlotAssignmentStaging S
   INNER JOIN Garden.Gardeners G ON G.Email = S.Email
   INNER JOIN Garden.Plots P ON P.PlotTag = S.PlotTag
) S ON S.GardenerId = T.GardenerId
   AND S.PlotId = T.PlotId
   AND S.StartDate = T.StartDate
WHEN MATCHED AND
   (
      ISNULL(T.EndDate, '1900-01-01') <> ISNULL(S.EndDate, '1900-01-01')
      OR ISNULL(T.Notes, N'') <> ISNULL(S.Notes, N'')
   ) THEN
   UPDATE
   SET T.EndDate = S.EndDate,
       T.Notes = S.Notes,
       T.UpdatedOn = SYSDATETIMEOFFSET()
WHEN NOT MATCHED THEN
   INSERT(PlotId, GardenerId, StartDate, EndDate, Notes)
   VALUES(S.PlotId, S.GardenerId, S.StartDate, S.EndDate, S.Notes);