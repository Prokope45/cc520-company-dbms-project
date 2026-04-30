import React, { useState, useEffect } from 'react';
import { departmentAPI } from '../api';
import Modal from './Modal';

const DepartmentTable = () => {
  const [departments, setDepartments] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingDepartment, setEditingDepartment] = useState(null);
  const [formData, setFormData] = useState({ company_id: '', name: '', description: '' });

  useEffect(() => {
    fetchDepartments();
  }, []);

  const fetchDepartments = async () => {
    try {
      const response = await departmentAPI.getAll();
      setDepartments(response.data);
    } catch (error) {
      console.error('Error fetching departments:', error);
    }
  };

  const handleDelete = async (id) => {
    if (window.confirm('Are you sure you want to delete this department?')) {
      try {
        await departmentAPI.delete(id);
        fetchDepartments();
      } catch (error) {
        console.error('Error deleting department:', error);
      }
    }
  };

  const handleEdit = (department) => {
    setEditingDepartment(department);
    setFormData({ 
      company_id: department.company_id || '', 
      name: department.name || '', 
      description: department.description || '' 
    });
    setIsModalOpen(true);
  };

  const handleAdd = () => {
    setEditingDepartment(null);
    setFormData({ company_id: '', name: '', description: '' });
    setIsModalOpen(true);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      // Convert company_id to number
      const payload = { ...formData, company_id: parseInt(formData.company_id, 10) };
      
      if (editingDepartment) {
        await departmentAPI.update(editingDepartment.id, payload);
      } else {
        await departmentAPI.create(payload);
      }
      setIsModalOpen(false);
      fetchDepartments();
    } catch (error) {
      console.error('Error saving department:', error);
    }
  };
  return (
    <div className="section">
      <h2>Departments</h2>
      <button className="btn btn-primary" onClick={handleAdd}>+ Add Department</button>
      
      <table className="display" style={{ width: '100%', marginTop: '15px' }}>
        <thead>
          <tr>
            <th>ID</th>
            <th>Company ID</th>
            <th>Name</th>
            <th>Description</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {departments.map(dept => (
            <tr key={dept.department_id}>
              <td className="table-cell-id">{dept.department_id}</td>
              <td className="table-cell-number">{dept.company_id}</td>
              <td className="table-cell-text">{dept.name}</td>
              <td className="table-cell-text">{dept.description}</td>
              <td className="table-cell-actions">
                <button className="btn btn-primary" onClick={() => handleEdit(dept)} style={{ marginRight: '5px' }}>Edit</button>
                <button className="btn btn-secondary" onClick={() => handleDelete(dept.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <Modal 
        isOpen={isModalOpen} 
        onClose={() => setIsModalOpen(false)}
        title={editingDepartment ? "Edit Department" : "Add Department"}
      >
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="deptCompanyId">Company ID:</label>
            <input 
              type="number" 
              id="deptCompanyId" 
              name="company_id" 
              required 
              value={formData.company_id}
              onChange={(e) => setFormData({ ...formData, company_id: e.target.value })}
            />
          </div>
          <div className="form-group">
            <label htmlFor="deptName">Department Name:</label>
            <input 
              type="text" 
              id="deptName" 
              name="name" 
              required 
              value={formData.name}
              onChange={(e) => setFormData({ ...formData, name: e.target.value })}
            />
          </div>
          <div className="form-group">
            <label htmlFor="deptDescription">Department Description:</label>
            <textarea 
              id="deptDescription" 
              name="description" 
              value={formData.description}
              onChange={(e) => setFormData({ ...formData, description: e.target.value })}
            ></textarea>
          </div>
          <button type="submit" className="btn btn-primary">Save</button>
          <button type="button" className="btn btn-secondary" onClick={() => setIsModalOpen(false)} style={{ marginLeft: '10px' }}>Cancel</button>
        </form>
      </Modal>
    </div>
  );
};

export default DepartmentTable;
