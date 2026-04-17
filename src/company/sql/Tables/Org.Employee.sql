IF OBJECT_ID(N'Org.Employee') IS NULL
BEGIN
    CREATE TABLE Org.[Employee]
    (
        [EmployeeID] INT NOT NULL PRIMARY KEY IDENTITY(1,1),
        [ManagerID] INT NULL,
        [PersonID] INT NOT NULL,
        [EmployeeTypeID] INT NOT NULL,
        [HireDate] DATETIME2 NULL,
        [TerminationDate] DATETIME2 NULL,
        -- CONSTRAINT [FK_Employee_Manager] FOREIGN KEY ([ManagerID])
        --     REFERENCES [Employee]([EmployeeID])
        --     ON DELETE SET NULL,
        --     ON UPDATE CASCADE,
        CONSTRAINT [FK_Employee_Person] FOREIGN KEY ([PersonID])
            REFERENCES [Person]([PersonID])
        -- CONSTRAINT [FK_Employee_EmployeeType] FOREIGN KEY ([EmployeeTypeID])
        --     REFERENCES [EmployeeType]([EmployeeTypeID])
        -- CONSTRAINT [UK_Employee_PersonID] UNIQUE ([PersonID]),
        -- CONSTRAINT [UK_Employee_EmployeeID] UNIQUE ([EmployeeID]),
        -- CONSTRAINT [CK_Employee_HireDate_NotNull_Salaried] CHECK (
        --     ([EmployeeTypeID] = 1 AND [HireDate] IS NOT NULL) OR
        --     ([EmployeeTypeID] = 2 AND [HireDate] IS NOT NULL) OR
        --     ([HireDate] IS NULL)
        -- ),
        -- CONSTRAINT [CK_Employee_TerminationDate_NotNull] CHECK (
        --     ([TerminationDate] IS NULL) OR
        --     ([TerminationDate] > [HireDate])
        -- ),
        -- CONSTRAINT [CK_Employee_Manager_NotSelf] CHECK (
        --     [ManagerID] IS NULL OR [ManagerID] <> [EmployeeID]
        -- )
    );
END;