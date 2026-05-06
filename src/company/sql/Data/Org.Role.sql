DECLARE @RoleStaging TABLE
(
    [Name] NVARCHAR(255) NOT NULL,
    [Description] NVARCHAR(255) NOT NULL
);

INSERT @RoleStaging([Name], [Description])
VALUES
(N'Intern', N'Junior intern; reports to Manager'),
(N'Developer I', N'Junior developer; reports to Manger'),
(N'Developer II', N'Mid-level developer; reports to Manger'),
(N'Developer III', N'Senior developer; reports to Manger'),
(N'Senior Engineer', N'Senior engineer; reports to Manager'),
(N'Member', N'Regular team member; Reports to Manager'),
(N'Manager', N'Department manager; Reports to Director'),
(N'HR Manager', N'Human resources manager; Reports to CHRO'),
(N'Director', N'Senior manager; reports to Executive in Department'),
(N'CTO', N'Chief Technology Officer; reports to CEO'),
(N'CFO', N'Chief Financial Officer; reports to CEO'),
(N'CHRO', N'Chief Human Resources Officer; reports to CEO'),
(N'CEO', N'Chief Executive Officer');

MERGE [Org].[Role] T
USING (SELECT * FROM @RoleStaging) S
    ON T.[Name] = S.[Name]
WHEN MATCHED AND (
    T.Description <> S.Description
) THEN
    UPDATE SET [Description] = S.Description
WHEN NOT MATCHED THEN
    INSERT([Name], [Description])
    VALUES(S.[Name], S.Description);
