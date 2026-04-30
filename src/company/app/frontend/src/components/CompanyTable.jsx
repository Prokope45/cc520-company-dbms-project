import React, { useState, useEffect } from 'react';
import { companyAPI } from '../api';
import Modal from './Modal';

const CompanyTable = () => {
  const [companies, setCompanies] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingCompany, setEditingCompany] = useState(null);
  const [formData, setFormData] = useState({ name: '' });
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    fetchCompanies();
  }, []);

  const fetchCompanies = async () => {
    try {
      const response = await companyAPI.getAll();
      setCompanies(response.data);
      setErrorMessage('');
    } catch (error) {
      console.error('Error fetching companies:', error);
      setErrorMessage(error.response?.data?.error || 'Failed to fetch companies');
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this company?')) {
      try {
        await companyAPI.delete(id);
        fetchCompanies();
        setErrorMessage('');
      } catch (error) {
        console.error('Error deleting company:', error);
        setErrorMessage(error.response?.data?.error || 'Failed to delete company');
      }
    }
  };

  const handleEdit = (company) => {
    setEditingCompany(company);
    setFormData({ id: company.company_id, name: company.name });
    setIsModalOpen(true);
    setErrorMessage('');
  };

  const handleAdd = () => {
    setEditingCompany(null);
    setFormData({ name: '' });
    setIsModalOpen(true);
    setErrorMessage('');
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setErrorMessage('');
    try {
      if (editingCompany) {
        await companyAPI.update(editingCompany.company_id, formData);
      } else {
        await companyAPI.create(formData);
      }
      setIsModalOpen(false);
      fetchCompanies();
    } catch (error) {
      console.error('Error saving company:', error);
      const errorMessage = error.response?.data?.error || error.response?.data?.message || 'Failed to save company';
      setErrorMessage(errorMessage);
    }
  };

  return (
    <div className="section">
      <h2>Companies</h2>
      <button className="btn btn-primary" onClick={handleAdd}>+ Add Company</button>
      
      {errorMessage && (
        <div className="alert alert-danger" style={{ marginTop: '10px', padding: '10px', backgroundColor: '#f8d7da', border: '1px solid #f5c6cb', borderRadius: '4px' }}>
          {errorMessage}
        </div>
      )}
      
      <table className="display" style={{ width: '100%', marginTop: '15px' }}>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Created Date</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {companies.map(company => (
            <tr key={company.company_id}>
              <td className="table-cell-id">{company.company_id}</td>
              <td className="table-cell-text">{company.name}</td>
              <td className="table-cell-text">{new Date(company.createdDate).toLocaleString()}</td>
              <td className="table-cell-actions">
                <button className="btn btn-primary" onClick={() => handleEdit(company)} style={{ marginRight: '5px' }}>Edit</button>
                <button className="btn btn-secondary" onClick={() => handleDelete(company.company_id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <Modal
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        title={editingCompany ? "Edit Company" : "Add Company"}
      >
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="companyName">Company Name:</label>
            <input
              type="text"
              id="companyName"
              name="name"
              required
              value={formData.name}
              onChange={(e) => setFormData({ name: e.target.value })}
            />
          </div>
          <button type="submit" className="btn btn-primary">Save</button>
          <button type="button" className="btn btn-secondary" onClick={() => setIsModalOpen(false)} style={{ marginLeft: '10px' }}>Cancel</button>
        </form>
      </Modal>
    </div>
  );
};

export default CompanyTable;
