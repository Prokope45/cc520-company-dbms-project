import React, { useState } from 'react';
import DataTable from 'react-data-table-component';
import { reportsApi } from '../api/reports';
import SearchBar from './SearchBar';

// Check if DataTable is the default export or named export
const Table = DataTable.default ? DataTable.default : DataTable;

const ReportsTab = () => {
    const [activeReport, setActiveReport] = useState('salary-ranks');
    const [data, setData] = useState([]);
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
        { id: 'highest-ceo', label: 'Highest Paid CEO' }
    ];

    const fetchReport = async () => {
        setLoading(true);
        setError(null);
        setData([]);
        try {
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
                    result = result ? [result] : []; // Convert single object to array for DataTable
                    break;
                default:
                    result = [];
            }
            setData(result);
        } catch (err) {
            setError(err.response?.data || err.message || 'Error fetching report');
        } finally {
            setLoading(false);
        }
    };

    const handleReportChange = (reportId) => {
        setActiveReport(reportId);
        setData([]);
        setError(null);
        setFilterText('');
        // We do NOT auto-fetch for parameter-required reports to allow users to set the date first
        if (reportId === 'unhired-managers' || reportId === 'highest-ceo') {
            // But we can auto-fetch these since they have no params
            setTimeout(() => {
                const fetchWithoutParams = async () => {
                    setLoading(true);
                    try {
                        let res = reportId === 'unhired-managers' 
                            ? await reportsApi.getUnhiredWithManager()
                            : [await reportsApi.getHighestPaidCEO()].filter(Boolean);
                        setData(res);
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
            {(activeReport === 'salary-ranks' || activeReport === 'terminated-hourly') && (
                <div className="report-filters">
                    {activeReport === 'salary-ranks' && (
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

                    <button className="btn btn-primary" onClick={fetchReport} disabled={loading} style={{ height: '36px', alignSelf: 'flex-end' }}>
                        {loading ? 'Loading...' : 'Generate Report'}
                    </button>
                </div>
            )}

            {error && <div style={{ color: 'red', margin: '10px 0' }}>{error}</div>}

            <div className="table-header-container">
                <h2>{reports.find(r => r.id === activeReport)?.label}</h2>
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
        </div>
    );
};

export default ReportsTab;