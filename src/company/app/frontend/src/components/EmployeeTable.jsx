import React, { useState, useEffect } from 'react';
import { employeeAPI } from '../api';
import Modal from './Modal';

const EmployeeTable = () => {
  const [employees, setEmployees] = useState([]);
  const [errorMessage, setErrorMessage] = useState('');
  
  // For standard add/edit modal
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingEmployee, setEditingEmployee] = useState(null);
  const [formActiveTab, setFormActiveTab] = useState('personal'); // 'personal', 'organization', 'status'
  
  const [formData, setFormData] = useState({
    company_id: '',
    department_id: '',
    department: '',
    first_name: '',
    last_name: '',
    email: '',
    phone: '',
    street: '',
    city: '',
    state: '',
    zip_code: '',
    role_title: '',
    manager_name: '',
    hire_date: '',
    termination_date: '',
    status_type: 'Salary',
    hourly_pay: '',
    max_hours_per_week: '',
    base_salary: '',
    bonus: '',
    deductions: '',
    paid_time_off_hours: '',
    sick_hours: '',
    effective_from: '',
    effective_to: ''
  });

  // For employee details modal (clicking the name)
  const [isDetailsModalOpen, setIsDetailsModalOpen] = useState(false);
  const [selectedEmployee, setSelectedEmployee] = useState(null);
  const [detailsActiveTab, setDetailsActiveTab] = useState('personal'); // 'personal', 'organization', 'status'

  useEffect(() => {
    fetchEmployees();
  }, []);

  const fetchEmployees = async () => {
    try {
      const response = await employeeAPI.getAll();
      setEmployees(response.data || []);
      setErrorMessage('');
    } catch (error) {
      console.error('Error fetching employees:', error);
      setErrorMessage(error.response?.data?.error || 'Failed to fetch employees');
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this employee?')) {
      try {
        await employeeAPI.delete(id);
        fetchEmployees();
        setErrorMessage('');
      } catch (error) {
        console.error('Error deleting employee:', error);
        setErrorMessage(error.response?.data?.error || 'Failed to delete employee');
      }
    }
  };

  const handleEdit = (emp) => {
    setEditingEmployee(emp);
    setFormActiveTab('personal');
    setFormData({
      company_id: emp.company_id || '',
      department_id: emp.department_id || '',
      department: emp.department || '',
      first_name: emp.first_name || '',
      last_name: emp.last_name || '',
      email: emp.email || '',
      phone: emp.phone || '',
      street: emp.street || '',
      city: emp.city || '',
      state: emp.state || '',
      zip_code: emp.zip_code || '',
      role_title: emp.role_title || '',
      manager_name: emp.manager_name || '',
      hire_date: emp.hire_date || '',
      termination_date: emp.termination_date || '',
      status_type: emp.status_type || 'Salary',
      hourly_pay: emp.hourly_pay || '',
      max_hours_per_week: emp.max_hours_per_week || '',
      base_salary: emp.base_salary || '',
      bonus: emp.bonus || '',
      deductions: emp.deductions || '',
      paid_time_off_hours: emp.paid_time_off_hours || '',
      sick_hours: emp.sick_hours || '',
      effective_from: emp.effective_from || '',
      effective_to: emp.effective_to || ''
    });
    setIsModalOpen(true);
  };

  const handleAdd = () => {
    setEditingEmployee(null);
    setFormActiveTab('personal');
    setFormData({
      company_id: '',
      department_id: '',
      department: '',
      first_name: '',
      last_name: '',
      email: '',
      phone: '',
      street: '',
      city: '',
      state: '',
      zip_code: '',
      role_title: '',
      manager_name: '',
      hire_date: '',
      termination_date: '',
      status_type: 'Salary',
      hourly_pay: '',
      max_hours_per_week: '',
      base_salary: '',
      bonus: '',
      deductions: '',
      paid_time_off_hours: '',
      sick_hours: '',
      effective_from: '',
      effective_to: ''
    });
    setIsModalOpen(true);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setErrorMessage('');
    try {
      const payload = { 
        ...formData, 
        company_id: parseInt(formData.company_id, 10),
        department_id: parseInt(formData.department_id, 10)
      };

      // Clean up fields based on status_type
      if (payload.status_type === 'Salary') {
        payload.hourly_pay = 0;
        payload.max_hours_per_week = 0;
        payload.base_salary = parseFloat(payload.base_salary) || 0;
        payload.bonus = parseFloat(payload.bonus) || 0;
        payload.deductions = parseFloat(payload.deductions) || 0;
        payload.paid_time_off_hours = parseInt(payload.paid_time_off_hours, 10) || 0;
        payload.sick_hours = parseInt(payload.sick_hours, 10) || 0;
      } else {
        payload.base_salary = 0;
        payload.bonus = 0;
        payload.deductions = 0;
        payload.paid_time_off_hours = 0;
        payload.sick_hours = 0;
        payload.effective_from = '';
        payload.effective_to = '';
        payload.hourly_pay = parseFloat(payload.hourly_pay) || 0;
        payload.max_hours_per_week = parseInt(payload.max_hours_per_week, 10) || 0;
      }
      
      if (editingEmployee) {
        await employeeAPI.update(editingEmployee.employee_id, payload);
      } else {
        await employeeAPI.create(payload);
      }
      setIsModalOpen(false);
      fetchEmployees();
    } catch (error) {
      console.error('Error saving employee:', error);
      const msg = error.response?.data?.error || error.response?.data?.message || 'Failed to save employee';
      setErrorMessage(msg);
    }
  };

  const handleNameClick = (emp) => {
    setSelectedEmployee(emp);
    setDetailsActiveTab('personal');
    setIsDetailsModalOpen(true);
  };

  const renderTabButtons = (activeTab, setActiveTab) => (
    <div style={{ display: 'flex', borderBottom: '1px solid #ccc', marginBottom: '15px' }}>
      <button 
        type="button"
        style={{ padding: '10px', background: activeTab === 'personal' ? '#e9ecef' : 'none', border: 'none', borderBottom: activeTab === 'personal' ? '2px solid #007bff' : 'none', cursor: 'pointer' }}
        onClick={() => setActiveTab('personal')}
      >
        Personal & Address
      </button>
      <button 
        type="button"
        style={{ padding: '10px', background: activeTab === 'organization' ? '#e9ecef' : 'none', border: 'none', borderBottom: activeTab === 'organization' ? '2px solid #007bff' : 'none', cursor: 'pointer' }}
        onClick={() => setActiveTab('organization')}
      >
        Organization
      </button>
      <button 
        type="button"
        style={{ padding: '10px', background: activeTab === 'status' ? '#e9ecef' : 'none', border: 'none', borderBottom: activeTab === 'status' ? '2px solid #007bff' : 'none', cursor: 'pointer' }}
        onClick={() => setActiveTab('status')}
      >
        Employment Status
      </button>
    </div>
  );

  return (
    <div className="section">
      <h2>Employees</h2>
      <button className="btn btn-primary" onClick={handleAdd}>+ Add Employee</button>

      {errorMessage && (
        <div className="alert alert-danger" style={{ marginTop: '10px', padding: '10px', backgroundColor: '#f8d7da', border: '1px solid #f5c6cb', borderRadius: '4px' }}>
          {errorMessage}
        </div>
      )}

      <table className="display" style={{ width: '100%', marginTop: '15px' }}>
        <thead>
          <tr>
            <th>ID</th>
            <th>First Name</th>
            <th>Last Name</th>
            <th>Department</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {employees.map(emp => (
            <tr key={emp.employee_id}>
              <td className="table-cell-id">{emp.employee_id}</td>
              <td className="table-cell-text">
                <a href="#" onClick={(e) => { e.preventDefault(); handleNameClick(emp); }}>
                  {emp.first_name}
                </a>
              </td>
              <td className="table-cell-text">
                <a href="#" onClick={(e) => { e.preventDefault(); handleNameClick(emp); }}>
                  {emp.last_name}
                </a>
              </td>
              <td className="table-cell-text">{emp.department}</td>
              <td className="table-cell-actions">
                <button className="btn btn-primary" onClick={() => handleEdit(emp)} style={{ marginRight: '5px' }}>Edit</button>
                <button className="btn btn-secondary" onClick={() => handleDelete(emp.employee_id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {/* Edit/Add Modal */}
      <Modal
        isOpen={isModalOpen} 
        onClose={() => setIsModalOpen(false)}
        title={editingEmployee ? "Edit Employee" : "Add Employee"}
      >
        <form onSubmit={handleSubmit}>
          {renderTabButtons(formActiveTab, setFormActiveTab)}

          {formActiveTab === 'personal' && (
            <div className="form-group" style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '10px' }}>
              <div><label>First Name:</label><input type="text" required value={formData.first_name} onChange={(e) => setFormData({...formData, first_name: e.target.value})} /></div>
              <div><label>Last Name:</label><input type="text" required value={formData.last_name} onChange={(e) => setFormData({...formData, last_name: e.target.value})} /></div>
              <div><label>Email:</label><input type="email" value={formData.email} onChange={(e) => setFormData({...formData, email: e.target.value})} /></div>
              <div><label>Phone:</label><input type="text" value={formData.phone} onChange={(e) => setFormData({...formData, phone: e.target.value})} /></div>
              <div style={{ gridColumn: '1 / -1' }}><hr/></div>
              <div><label>Street:</label><input type="text" value={formData.street} onChange={(e) => setFormData({...formData, street: e.target.value})} /></div>
              <div><label>City:</label><input type="text" value={formData.city} onChange={(e) => setFormData({...formData, city: e.target.value})} /></div>
              <div><label>State:</label><input type="text" value={formData.state} onChange={(e) => setFormData({...formData, state: e.target.value})} /></div>
              <div><label>Zip:</label><input type="text" value={formData.zip_code} onChange={(e) => setFormData({...formData, zip_code: e.target.value})} /></div>
            </div>
          )}

          {formActiveTab === 'organization' && (
            <div className="form-group" style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '10px' }}>
              <div><label>Company ID:</label><input type="number" required value={formData.company_id} onChange={(e) => setFormData({...formData, company_id: e.target.value})} /></div>
              <div><label>Department ID:</label><input type="number" required value={formData.department_id} onChange={(e) => setFormData({...formData, department_id: e.target.value})} /></div>
              <div><label>Department Name:</label><input type="text" required value={formData.department} onChange={(e) => setFormData({...formData, department: e.target.value})} /></div>
              <div><label>Role:</label><input type="text" value={formData.role_title} onChange={(e) => setFormData({...formData, role_title: e.target.value})} /></div>
              <div><label>Manager Name:</label><input type="text" value={formData.manager_name} onChange={(e) => setFormData({...formData, manager_name: e.target.value})} /></div>
              <div><label>Hire Date (YYYY-MM-DD):</label><input type="date" value={formData.hire_date} onChange={(e) => setFormData({...formData, hire_date: e.target.value})} /></div>
              <div><label>Termination Date:</label><input type="date" value={formData.termination_date} onChange={(e) => setFormData({...formData, termination_date: e.target.value})} /></div>
            </div>
          )}

          {formActiveTab === 'status' && (
            <div className="form-group" style={{ display: 'flex', flexDirection: 'column', gap: '10px' }}>
              <div>
                <label>Status Type:</label>
                <select value={formData.status_type} onChange={(e) => setFormData({...formData, status_type: e.target.value})}>
                  <option value="Salary">Salary</option>
                  <option value="Hourly">Hourly</option>
                </select>
              </div>

              {formData.status_type === 'Hourly' && (
                <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '10px', marginTop: '10px' }}>
                  <div><label>Hourly Pay:</label><input type="number" step="0.01" value={formData.hourly_pay} onChange={(e) => setFormData({...formData, hourly_pay: e.target.value})} /></div>
                  <div><label>Max Hours Per Week:</label><input type="number" value={formData.max_hours_per_week} onChange={(e) => setFormData({...formData, max_hours_per_week: e.target.value})} /></div>
                </div>
              )}

              {formData.status_type === 'Salary' && (
                <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '10px', marginTop: '10px' }}>
                  <div><label>Base Salary:</label><input type="number" step="0.01" value={formData.base_salary} onChange={(e) => setFormData({...formData, base_salary: e.target.value})} /></div>
                  <div><label>Bonus:</label><input type="number" step="0.01" value={formData.bonus} onChange={(e) => setFormData({...formData, bonus: e.target.value})} /></div>
                  <div><label>Deductions:</label><input type="number" step="0.01" value={formData.deductions} onChange={(e) => setFormData({...formData, deductions: e.target.value})} /></div>
                  <div><label>PTO Hours:</label><input type="number" value={formData.paid_time_off_hours} onChange={(e) => setFormData({...formData, paid_time_off_hours: e.target.value})} /></div>
                  <div><label>Sick Hours:</label><input type="number" value={formData.sick_hours} onChange={(e) => setFormData({...formData, sick_hours: e.target.value})} /></div>
                  <div><label>Effective From:</label><input type="date" value={formData.effective_from} onChange={(e) => setFormData({...formData, effective_from: e.target.value})} /></div>
                  <div><label>Effective To:</label><input type="date" value={formData.effective_to} onChange={(e) => setFormData({...formData, effective_to: e.target.value})} /></div>
                </div>
              )}
            </div>
          )}

          <div style={{ marginTop: '20px', borderTop: '1px solid #eee', paddingTop: '15px' }}>
            <button type="submit" className="btn btn-primary">Save</button>
            <button type="button" className="btn btn-secondary" onClick={() => setIsModalOpen(false)} style={{ marginLeft: '10px' }}>Cancel</button>
          </div>
        </form>
      </Modal>

      {/* Details View Modal */}
      {selectedEmployee && (
        <Modal
          isOpen={isDetailsModalOpen}
          onClose={() => setIsDetailsModalOpen(false)}
          title={`Employee Details: ${selectedEmployee.first_name} ${selectedEmployee.last_name}`}
        >
          {renderTabButtons(detailsActiveTab, setDetailsActiveTab)}

          <div style={{ lineHeight: '1.6', minHeight: '200px' }}>
            {detailsActiveTab === 'personal' && (
              <div>
                <h3 style={{ borderBottom: '1px solid #ccc', paddingBottom: '5px' }}>Contact Info</h3>
                <p><strong>Email:</strong> {selectedEmployee.email}</p>
                <p><strong>Phone:</strong> {selectedEmployee.phone}</p>

                <h3 style={{ borderBottom: '1px solid #ccc', paddingBottom: '5px', marginTop: '15px' }}>Address</h3>
                <p>{selectedEmployee.street}</p>
                <p>{selectedEmployee.city}, {selectedEmployee.state} {selectedEmployee.zip_code}</p>
              </div>
            )}

            {detailsActiveTab === 'organization' && (
              <div>
                <p><strong>Company ID:</strong> {selectedEmployee.company_id}</p>
                <p><strong>Department:</strong> {selectedEmployee.department} (ID: {selectedEmployee.department_id})</p>
                <p><strong>Role:</strong> {selectedEmployee.role_title}</p>
                <p><strong>Manager:</strong> {selectedEmployee.manager_name}</p>
                <p><strong>Hire Date:</strong> {selectedEmployee.hire_date || 'N/A'}</p>
                <p><strong>Termination Date:</strong> {selectedEmployee.termination_date || 'N/A'}</p>
              </div>
            )}

            {detailsActiveTab === 'status' && (
              <div>
                <p><strong>Employee Type:</strong> {selectedEmployee.status_type}</p>
                
                {selectedEmployee.status_type === 'Hourly' ? (
                  <div style={{ marginTop: '10px' }}>
                    <p><strong>Hourly Pay:</strong> ${selectedEmployee.hourly_pay}</p>
                    <p><strong>Max Hours/Week:</strong> {selectedEmployee.max_hours_per_week}</p>
                  </div>
                ) : (
                  <div style={{ marginTop: '10px' }}>
                    <p><strong>Base Salary:</strong> ${selectedEmployee.base_salary}</p>
                    <p><strong>Bonus:</strong> ${selectedEmployee.bonus}</p>
                    <p><strong>Deductions:</strong> ${selectedEmployee.deductions}</p>
                    <p><strong>PTO Hours:</strong> {selectedEmployee.paid_time_off_hours}</p>
                    <p><strong>Sick Hours:</strong> {selectedEmployee.sick_hours}</p>
                    <p><strong>Effective Period:</strong> {selectedEmployee.effective_from || 'N/A'} to {selectedEmployee.effective_to || 'N/A'}</p>
                  </div>
                )}
              </div>
            )}
          </div>
          
          <div style={{ marginTop: '20px', borderTop: '1px solid #eee', paddingTop: '15px' }}>
            <button className="btn btn-secondary" onClick={() => setIsDetailsModalOpen(false)}>Close</button>
          </div>
        </Modal>
      )}
    </div>
  );
};

export default EmployeeTable;
