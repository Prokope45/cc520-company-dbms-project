import React, { useState, useEffect } from 'react';
import { companyAPI } from '../api';
import Modal from './Modal';

const CompanyTable = () => {
  const [companies, setCompanies] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingCompany, setEditingCompany] = useState(null);
  const [formData, setFormData] = useState({ name: '' });

  useEffect(() => {
    fetchCompanies();
  }, []);

  const fetchCompanies = async () => {
    try {
      const response = await companyAPI.getAll();
      setCompanies(response.data);
    } catch (error) {
      console.error('Error fetching companies:', error);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this company?')) {
      try {
        await companyAPI.delete(id);
        fetchCompanies();
      } catch (error) {
        console.error('Error deleting company:', error);
      }
    }
  };

  const handleEdit = (company) => {
    setEditingCompany(company);
    setFormData({ name: company.name });
    setIsModalOpen(true);
  };

  const handleAdd = () => {
    setEditingCompany(null);
    setFormData({ name: '' });
    setIsModalOpen(true);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (editingCompany) {
        await companyAPI.update(editingCompany.id, formData);
      } else {
        await companyAPI.create(formData);
      }
      setIsModalOpen(false);
      fetchCompanies();
    } catch (error) {
      console.error('Error saving company:', error);
    }
  };

  return (
    <div className="section">
      <h2>Companies</h2>
      <button className="btn btn-primary" onClick={handleAdd}>+ Add Company</button>
      
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
            <tr key={company.id}>
              <td>{company.id}</td>
              <td>{company.name}</td>
              <td>{new Date(company.createdAt).toLocaleString()}</td>
              <td>
                <button className="btn btn-primary" onClick={() => handleEdit(company)} style={{ marginRight: '5px' }}>Edit</button>
                <button className="btn btn-secondary" onClick={() => handleDelete(company.id)}>Delete</button>
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
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
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
