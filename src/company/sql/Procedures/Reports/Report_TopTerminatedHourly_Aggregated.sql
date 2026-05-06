CREATE PROCEDURE [Org].[sp_Report_TopTerminatedHourly_Aggregated]
    @TerminationDate DATETIME2
AS
BEGIN
    SELECT TOP(5)
        C.[Name] AS CompanyName,
        D.[Name] AS DepartmentName,
        COUNT(*) AS TerminatedCount,
        AVG(HE.HourlyPay) AS AvgTerminatedHourlyPay,
        MAX(E.TerminationDate) AS LatestTerminationDate
    FROM [Org].[Employee] E
        INNER JOIN [Org].[HourlyEmployee] HE ON HE.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[DepEmpRole] DER ON DER.EmployeeID = E.EmployeeID
        INNER JOIN [Org].[Department] D ON D.DepartmentID = DER.DepartmentID
        INNER JOIN [Org].[Company] C ON C.CompanyID = D.CompanyID
    WHERE E.TerminationDate IS NOT NULL 
      AND E.TerminationDate >= @TerminationDate
    GROUP BY C.[Name], D.[Name]
    ORDER BY AvgTerminatedHourlyPay DESC, LatestTerminationDate DESC;
END
GO
