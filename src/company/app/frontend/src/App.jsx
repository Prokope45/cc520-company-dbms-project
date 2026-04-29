import React, { useState, useEffect } from 'react';
import CompanyTable from './components/CompanyTable';
import DepartmentTable from './components/DepartmentTable';
import './App.css';

function App() {
  return (
    <div className="container">
      <h1>Company Management</h1>
      <CompanyTable />
      <DepartmentTable />
    </div>
  );
}

export default App;
