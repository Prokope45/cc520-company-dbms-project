IF OBJECT_ID(N'Org.HourlyEmployee') IS NULL
BEGIN
    CREATE TABLE Org.[HourlyEmployee] (
        [EmployeeID] INT NOT NULL PRIMARY KEY,
        [EmployeeTypeID] INT NOT NULL,
        [HourlyPay] DECIMAL(18, 2) NOT NULL,
        [MaxHoursPerWeek] INT NULL,
        CONSTRAINT [FK_HourlyEmployee_Employee] FOREIGN KEY ([EmployeeID])
            REFERENCES Org.[Employee]([EmployeeID])
            ON DELETE CASCADE,
        CONSTRAINT [FK_HourlyEmployee_EmployeeType] FOREIGN KEY ([EmployeeTypeID])
            REFERENCES Org.[EmployeeType]([EmployeeTypeID]),
        CONSTRAINT [CK_HourlyRate_Positive] CHECK ([HourlyPay] > 0)
    );
END;