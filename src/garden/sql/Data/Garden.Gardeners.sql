DECLARE @GardenerStaging TABLE
(
   FirstName NVARCHAR(32) NOT NULL,
   LastName NVARCHAR(32) NOT NULL,
   Phone NVARCHAR(24) NULL,
   Email NVARCHAR(128) NOT NULL PRIMARY KEY,
   JoinDate DATE NOT NULL
);

/***************************** Modify values here *****************************/
-- The following data was gnerated by Claude Opus 4.6
INSERT @GardenerStaging(FirstName, LastName, Phone, Email, JoinDate)
VALUES
   (N'Ada', N'Lovelace', N'555-0101', N'ada@example.com', '2024-03-12'),
   (N'Grace', N'Hopper', N'555-0102', N'grace@example.com', '2024-04-03'),
   (N'Alan', N'Turing', N'555-0103', N'alan@example.com', '2024-04-18'),
   (N'Katherine', N'Johnson', N'555-0104', N'katherine@example.com', '2024-05-02'),
   (N'Margaret', N'Hamilton', N'555-0105', N'margaret@example.com', '2024-05-17'),
   (N'Linus', N'Torvalds', N'555-0106', N'linus@example.com', '2024-06-01'),
   (N'Barbara', N'Liskov', N'555-0107', N'barbara@example.com', '2024-06-10'),
   (N'Donald', N'Knuth', N'555-0108', N'donald@example.com', '2024-06-22'),
   (N'Hedy', N'Lamarr', N'555-0109', N'hedy@example.com', '2024-07-04'),
   (N'Tim', N'Berners-Lee', N'555-0110', N'tim@example.com', '2024-07-15'),
   (N'Radia', N'Perlman', N'555-0111', N'radia@example.com', '2024-07-28'),
   (N'Dennis', N'Ritchie', N'555-0112', N'dennis@example.com', '2024-08-05'),
   (N'Frances', N'Allen', N'555-0113', N'frances@example.com', '2024-08-18'),
   (N'John', N'McCarthy', N'555-0114', N'john@example.com', '2024-08-30'),
   (N'Anita', N'Borg', N'555-0115', N'anita@example.com', '2024-09-10'),
   (N'Vint', N'Cerf', N'555-0116', N'vint@example.com', '2024-09-22'),
   (N'Sister', N'Keller', N'555-0117', N'sister@example.com', '2024-10-01'),
   (N'Ken', N'Thompson', N'555-0118', N'ken@example.com', '2024-10-14'),
   (N'Fei-Fei', N'Li', N'555-0119', N'feifei@example.com', '2024-10-28'),
   (N'Edsger', N'Dijkstra', N'555-0120', N'edsger@example.com', '2024-11-05'),
   (N'Jean', N'Bartik', N'555-0121', N'jean@example.com', '2024-11-18'),
   (N'Guido', N'van Rossum', N'555-0122', N'guido@example.com', '2024-12-01'),
   (N'Lynn', N'Conway', N'555-0123', N'lynn@example.com', '2024-12-12'),
   (N'James', N'Gosling', N'555-0124', N'james@example.com', '2025-01-05'),
   (N'Adele', N'Goldberg', N'555-0125', N'adele@example.com', '2025-01-15'),
   (N'Bjarne', N'Stroustrup', N'555-0126', N'bjarne@example.com', '2025-02-02'),
   (N'Sophie', N'Wilson', N'555-0127', N'sophie@example.com', '2025-02-18'),
   (N'Brian', N'Kernighan', N'555-0128', N'brian@example.com', '2025-03-01'),
   (N'Shafi', N'Goldwasser', N'555-0129', N'shafi@example.com', '2025-03-14'),
   (N'Anders', N'Hejlsberg', N'555-0130', N'anders@example.com', '2025-03-28');

/******************************************************************************/

MERGE Garden.Gardeners T
USING @GardenerStaging S ON S.Email = T.Email
WHEN MATCHED AND
   (
      T.FirstName <> S.FirstName
      OR T.LastName <> S.LastName
      OR ISNULL(T.Phone, N'') <> ISNULL(S.Phone, N'')
      OR ISNULL(T.JoinDate, '1900-01-01') <> S.JoinDate
   ) THEN
   UPDATE
   SET T.FirstName = S.FirstName,
       T.LastName = S.LastName,
       T.Phone = S.Phone,
       T.JoinDate = S.JoinDate,
       T.UpdatedOn = SYSDATETIMEOFFSET()
WHEN NOT MATCHED THEN
   INSERT(FirstName, LastName, Phone, Email, JoinDate)
   VALUES(S.FirstName, S.LastName, S.Phone, S.Email, S.JoinDate);