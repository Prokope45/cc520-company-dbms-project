IF OBJECT_ID(N'Org.DepEmpRole') IS NULL
BEGIN
    CREATE TABLE Org.[DepEmpRole]
    (
        [RoleID] INT NOT NULL,
        [DepartmentID] INT NOT NULL,
        [EmployeeID] INT NOT NULL,
        PRIMARY KEY ([RoleID], [DepartmentID], [EmployeeID]),
        CONSTRAINT [FK_DepEmpRole_Role] FOREIGN KEY ([RoleID])
            REFERENCES Org.[Role]([RoleID]),
        CONSTRAINT [FK_DepEmpRole_Department] FOREIGN KEY ([DepartmentID])
            REFERENCES Org.[Department]([DepartmentID]),
        CONSTRAINT [FK_DepEmpRole_Employee] FOREIGN KEY ([EmployeeID])
            REFERENCES Org.[Employee]([EmployeeID])
    );
END;
