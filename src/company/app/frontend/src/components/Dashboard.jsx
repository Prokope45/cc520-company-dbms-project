import React, { useState } from 'react';
import CompanyTable from './CompanyTable';
import DepartmentTable from './DepartmentTable';
import EmployeeTable from './EmployeeTable';
import ReportsTab from './ReportsTab';

const Dashboard = () => {
  const [activeTab, setActiveTab] = useState('companies');

  const tabs = [
    { id: 'companies', label: 'Companies' },
    { id: 'departments', label: 'Departments' },
    { id: 'employees', label: 'Employees' },
    { id: 'reports', label: 'Reports' },
  ];

  return (
    <div className="container">
      <h1>Company Management Dashboard</h1>
      
      <div className="tabs-container">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            className={`tab-button ${activeTab === tab.id ? 'active' : ''}`}
            onClick={() => setActiveTab(tab.id)}
          >
            {tab.label}
          </button>
        ))}
      </div>

      <div className="tab-content">
        {activeTab === 'companies' && <CompanyTable />}
        {activeTab === 'departments' && <DepartmentTable />}
        {activeTab === 'employees' && <EmployeeTable />}
        {activeTab === 'reports' && <ReportsTab />}
      </div>
    </div>
  );
};

export default Dashboard;