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
  