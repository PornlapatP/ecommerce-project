export interface LoginRequest {
    email: string;
    password: string;
  }
  
  export interface RegisterRequest {
    username: string;
    email: string;
    password: string;
  }
  
  export interface AuthResponse {
    token: string;
    message: string;
    // userId: string;
  }

  export interface TokenType {
    exp: number;     // เวลาหมดอายุของ token (timestamp)
    id: number;      // ID ของผู้ใช้
    role: string;    // Role ของผู้ใช้ (เช่น admin)
    username: string; // ชื่อผู้ใช้
  }
  // interface LoginResponse {
  //   token: string;
  // }
  