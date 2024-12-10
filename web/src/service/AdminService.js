import axios from 'axios';

export const AdminService = {
    async getStatisticsData() {
        try {
            const response = await axios.get(`/api/v1/admin/dashboard-statistic`);
            return response.data;
        } catch (error) {
            console.error('There was a problem with the axios operation:', error);
            throw error;
        }
    },

    async getLastOrders({ limit }) {
        try {
            const response = await axios.get(`/api/v1/admin/last-orders`, {
                params: { limit }
            });
            return response.data;
        } catch (error) {
            console.error('There was a problem with the axios operation:', error);
            throw error;
        }
    }
};
