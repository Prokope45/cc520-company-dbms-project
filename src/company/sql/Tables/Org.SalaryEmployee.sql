IF OBJECT_ID(N'Org.SalaryEmployee') IS NULL
BEGIN
    CREATE TABLE Org.[SalaryEmployee] (
        [EmployeeID] INT NOT NULL PRIMARY KEY,
        [PaidTimeOffHours] INT NOT NULL,
        [SickHours] INT NOT NULL,
        CONSTRAINT [FK_SalaryEmployee_Employee] FOREIGN KEY ([EmployeeID])
            REFERENCES Org.[Employee]([EmployeeID])
            ON DELETE CASCADE,
        CONSTRAINT [FK_SalaryEmployee_Salary] FOREIGN KEY ([EmployeeID])
            REFERENCES Org.[Salary]([SalaryID])
            ON DELETE CASCADE,
        CONSTRAINT [CK_PaidTimeOffHours_Positive] CHECK ([PaidTimeOffHours] >= 0),
        CONSTRAINT [CK_SickHours_Positive] CHECK ([SickHours] >= 0)
    );
END;
