import { ref, onMounted } from 'vue';
import axios from 'axios';

export function useUserRole() {
    const userRole = ref(null);

    const fetchUserRole = async () => {
        try {
            const response = await axios.get(`/api/v1/user-info`);
            userRole.value = response.data.data.role;
        } catch (error) {
            console.error('Failed to fetch user role:', error);
        }
    };

    onMounted(() => {
        fetchUserRole();
    });

    return { userRole };
}
