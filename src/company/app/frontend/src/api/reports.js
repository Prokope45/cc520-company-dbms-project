import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080';

export const reportsApi = {
    getDepartmentSalaryRanks: async (hireDate) => {
        const response = await axios.get(`${API_BASE_URL}/reports/department-salary-ranks`, {
            params: { hireDate }
        });
        return response.data;
    },

    GetDepartmentSalary_Aggregated: async (hireDate) => {
        const response = await axios.get(`${API_BASE_URL}/reports/department-salary-ranks-aggregated`, {
            params: { hireDate }
        });
        return response.data;
    },

    getTopTerminatedHourly: async (terminationDate) => {
        const response = await axios.get(`${API_BASE_URL}/reports/top-terminated-hourly`, {
            params: { terminationDate }
        });
        return response.data;
    },

    getTopTerminatedHourly_Aggregated: async (terminationDate) => {
        const response = await axios.get(`${API_BASE_URL}/reports/top-terminated-hourly-aggregated`, {
            params: { terminationDate }
        });
        return response.data;
    },

    getUnhiredWithManager: async () => {
        const response = await axios.get(`${API_BASE_URL}/reports/unhired-with-manager`);
        return response.data;
    },

    getUnhiredWithManager_Aggregated: async () => {
        const response = await axios.get(`${API_BASE_URL}/reports/unhired-with-manager-aggregated`);
        return response.data;
    },

    getHighestPaidCEO: async () => {
        const response = await axios.get(`${API_BASE_URL}/reports/highest-paid-ceo`);
        return response.data;
    },

    getHighestPaidCEO_Aggregated: async () => {
        const response = await axios.get(`${API_BASE_URL}/reports/highest-paid-ceo-aggregated`);
        return response.data;
    }
};