import React, { useState, useEffect } from 'react';
import axios from 'axios';
import DataTable from 'react-data-table-component';
import Modal from './Modal';
import SearchBar from './SearchBar';
import { FaEdit } from "react-icons/fa";
import { FaTrashCan } from "react-icons/fa6";

const API_BASE_URL = 'http://localhost:8080';

// Check if DataTable is the default export or named export
const Table = DataTable.default ? DataTable.default : DataTable;

const DepartmentTable = () => {
    const [departments, setDepartments] = useState([]);
    const [companies, setCompanies] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [currentDepartment, setCurrentDepartment] = useState({ company_id: '', name: '', description: '' });
    const [isEditing, setIsEditing] = useState(false);
    const [filterText, setFilterText] = useState('');

    useEffect(() => {
        fetchDepartments();
        fetchCompanies();
    }, []);

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

    const handleOpenModal = (department = { company_id: '', name: '', description: '' }) => {
        setCurrentDepartment(department);
        setIsEditing(!!department.department_id);
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        setCurrentDepartment({ company_id: '', name: '', description: '' });
        setIsEditing(false);
    };

    const handleSave = async () => {
        try {
            // Ensure company_id is a number
            const deptToSave = {
                ...currentDepartment,
                company_id: parseInt(currentDepartment.company_id, 10)
            };

            if (isEditing) {
                await axios.put(`${API_BASE_URL}/departments/${currentDepartment.department_id}`, deptToSave);
            } else {
                await axios.post(`${API_BASE_URL}/departments`, deptToSave);
            }
            fetchDepartments();
            handleCloseModal();
        } catch (error) {
            console.error('Error saving department:', error);
            alert('Failed to save department. Please check the console for details.');
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Are you sure you want to delete this department?')) {
            try {
                await axios.delete(`${API_BASE_URL}/departments/${id}`);
                fetchDepartments();
            } catch (error) {
                console.error('Error deleting department:', error);
            }
        }
    };

    const filteredItems = departments.filter(
        item => {
            const companyName = companies.find(c => c.company_id === item.company_id)?.name || '';
            const searchStr = `${item.name} ${item.description} ${companyName}`.toLowerCase();
            return searchStr.includes(filterText.toLowerCase());
        }
    );

    const columns = [
        {
            name: 'ID',
            selector: row => row.department_id,
            sortable: true,
            width: '80px',
        },
        {
            name: 'Company',
            selector: row => companies.find(c => c.company_id === row.company_id)?.name || 'Unknown',
            sortable: true,
            width: '200px',
        },
        {
            name: 'Name',
            selector: row => row.name,
            sortable: true,
        },
        {
            name: 'Description',
            selector: row => row.description,
            sortable: true,
        },
        {
            name: 'Actions',
            cell: row => (
                <div style={{ display: 'flex', gap: '8px' }}>
                    <button className="btn btn-sm btn-outline-primary" onClick={() => handleOpenModal(row)}><FaEdit /></button>
                    <button className="btn btn-sm btn-outline-danger" onClick={() => handleDelete(row.department_id)}><FaTrashCan /></button>
                </div>
            ),
            ignoreRowClick: true,
        },
    ];

    return (
        <div>
            <div className="table-header-container mb-3">
                <button className="btn btn-primary" onClick={() => handleOpenModal()}>Add Department</button>
                <SearchBar 
                    filterText={filterText} 
                    onFilter={e => setFilterText(e.target.value)} 
                    placeholderText="Search departments..."
                />
            </div>

            <Table
                columns={columns}
                data={filteredItems}
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

            <Modal isOpen={isModalOpen} onClose={handleCloseModal} title={isEditing ? 'Edit Department' : 'Add Department'}>
                <div className="form-group">
                    <label>Company:</label>
                    <select
                        className='form-select'
                        value={currentDepartment.company_id || ''}
                        onChange={(e) => setCurrentDepartment({ ...currentDepartment, company_id: e.target.value })}
                        required
                    >
                        <option value="">Select a Company</option>
                        {companies.map(company => (
                            <option key={company.company_id} value={company.company_id}>
                                {company.name}
                            </option>
                        ))}
                    </select>
                </div>
                <div className="form-group">
                    <label>Name:</label>
                    <input
                        type="text"
                        value={currentDepartment.name || ''}
                        onChange={(e) => setCurrentDepartment({ ...currentDepartment, name: e.target.value })}
                        required
                    />
                </div>
                <div className="form-group">
                    <label>Description:</label>
                    <textarea
                        value={currentDepartment.description || ''}
                        onChange={(e) => setCurrentDepartment({ ...currentDepartment, description: e.target.value })}
                        rows="3"
                    />
                </div>
                <div className="modal-actions mt-4 d-flex gap-2">
                    <button className="btn btn-success" onClick={handleSave}>Save</button>
                    <button className="btn btn-secondary" onClick={handleCloseModal}>Cancel</button>
                </div>
            </Modal>
        </div>
    );
};

export default DepartmentTable;
