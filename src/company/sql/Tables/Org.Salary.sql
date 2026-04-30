IF OBJECT_ID(N'Org.Salary') IS NULL
BEGIN
    CREATE TABLE [Org].[Salary] (
        [SalaryID] INT NOT NULL PRIMARY KEY,
        [BaseSalary] DECIMAL(18, 2) NOT NULL,
        [Bonus] DECIMAL(18, 2) DEFAULT 0.00,
        [Deductions] DECIMAL(18, 2) DEFAULT 0.00,
        [EffectiveFrom] DATETIME2 DEFAULT GETDATE(),
        [EffectiveTo] DATETIME2 DEFAULT GETDATE(),
        CONSTRAINT [CK_Salary_BaseSalary_Positive] CHECK ([BaseSalary] >= 0),
        CONSTRAINT [CK_Salary_Bonus_Positive] CHECK ([Bonus] >= 0),
        CONSTRAINT [CK_Salary_Deductions_Positive] CHECK ([Deductions] >= 0),
        CONSTRAINT [CK_Salary_EffectiveDates] CHECK ([EffectiveFrom] < [EffectiveTo])
    );
END;
