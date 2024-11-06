import apiClient from '../utils/apiClient';

const authService = {
  login: async (data: { email: string; password: string }) => {
    const response = await apiClient.post('/login', data);
    return response.data;
  },
  register: async (data: { username: string ; email: string; password: string }) => {
    const response = await apiClient.post('/register', data);
    return response.data;
  },
};

export default authService;
