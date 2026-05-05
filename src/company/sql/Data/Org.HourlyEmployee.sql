DECLARE @HourlyEmployeeStaging TABLE
(
    EmployeeID INT NOT NULL,
    HourlyPay DECIMAL(18, 2) NOT NULL,
    MaxHoursPerWeek INT NULL
);

INSERT @HourlyEmployeeStaging(EmployeeID, HourlyPay, MaxHoursPerWeek)
VALUES
(1, 17.5, 20),
(9, 43.5, 20),
(26, 44.5, 20),
(27, 33, 20),
(28, 20, 40),
(48, 39.5, 20),
(52, 34.5, 20),
(57, 45.5, 20);

MERGE [Org].HourlyEmployee T
USING (SELECT * FROM @HourlyEmployeeStaging) S ON T.EmployeeID = S.EmployeeID
WHEN MATCHED AND (T.HourlyPay <> S.HourlyPay OR T.MaxHoursPerWeek <> S.MaxHoursPerWeek
) THEN
    UPDATE SET
        HourlyPay = S.HourlyPay,
        MaxHoursPerWeek = S.MaxHoursPerWeek
WHEN NOT MATCHED THEN
    INSERT(EmployeeID, HourlyPay, MaxHoursPerWeek)
    VALUES(S.EmployeeID, S.HourlyPay, S.MaxHoursPerWeek);
