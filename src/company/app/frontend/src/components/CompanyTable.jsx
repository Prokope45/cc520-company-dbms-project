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

const CompanyTable = () => {
    const [companies, setCompanies] = useState([]);
    const [isModalOpen, setIsModalOpen] = useState(false);
    const [currentCompany, setCurrentCompany] = useState({ name: '' });
    const [isEditing, setIsEditing] = useState(false);
    const [filterText, setFilterText] = useState('');

    useEffect(() => {
        fetchCompanies();
    }, []);

    const fetchCompanies = async () => {
        try {
            const response = await axios.get(`${API_BASE_URL}/companies`);
            setCompanies(response.data || []);
        } catch (error) {
            console.error('Error fetching companies:', error);
        }
    };

    const handleOpenModal = (company = { name: '' }) => {
        setCurrentCompany(company);
        setIsEditing(!!company.company_id);
        setIsModalOpen(true);
    };

    const handleCloseModal = () => {
        setIsModalOpen(false);
        setCurrentCompany({ name: '' });
        setIsEditing(false);
    };

    const handleSave = async () => {
        try {
            if (isEditing) {
                await axios.put(`${API_BASE_URL}/companies/${currentCompany.company_id}`, currentCompany);
            } else {
                await axios.post(`${API_BASE_URL}/companies`, currentCompany);
            }
            fetchCompanies();
            handleCloseModal();
        } catch (error) {
            console.error('Error saving company:', error);
        }
    };

    const handleDelete = async (id) => {
        if (window.confirm('Are you sure you want to delete this company?')) {
            try {
                await axios.delete(`${API_BASE_URL}/companies/${id}`);
                fetchCompanies();
            } catch (error) {
                console.error('Error deleting company:', error);
            }
        }
    };

    const filteredItems = companies.filter(
        item => item.name && item.name.toLowerCase().includes(filterText.toLowerCase())
    );

    const columns = [
        {
            name: 'ID',
            selector: row => row.company_id,
            sortable: true,
            width: '100px',
        },
        {
            name: 'Name',
            selector: row => row.name,
            sortable: true,
        },
        {
            name: 'Created Date',
            selector: row => new Date(row.created_date).toLocaleDateString(),
            sortable: true,
        },
        {
            name: 'Actions',
            cell: row => (
                <div style={{ display: 'flex', gap: '8px' }}>
                    <button className="btn btn-sm btn-outline-primary" onClick={() => handleOpenModal(row)}><FaEdit /></button>
                    <button className="btn btn-sm btn-outline-danger" onClick={() => handleDelete(row.company_id)}><FaTrashCan /></button>
                </div>
            ),
            ignoreRowClick: true,
        },
    ];

    return (
        <div>
            <div className="table-header-container mb-3">
                <button className="btn btn-primary" onClick={() => handleOpenModal()}>Add Company</button>
                <SearchBar 
                    filterText={filterText} 
                    onFilter={e => setFilterText(e.target.value)} 
                    placeholderText="Search companies..."
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

            <Modal isOpen={isModalOpen} onClose={handleCloseModal} title={isEditing ? 'Edit Company' : 'Add Company'}>
                <div className="form-group">
                    <label>Name:</label>
                    <input
                        type="text"
                        value={currentCompany.name || ''}
                        onChange={(e) => setCurrentCompany({ ...currentCompany, name: e.target.value })}
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

export default CompanyTable;
