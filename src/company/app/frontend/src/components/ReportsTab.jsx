import { useState } from 'react';
import DataTable from 'react-data-table-component';
import { reportsApi } from '../api/reports';
import SearchBar from './SearchBar';

// Check if DataTable is the default export or named export
const Table = DataTable.default ? DataTable.default : DataTable;

const ReportsTab = () => {
    const [activeReport, setActiveReport] = useState('salary-ranks');
    const [data, setData] = useState([]);
    const [aggregatedData, setAggregatedData] = useState([]);
    const [detailedData, setDetailedData] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [filterText, setFilterText] = useState('');

    // Parameters state
    const [hireDate, setHireDate] = useState(new Date("05/05/2022").toISOString().split('T')[0]);
    const [terminationDate, setTerminationDate] = useState(new Date(new Date().setFullYear(new Date().getFullYear() - 1)).toISOString().split('T')[0]);

    const reports = [
        { id: 'salary-ranks', label: 'Department Salary Ranks' },
        { id: 'terminated-hourly', label: 'Top Hourly Pay of Terminated Employees' },
        { id: 'unhired-managers', label: 'Unhired Employees Assigned to Managers' },
        { id: 'highest-ceo', label: 'Highest Paid CEO' },
        { id: 'salary-ranks-aggregated', label: 'Department Salary Ranks' },
        { id: 'terminated-hourly-aggregated', label: 'Terminated Hourly' },
        { id: 'unhired-managers-aggregated', label: 'Unhired Employees' },
        { id: 'highest-ceo-aggregated', label: 'Highest Paid CEO' }
    ];

    const fetchReport = async () => {
        setLoading(true);
        setError(null);
        setData([]);
        setAggregatedData([]);
        setDetailedData([]);
        try {
            const isAggregated = activeReport.includes('-aggregated');
            if (isAggregated) {
                let aggResult;
                switch (activeReport) {
                    case 'salary-ranks-aggregated':
                        aggResult = await reportsApi.GetDepartmentSalary_Aggregated(hireDate);
                        break;
                    case 'terminated-hourly-aggregated':
                        aggResult = await reportsApi.getTopTerminatedHourly_Aggregated(terminationDate);
                        break;
                    case 'unhired-managers-aggregated':
                        aggResult = await reportsApi.getUnhiredWithManager_Aggregated();
                        break;
                    case 'highest-ceo-aggregated':
                        aggResult = await reportsApi.getHighestPaidCEO_Aggregated();
                        break;
                    default:
                        aggResult = [];
                }
                setAggregatedData(aggResult);

                let detailResult;
                switch (activeReport) {
                    case 'salary-ranks-aggregated':
                        detailResult = await reportsApi.getDepartmentSalaryRanks(hireDate);
                        break;
                    case 'terminated-hourly-aggregated':
                        detailResult = await reportsApi.getTopTerminatedHourly(terminationDate);
                        break;
                    case 'unhired-managers-aggregated':
                        detailResult = await reportsApi.getUnhiredWithManager();
                        break;
                    case 'highest-ceo-aggregated':
                        detailResult = await reportsApi.getHighestPaidCEO();
                        detailResult = detailResult ? [detailResult] : [];
                        break;
                    default:
                        detailResult = [];
                }
                setDetailedData(detailResult);
            } else {
                let result;
                switch (activeReport) {
                    case 'salary-ranks':
                        result = await reportsApi.getDepartmentSalaryRanks(hireDate);
                        break;
                    case 'terminated-hourly':
                        result = await reportsApi.getTopTerminatedHourly(terminationDate);
                        break;
                    case 'unhired-managers':
                        result = await reportsApi.getUnhiredWithManager();
                        break;
                    case 'highest-ceo':
                        result = await reportsApi.getHighestPaidCEO();
                        result = result ? [result] : [];
                        break;
                    default:
                        result = [];
                }
                setData(result);
            }
        } catch (err) {
            setError(err.response?.data || err.message || 'Error fetching report');
        } finally {
            setLoading(false);
        }
    };

    const handleReportChange = (reportId) => {
        setActiveReport(reportId);
        setData([]);
        setAggregatedData([]);
        setDetailedData([]);
        setError(null);
        setFilterText('');
        // We do NOT auto-fetch for parameter-required reports to allow users to set the date first
        if (reportId === 'unhired-managers' || reportId === 'highest-ceo' || reportId === 'unhired-managers-aggregated' || reportId === 'highest-ceo-aggregated') {
            // Auto-fetch these since they have no params
            setTimeout(() => {
                const fetchWithoutParams = async () => {
                    setLoading(true);
                    try {
                        if (reportId === 'unhired-managers') {
                            const res = await reportsApi.getUnhiredWithManager();
                            setData(res);
                        } else if (reportId === 'unhired-managers-aggregated') {
                            const aggRes = await reportsApi.getUnhiredWithManager_Aggregated();
                            setAggregatedData(aggRes);
                            const detRes = await reportsApi.getUnhiredWithManager();
                            console.log(detRes)
                            setDetailedData(detRes);
                        } else if (reportId === 'highest-ceo-aggregated') {
                            const aggRes = await reportsApi.getHighestPaidCEO_Aggregated();
                            setAggregatedData(aggRes);
                            const detRes = await reportsApi.getHighestPaidCEO();
                            setDetailedData(detRes ? [detRes] : []);
                        } else {
                            const res = [await reportsApi.getHighestPaidCEO()].filter(Boolean);
                            setData(res);
                        }
                    } catch (e) {
                        setError(e.message);
                    }
                    setLoading(false);
                };
                fetchWithoutParams();
            }, 0);
        }
    };

    // Define columns based on active report
    const getColumns = () => {
        switch (activeReport) {
            case 'salary-ranks':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Rank', selector: row => row.salary_rank, sortable: true },
                    { name: 'Employee', selector: row => row.employee_name, sortable: true },
                    { name: 'Salary', selector: row => `$${row.employee_salary?.toLocaleString()}`, sortable: true },
                    { name: 'Hire Month/Year', selector: row => `${row.hire_month}/${row.hire_year}`, sortable: true },
                ];
            case 'terminated-hourly':
                return [
                    { name: 'CompanyName', selector: row => row.company_name, sortable: true },
                    { name: 'DepartmentName', selector: row => row.department_name, sortable: true },
                    { name: 'Employee', selector: row => row.employee_name, sortable: true },
                    { name: 'Hourly Pay', selector: row => `$${row.hourly_pay?.toFixed(2)}`, sortable: true },
                    { name: 'Date Terminated', selector: row => row.date_terminated, sortable: true },
                ];
            case 'unhired-managers':
                return [
                    { name: 'CompanyName', selector: row => row.company_name, sortable: true },
                    { name: 'DepartmentName', selector: row => row.department_name, sortable: true },
                    { name: 'Employee', selector: row => row.employee_name, sortable: true },
                    { name: 'Manager Assigned', selector: row => row.manager_assigned, sortable: true },
                ];
            case 'highest-ceo':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'CEO Name', selector: row => row.ceo_name, sortable: true },
                    { name: 'Salary', selector: row => `$${row.ceo_salary?.toLocaleString()}`, sortable: true },
                ];
            case 'salary-ranks-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Employee Count', selector: row => row.employee_count, sortable: true },
                    { name: 'Average Salary', selector: row => `$${row.average_salary?.toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2})}`, sortable: true },
                    { name: 'Highest Salary', selector: row => `$${row.highest_salary?.toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2})}`, sortable: true },
                    { name: 'Lowest Salary', selector: row => `$${row.lowest_salary?.toLocaleString(undefined, {minimumFractionDigits: 2, maximumFractionDigits: 2})}`, sortable: true },
                ];
            case 'terminated-hourly-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Terminated Count', selector: row => row.terminated_count, sortable: true },
                ];
            case 'unhired-managers-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Unhired Count', selector: row => row.unhired_count, sortable: true },
                    { name: 'Manager Names', selector: row => row.manager_names, sortable: true },
                ];
            case 'highest-ceo-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'CEO Count', selector: row => row.ceo_count, sortable: true },
                    { name: 'Highest Salary', selector: row => `$${row.highest_ceo_salary?.toLocaleString()}`, sortable: true },
                ];
            default:
                return [];
        }
    };

    const getDetailedColumns = () => {
        switch (activeReport) {
            case 'salary-ranks-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Rank', selector: row => row.salary_rank, sortable: true },
                    { name: 'Employee', selector: row => row.employee_name, sortable: true },
                    { name: 'Salary', selector: row => `$${row.employee_salary?.toLocaleString()}`, sortable: true },
                    { name: 'Hire Month/Year', selector: row => `${row.hire_month}/${row.hire_year}`, sortable: true },
                ];
            case 'terminated-hourly-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Employee', selector: row => row.employee_name, sortable: true },
                    { name: 'Hourly Pay', selector: row => `$${row.hourly_pay?.toFixed(2)}`, sortable: true },
                    { name: 'Date Terminated', selector: row => row.date_terminated, sortable: true },
                ];
            case 'unhired-managers-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'Department', selector: row => row.department_name, sortable: true },
                    { name: 'Employee', selector: row => row.employee_name, sortable: true },
                    { name: 'Manager Assigned', selector: row => row.manager_assigned, sortable: true },
                ];
            case 'highest-ceo-aggregated':
                return [
                    { name: 'Company', selector: row => row.company_name, sortable: true },
                    { name: 'CEO Name', selector: row => row.ceo_name, sortable: true },
                    { name: 'Salary', selector: row => `$${row.ceo_salary?.toLocaleString()}`, sortable: true },
                ];
            default:
                return [];
        }
    };

    // Global Filter Logic
    const filteredData = data.filter(item => {
        if (!filterText) return true;
        return Object.values(item).some(val => 
            String(val).toLowerCase().includes(filterText.toLowerCase())
        );
    });

    return (
        <div>
            <div className="sub-tabs-container">
                {reports.map(report => (
                    <button
                        key={report.id}
                        className={`sub-tab-button ${activeReport === report.id ? 'active' : ''}`}
                        onClick={() => handleReportChange(report.id)}
                    >
                        {report.label}
                    </button>
                ))}
            </div>

            {/* Parameter Inputs for specific reports */}
            {(activeReport === 'salary-ranks' || activeReport === 'salary-ranks-aggregated' || activeReport === 'terminated-hourly' || activeReport === 'terminated-hourly-aggregated') && (
                <div className="report-filters">
                    {(activeReport === 'salary-ranks' || activeReport === 'salary-ranks-aggregated') && (
                        <div className="filter-group">
                            <label>Hired After Date:</label>
                            <input 
                                type="date" 
                                value={hireDate} 
                                onChange={(e) => setHireDate(e.target.value)} 
                            />
                        </div>
                    )}
                    
                    {activeReport === 'terminated-hourly' && (
                        <div className="filter-group">
                            <label>Terminated After Date:</label>
                            <input 
                                type="date" 
                                value={terminationDate} 
                                onChange={(e) => setTerminationDate(e.target.value)} 
                            />
                        </div>
                    )}

                    {activeReport === 'terminated-hourly-aggregated' && (
                        <div className="filter-group">
                            <label>Terminated After Date:</label>
                            <input 
                                type="date" 
                                value={terminationDate} 
                                onChange={(e) => setTerminationDate(e.target.value)} 
                            />
                        </div>
                    )}

                    <button className="btn btn-primary" onClick={fetchReport} disabled={loading} style={{ height: '36px', alignSelf: 'flex-end' }}>
                        {loading ? 'Loading...' : 'Generate Report'}
                    </button>
                </div>
            )}

            {error && <div style={{ color: 'red', margin: '10px 0' }}>{error}</div>}

            {(() => {
                const isAggregated = activeReport.includes('-aggregated');
                // const aggregatedReportId = activeReport;
                const reportLabel = reports.find(r => r.id === activeReport)?.label;
                const aggFilteredData = aggregatedData.filter(item => {
                    if (!filterText) return true;
                    return Object.values(item).some(val =>
                        String(val).toLowerCase().includes(filterText.toLowerCase())
                    );
                });
                const detFilteredData = detailedData.filter(item => {
                    if (!filterText) return true;
                    return Object.values(item).some(val =>
                        String(val).toLowerCase().includes(filterText.toLowerCase())
                    );
                });

                if (isAggregated) {
                    return (
                        <>
                            <div className="table-header-container">
                                <h2>{reportLabel}</h2>
                                <SearchBar
                                    filterText={filterText}
                                    onFilter={e => setFilterText(e.target.value)}
                                    placeholderText="Filter results..."
                                />
                            </div>

                            <div style={{ marginBottom: '20px' }}>
                                <h3 style={{ margin: '0 0 10px 0', color: '#555' }}>Summary</h3>
                                <Table
                                    columns={getColumns()}
                                    data={aggFilteredData}
                                    progressPending={loading}
                                    pagination
                                    highlightOnHover
                                    pointerOnHover
                                    customStyles={{
                                        headRow: {
                                            style: {
                                                backgroundColor: '#e8f4f8',
                                                fontWeight: 'bold',
                                            }
                                        }
                                    }}
                                />
                            </div>

                            <div>
                                <h3 style={{ margin: '0 0 10px 0', color: '#555' }}>Details</h3>
                                <Table
                                    columns={getDetailedColumns()}
                                    data={detFilteredData}
                                    progressPending={loading}
                                    pagination
                                    highlightOnHover
                                    pointerOnHover
                                    customStyles={{
                                        headRow: {
                                            style: {
                                                backgroundColor: '#f8f9fa',
                                                fontWeight: 'bold',
                                            }
                                        }
                                    }}
                                />
                            </div>
                        </>
                    );
                }

                return (
                    <>
                        <div className="table-header-container">
                            <h2>{reportLabel}</h2>
                            <SearchBar
                                filterText={filterText}
                                onFilter={e => setFilterText(e.target.value)}
                                placeholderText="Filter results..."
                            />
                        </div>

                        <Table
                            columns={getColumns()}
                            data={filteredData}
                            progressPending={loading}
                            pagination
                            highlightOnHover
                            pointerOnHover
                            customStyles={{
                                headRow: {
                                    style: {
                                        backgroundColor: '#f8f9fa',
                                        fontWeight: 'bold',
                                    }
                                }
                            }}
                        />
                    </>
                );
            })()}
        </div>
    );
};

export default ReportsTab;