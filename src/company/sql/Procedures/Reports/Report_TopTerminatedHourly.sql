CREATE PROCEDURE [Org].[sp_Report_TopTerminatedHourly]
    @TerminationDate DATETIME2
AS
BEGIN
    SELECT TOP(5)
        CONVERT(VARCHAR(10), E.TerminationDate, 120) AS DateTerminated,
        HE.HourlyPay,
        CONCAT(P.FirstName, ' ', P.LastName) AS EmployeeName
    FROM [Org].[Employee] E
        INNER JOIN [Org].[Person] P ON P.PersonID = E.PersonID
        INNER JOIN [Org].[HourlyEmployee] HE ON HE.EmployeeID E.EmployeeID
    WHERE E.TerminationDate IS NOT NULL 
      AND E.TerminationDate >= @TerminationDate
    ORDER BY HE.HourlyPay DESC, E.TerminationDate DESC;
END
GO
