import React, { useState, useEffect } from 'react';
import axios from 'axios';
import DataTable from 'react-data-table-component';
import Select from 'react-select';
import Modal from './Modal';
import SearchBar from './SearchBar';
import { FaEdit } from "react-icons/fa";
import { FaTrashCan } from "react-icons/fa6";

const API_BASE_URL = 'http://localhost:8080';

// Check if DataTable is the default export or named export
const Table = DataTable.default ? DataTable.default : DataTable;

const EmployeeTable = () => {
    const [employees, setEmployees] = useState([]);
    const [departments, setDepartments] = useState([]);
    const [companies, setCompanies] = useState([]);
    const [roles, setRoles] = useState([
        'Intern', 'Developer I', 'Developer II', 'Developer III',
        'Senior Engineer', 'Member', 'Manager', 'Director',
        'HR Manager', 'CFO', 'CTO', 'CHRO', 'CEO'
    ]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [filterText, setFilterText] = useState('');

    const emptyEmployee = {
        company_id: '', department_id: '', first_name: '', last_name: '',
        email: '', phone: '', street: '', address_line_two: '', city: '', state: '', zip_code: '',
        role_title: '', manager_id: '', hire_date: '', termination_date: '',
        status_type: 'Salary', hourly_pay: 15, max_hours_per_week: 40,
        base_salary: 0, bonus: 0, deductions: 0,
        paid_time_off_hours: 0, sick_hours: 0,
        effective_from: new Date().toISOString().split('T')[0],
        effective_to: new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString().split('T')[0]
    };

    const [currentEmployee, setCurrentEmployee] = useState(emptyEmployee);
    const [isEditing, setIsEditing] = useState(false);
    const [isReadOnly, setIsReadOnly] = useState(false);
    const [activeModalTab, setActiveModalTab] = useState('person');

    useEffect(() => {
        fetchEmployees();
        fetchDepartments();
        fetchCompanies();
    }, []);

    const fetchEmployees = async () => {
        try {
            const response = await axios.get(`${API_BASE_URL}/employees`);
            setEmployees(response.data || []);
        } catch (error) {
            console.error('Error fetching employees:', error);
        }
    };

    const fetchDepartments = async () => {
        try {
            const response = await axios.get(`${API_BASE_URL}/departments`);
            setDepartments(response.data || []);
        } catch (error) {
            console.error('Error fetching departments:', error);
        }
    };

    const fetchCompanies = async () => {
        try {
            const response = await axios.get(`${API_BASE_URL}/companies`);
            setCompanies(response.data || []);
        } catch (error) {
            console.error('Error fetching companies:', error);
        }
    };

    const handleOpenModal = (employee = null, readOnly = false) => {
        if (employee) {
            // Convert to format suitable for form inputs
            setCurrentEmployee({
                ...employee,
                manager_id: employee.manager_id || '',
                hire_date: employee.hire_date ? employee.hire_date.split('T')[0] : '',
                termination_date: employee.termination_date ? employee.termination_date.split('T')[0] : '',
                effective_from: employee.effective_from ? employee.effective_from.split('T')[0] : '',
                effective_to: employee.effective_to ? employee.effective_to.split('T')[0] : ''
            });
            setIsEditing(true);
        } else {
            setCurrentEmployee(emptyEmployee);
            setIsEditing(false);
        }
        setIsReadOnly(readOnly);
        setActiveModalTab('person');
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        setCurrentEmployee(emptyEmployee);
        setIsEditing(false);
        setIsReadOnly(false);
    };

    const handleSave = async () => {
        try {
            // Format numeric values
            const empToSave = {
                ...currentEmployee,
                company_id: parseInt(currentEmployee.company_id) || 0,
                department_id: parseInt(currentEmployee.department_id) || 0,
                manager_id: currentEmployee.manager_id ? parseInt(currentEmployee.manager_id) : null,
                hourly_pay: parseFloat(currentEmployee.hourly_pay) || 0,
                max_hours_per_week: parseInt(currentEmployee.max_hours_per_week) || 0,
                base_salary: parseFloat(currentEmployee.base_salary) || 0,
                bonus: parseFloat(currentEmployee.bonus) || 0,
                deductions: parseFloat(currentEmployee.deductions) || 0,
                paid_time_off_hours: parseInt(currentEmployee.paid_time_off_hours) || 0,
                sick_hours: parseInt(currentEmployee.sick_hours) || 0,
            };

            if (isEditing) {
                await axios.put(`${API_BASE_URL}/employees/${currentEmployee.employee_id}`, empToSave);
            } else {
                await axios.post(`${API_BASE_URL}/employees`, empToSave);
            }
            fetchEmployees();
            handleCloseModal();
        } catch (error) {
            console.error('Error saving employee:', error);
            alert(`Failed to save employee: ${error.response?.data || error.message}`);
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Are you sure you want to delete this employee? This will cascade delete associated pay and role records.')) {
            try {
                await axios.delete(`${API_BASE_URL}/employees/${id}`);
                fetchEmployees();
            } catch (error) {
                console.error('Error deleting employee:', error);
                alert('Failed to delete employee.');
            }
        }
    };

    const handleInputChange = (e) => {
        const { name, value, type, checked } = e.target;
        setCurrentEmployee(prev => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : value
        }));
    };

    const filteredItems = employees.filter(
        item => {
            const searchStr = `${item.first_name} ${item.last_name} ${item.email} ${item.department} ${item.role_title}`.toLowerCase();
            return searchStr.includes(filterText.toLowerCase());
        }
    );

    const columns = [
        {
            name: 'ID',
            selector: row => row.employee_id,
            sortable: true,
            width: '80px',
        },
        {
            name: 'Name',
            selector: row => `${row.first_name} ${row.last_name}`,
            sortable: true,
            width: '180px',
        },
        {
            name: 'Email',
            selector: row => row.email,
            sortable: true,
            width: '200px',
        },
        {
            name: 'Department',
            selector: row => row.department,
            sortable: true,
        },
        {
            name: 'Role',
            selector: row => row.role_title,
            sortable: true,
        },
        {
            name: 'Type',
            selector: row => row.status_type,
            sortable: true,
            width: '100px',
        },
        {
            name: 'Manager',
            selector: row => row.manager_name,
            sortable: true,
        },
        {
            name: 'Actions',
            cell: row => (
                <div style={{ display: 'flex', gap: '8px' }}>
                    <button className="btn btn-sm btn-outline-primary" onClick={(e) => { e.stopPropagation(); handleOpenModal(row, false); }}><FaEdit /></button>
                    <button className="btn btn-sm btn-outline-danger" onClick={(e) => { e.stopPropagation(); handleDelete(row.employee_id); }}><FaTrashCan /></button>
                </div>
            ),
            ignoreRowClick: true,
        },
    ];

    // Get departments for the selected company
    const availableDepartments = departments.filter(
        d => currentEmployee.company_id ? d.company_id === parseInt(currentEmployee.company_id) : true
    );

    const handleSelectChange = (selectedOption, actionMeta) => {
        setCurrentEmployee(prev => ({
            ...prev,
            [actionMeta.name]: selectedOption ? selectedOption.value : ''
        }));
    };

    return (
        <div>
            <div className="table-header-container mb-3">
                <button className="btn btn-primary" onClick={() => handleOpenModal()}>Add Employee</button>
                <SearchBar 
                    filterText={filterText} 
                    onFilter={e => setFilterText(e.target.value)} 
                    placeholderText="Search employees..."
                />
            </div>

            <Table
                columns={columns}
                data={filteredItems}
                pagination
                onRowClicked={(row) => handleOpenModal(row, true)}
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

            <Modal isOpen={isModalOpen} onClose={handleCloseModal} title={isReadOnly ? 'View Employee Profile' : (isEditing ? 'Edit Employee Profile' : 'Add Employee Profile')} wide={true}>
                <div className="tabs-container mb-3">
                    <button className={`tab-button ${activeModalTab === 'person' ? 'active' : ''}`} onClick={() => setActiveModalTab('person')}>Person & Address</button>
                    <button className={`tab-button ${activeModalTab === 'employment' ? 'active' : ''}`} onClick={() => setActiveModalTab('employment')}>Employment Data</button>
                    <button className={`tab-button ${activeModalTab === 'pay' ? 'active' : ''}`} onClick={() => setActiveModalTab('pay')}>Pay Structure</button>
                </div>
                
                <div className="tab-content" style={{ minHeight: '400px' }}>
                    {activeModalTab === 'person' && (
                        <div>
                            <h3>Personal Info</h3>
                            <div className="form-group row">
                                <div className="col">
                                    <label>First Name:</label>
                                    <input className="form-control" type="text" name="first_name" value={currentEmployee.first_name || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                                </div>
                                <div className="col">
                                    <label>Last Name:</label>
                                    <input className="form-control" type="text" name="last_name" value={currentEmployee.last_name || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                                </div>
                            </div>
                            <div className="form-group row">
                                <div className="col">
                                    <label>Email:</label>
                                    <input className="form-control" type="email" name="email" value={currentEmployee.email || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                                </div>
                                <div className="col">
                                    <label>Phone:</label>
                                    <input className="form-control" type="text" name="phone" value={currentEmployee.phone || ''} onChange={handleInputChange}  disabled={isReadOnly} />
                                </div>
                            </div>
                            
                            <h3 className="mt-4">Address</h3>
                            <div className="form-group">
                                <label>Address Line 1:</label>
                                <input className="form-control" type="text" name="street" value={currentEmployee.street || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                            </div>
                            <div className="form-group mt-2">
                                <label>Address Line 2:</label>
                                <input className="form-control" type="text" name="address_line_two" placeholder="Apt. number, gate number, etc." value={currentEmployee.address_line_two || ''} onChange={handleInputChange}  disabled={isReadOnly} />
                            </div>
                            <div className="form-group row mt-2">
                                <div className="col">
                                    <label>City:</label>
                                    <input className="form-control" type="text" name="city" value={currentEmployee.city || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                                </div>
                                <div className="col" style={{flex: 0.5}}>
                                    <label>State:</label>
                                    <input className="form-control" type="text" name="state" value={currentEmployee.state || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                                </div>
                                <div className="col" style={{flex: 0.5}}>
                                    <label>Zip:</label>
                                    <input className="form-control" type="text" name="zip_code" value={currentEmployee.zip_code || ''} onChange={handleInputChange} required  disabled={isReadOnly} />
                                </div>
                            </div>
                        </div>
                    )}

                    {activeModalTab === 'employment' && (
                        <div>
                            <h3>Organization</h3>
                            <div className="form-group row mb-2">
                                <div className="col">
                                    <label>Company:</label>
                                    <Select 
                                        name="company_id" 
                                        value={companies.map(c => ({ value: c.company_id, label: c.name })).find(o => o.value == currentEmployee.company_id) || null}
                                        onChange={handleSelectChange}
                                        options={companies.map(c => ({ value: c.company_id, label: c.name }))}
                                        placeholder="Select Company"
                                        isClearable
                                        required
                                        isDisabled={isReadOnly}
                                    />
                                </div>
                                <div className="col">
                                    <label>Department:</label>
                                    <Select 
                                        name="department_id" 
                                        value={availableDepartments.map(d => ({ value: d.department_id, label: d.name })).find(o => o.value == currentEmployee.department_id) || null}
                                        onChange={handleSelectChange}
                                        options={availableDepartments.map(d => ({ value: d.department_id, label: d.name }))}
                                        placeholder="Select Department"
                                        isClearable
                                        required
                                        isDisabled={isReadOnly}
                                    />
                                </div>
                            </div>
                            
                            <div className="form-group row mb-2">
                                <div className="col">
                                    <label>Role:</label>
                                    <Select 
                                        name="role_title" 
                                        value={roles.map(r => ({ value: r, label: r })).find(o => o.value === currentEmployee.role_title) || null}
                                        onChange={handleSelectChange}
                                        options={roles.map(r => ({ value: r, label: r }))}
                                        placeholder="Select Role"
                                        isClearable
                                        required
                                        isDisabled={isReadOnly}
                                    />
                                </div>
                                <div className="col">
                                    <label>Manager:</label>
                                    <Select 
                                        name="manager_id" 
                                        value={employees.filter(e => e.employee_id !== currentEmployee.employee_id).map(e => ({ value: e.employee_id, label: `${e.first_name} ${e.last_name}` })).find(o => o.value == currentEmployee.manager_id) || null}
                                        onChange={handleSelectChange}
                                        options={employees.filter(e => e.employee_id !== currentEmployee.employee_id).map(e => ({ value: e.employee_id, label: `${e.first_name} ${e.last_name}` }))}
                                        placeholder="None"
                                        isClearable
                                        isDisabled={isReadOnly}
                                    />
                                </div>
                            </div>

                            <h3 className="mt-4">Dates</h3>
                            <div className="form-group row">
                                <div className="col">
                                    <label>Hire Date:</label>
                                    <input className="form-control" type="date" name="hire_date" value={currentEmployee.hire_date || ''} onChange={handleInputChange}  disabled={isReadOnly} />
                                </div>
                                <div className="col">
                                    <label>Termination Date:</label>
                                    <input className="form-control" type="date" name="termination_date" value={currentEmployee.termination_date || ''} onChange={handleInputChange}  disabled={isReadOnly} />
                                </div>
                            </div>
                        </div>
                    )}

                    {activeModalTab === 'pay' && (
                        <div>
                            <h3>Pay Structure</h3>
                            <div className="form-group">
                                <label>Status Type:</label>
                                <select className="form-select" name="status_type" value={currentEmployee.status_type || 'Salary'} onChange={handleInputChange} disabled={isReadOnly}>
                                    <option value="Salary">Salary</option>
                                    <option value="Hourly">Hourly</option>
                                </select>
                            </div>

                            {currentEmployee.status_type === 'Hourly' ? (
                                <div className="form-group row">
                                    <div className="col">
                                        <label>Hourly Pay ($):</label>
                                        <input className="form-control" type="number" step="0.01" name="hourly_pay" value={currentEmployee.hourly_pay || 0} onChange={handleInputChange}  disabled={isReadOnly} />
                                    </div>
                                    <div className="col">
                                        <label>Max Hours/Week:</label>
                                        <input className="form-control" type="number" name="max_hours_per_week" value={currentEmployee.max_hours_per_week || 40} onChange={handleInputChange}  disabled={isReadOnly} />
                                    </div>
                                </div>
                            ) : (
                                <>
                                    <div className="form-group row">
                                        <div className="col">
                                            <label>Base Salary ($):</label>
                                            <input className="form-control" type="number" name="base_salary" value={currentEmployee.base_salary || 0} onChange={handleInputChange}  disabled={isReadOnly} />
                                        </div>
                                        <div className="col">
                                            <label>Bonus ($):</label>
                                            <input className="form-control" type="number" name="bonus" value={currentEmployee.bonus || 0} onChange={handleInputChange}  disabled={isReadOnly} />
                                        </div>
                                    </div>
                                    <div className="form-group row">
                                        <div className="col">
                                            <label>PTO Hours:</label>
                                            <input className="form-control" type="number" name="paid_time_off_hours" value={currentEmployee.paid_time_off_hours || 0} onChange={handleInputChange}  disabled={isReadOnly} />
                                        </div>
                                        <div className="col">
                                            <label>Sick Hours:</label>
                                            <input className="form-control" type="number" name="sick_hours" value={currentEmployee.sick_hours || 0} onChange={handleInputChange}  disabled={isReadOnly} />
                                        </div>
                                    </div>
                                </>
                            )}
                        </div>
                    )}
                </div>
                
                <div className="modal-actions mt-4 d-flex gap-2" style={{ borderTop: '1px solid #ddd', paddingTop: '15px' }}>
                    {isReadOnly ? (
                        <>
                            <button className="btn btn-primary" onClick={() => setIsReadOnly(false)}>Edit</button>
                            <button className="btn btn-secondary" onClick={handleCloseModal}>Close</button>
                        </>
                    ) : (
                        <>
                            <button className="btn btn-success" onClick={handleSave}>Save Profile</button>
                            <button className="btn btn-secondary" onClick={handleCloseModal}>Cancel</button>
                        </>
                    )}
                </div>
            </Modal>
        </div>
    );
};

export default EmployeeTable;
