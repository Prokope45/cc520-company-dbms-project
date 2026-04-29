// Authors:
// 	- Jared Paubel
//  - RooCode agent - local qwen/qwen3.5-9b
// Percentage written by Agent: 100%

// API Base URL - Frontend server
const API_URL = 'http://localhost:3000';

// Initialize DataTables
let companiesTable, departmentsTable;

$(document).ready(function() {
    // Initialize DataTables with column definitions
    companiesTable = $('#companiesTable').DataTable({
        columns: [
            { data: 'company_id', title: 'ID' },
            { data: 'name', title: 'Name' },
            { data: 'created_date', title: 'Created Date' },
            { title: 'Actions', orderable: false, render: function(data, type, row) {
                return '<a href="#" class="btn btn-sm btn-edit" data-action="edit">Edit</a>' +
                    '<a href="#" class="btn btn-sm btn-delete" data-action="delete">Delete</a>';
            }}
        ],
        order: [[0, 'asc']]
    });
    
    departmentsTable = $('#departmentsTable').DataTable({
        columns: [
            { data: 'department_id', title: 'ID' },
            { data: 'company_id', title: 'Company ID' },
            { data: 'name', title: 'Name' },
            { data: 'description', title: 'Description' },
            { title: 'Actions', orderable: false, render: function(data, type, row) {
                return '<a href="#" class="btn btn-sm btn-edit" data-action="edit">Edit</a>' +
                    '<a href="#" class="btn btn-sm btn-delete" data-action="delete">Delete</a>';
            }}
        ],
        order: [[0, 'asc']]
    });
    
    // Load data on page load
    loadCompanies();
    loadDepartments();
    
    // Add Company button
    $('#addCompanyBtn').click(function() {
        openModal('Add Company', 'company');
    });
    
    // Add Department button
    $('#addDepartmentBtn').click(function() {
        openModal('Add Department', 'department');
    });
    
    // Form submit handler
    $('#crudForm').submit(function(e) {
        e.preventDefault();
        
        const recordId = $('#recordId').val();
        const type = $('#crudForm').data('type');
        
        if (recordId) {
            // Update existing record
            updateRecord(type, recordId);
        } else {
            // Create new record
            createRecord(type);
        }
    });
    
    // Cancel button
    $('#cancelBtn').click(function() {
        closeModal();
    });
});

// Load all companies
function loadCompanies() {
    $.get(`${API_URL}/companies`, function(data) {
        companiesTable.clear();
        companiesTable.rows.add(data).draw();
    }).fail(function(err) {
        alert('Error loading companies: ' + err.responseText);
    });
}

// Load all departments
function loadDepartments() {
    $.get(`${API_URL}/departments`, function(data) {
        departmentsTable.clear();
        departmentsTable.rows.add(data).draw();
    }).fail(function(err) {
        alert('Error loading departments: ' + err.responseText);
    });
}

// Create company
function createRecord(type) {
    const data = {
        name: $('#companyName').val() || $('#deptName').val()
    };
    
    if (type === 'department') {
        data.description = $('#deptDescription').val();
    }
    
    $.post(`${API_URL}/${type}s`, data)
        .done(function() {
            alert('Record created successfully');
            closeModal();
            loadCompanies();
            loadDepartments();
        })
        .fail(function(err) {
            alert('Error creating record: ' + err.responseText);
        });
}

// Update record
function updateRecord(type, id) {
    const data = {
        name: $('#companyName').val() || $('#deptName').val()
    };
    
    if (type === 'department') {
        data.description = $('#deptDescription').val();
    }
    
    $.put(`${API_URL}/${type}s/${id}`, data)
        .done(function() {
            alert('Record updated successfully');
            closeModal();
            loadCompanies();
            loadDepartments();
        })
        .fail(function(err) {
            alert('Error updating record: ' + err.responseText);
        });
}

// Delete record
function deleteRecord(type, id) {
    if (!confirm(`Are you sure you want to delete this ${type}?`)) {
        return;
    }
    
    $.ajax({
        url: `${API_URL}/${type}s/${id}`,
        method: 'DELETE',
        success: function() {
            alert('Record deleted successfully');
            loadCompanies();
            loadDepartments();
        },
        error: function(err) {
            alert('Error deleting record: ' + err.responseText);
        }
    });
}

// Open modal for add/edit
function openModal(title, type) {
    $('#modalTitle').text(title);
    $('#crudForm').data('type', type);
    $('#recordId').val('');
    $('#companyName').val('');
    $('#deptName').val('');
    $('#deptDescription').val('');
    modal.style.display = "block";
}

// Close modal
function closeModal() {
    modal.style.display = "none";
    $('#crudForm').trigger('reset');
}

// Edit button click handler (added to DataTables)
function addEditButtons(table) {
    table.column(3).every(function() {
        $(this).node().innerHTML = '<a href="#" class="btn btn-sm btn-edit" data-action="edit">Edit</a>' +
            '<a href="#" class="btn btn-sm btn-delete" data-action="delete">Delete</a>';
    });
}

// Handle edit action
$('#companiesTable tbody, #departmentsTable tbody').on('click', '.btn-edit', function(e) {
    e.preventDefault();
    const table = $(this).closest('table');
    const row = table.row($(this)).data();
    const id = row.company_id;
    const type = table[0].config.columns[0].data;
    
    $('#recordId').val(id);
    if (type === 'company') {
        $('#companyName').val(row.name);
    } else {
        $('#deptName').val(row.name);
        $('#deptDescription').val(row.description);
    }
    openModal('Edit Record', type);
});

// Handle delete action
$('#companiesTable tbody, #departmentsTable tbody').on('click', '.btn-delete', function(e) {
    e.preventDefault();
    const table = $(this).closest('table');
    const row = table.row($(this)).data();
    const id = row.company_id;
    const type = table[0].config.columns[0].data;
    
    deleteRecord(type, id);
});
