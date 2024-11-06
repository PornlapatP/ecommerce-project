import apiClient from '../utils/apiClient';
import { LoginRequest } from '@/types/auth';
import { RegisterRequest } from '@/types/auth';
import { AuthResponse } from '@/types/auth';
const authService = {
  login: async (data: LoginRequest ) => {
    try{
      const response = await apiClient.post<AuthResponse>('/login', data);
      const { token } = response.data;
      localStorage.setItem('token', token); // จัดเก็บ token
      return response.data;
    }catch (error){
      throw error;
    }
  },
  register: async (data: RegisterRequest ) => {
    try{
      const response = await apiClient.post('/register', data);
      return response.data;
    }catch (error){
      throw error;
    }

  },
  logout: () => {
    localStorage.removeItem('token');
  },

  // ตรวจสอบ token ใน localStorage ว่าเข้าสู่ระบบอยู่หรือไม่
  isLoggedIn: (): boolean => {
    const token = localStorage.getItem('token');
    return !!token; // true ถ้ามี token, false ถ้าไม่มี
  },

  // ดึงข้อมูล token ปัจจุบัน
  getToken: (): string | null => {
    return localStorage.getItem('token');
  },
};

export default authService;
