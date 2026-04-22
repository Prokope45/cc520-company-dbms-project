IF OBJECT_ID(N'Org.HourlyEmployee') IS NULL
BEGIN
    CREATE TABLE Org.[HourlyEmployee] (
        [EmployeeID] INT NOT NULL PRIMARY KEY,
        [HourlyPay] DECIMAL(18, 2) NOT NULL,
        [MaxHoursPerWeek] INT NULL,
        CONSTRAINT [FK_HourlyEmployee_Employee] FOREIGN KEY ([EmployeeID])
            REFERENCES Org.[Employee]([EmployeeID])
            ON DELETE CASCADE,
        CONSTRAINT [CK_HourlyPay_Positive] CHECK ([HourlyPay] > 0)
    );
END;