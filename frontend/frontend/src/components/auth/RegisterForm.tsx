// components/RegisterForm.tsx
import { useState, ChangeEvent, FormEvent } from 'react';
// import axios from 'axios';
import { useRouter } from 'next/router';
import { RegisterRequest } from '../../types/auth'
import authService from '../../services/authService';
import styles from '../../style/Register.module.css'; // Import the CSS module
const RegisterForm = () => {
  const [formData, setFormData] = useState<RegisterRequest>({ username: '', email: '', password: '' });
  const [error, setError] = useState<string | null>(null);
  const router = useRouter();

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError(null);
    try {
      await authService.register(formData);
      router.push('/login'); // เปลี่ยนเส้นทางไปยังหน้าล็อกอินหลังจากสมัครสมาชิกสำเร็จ
    } catch (err: any) {
      setError(err.response?.data?.error || 'Registration failed');
    }
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.heading}>Register</h2>
      {error && <p className={styles.error}>{error}</p>}
      <form className={styles.form} onSubmit={handleSubmit}>
        <input
          className={styles.input}
          type="text"
          name="username"
          placeholder="Username"
          value={formData.username}
          onChange={handleChange}
          required
        />
        <input
          className={styles.input}
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          required
        />
        <input
          className={styles.input}
          type="password"
          name="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange}
          required
        />
        <button className={styles.button} type="submit">Register</button>
      </form>
    </div>
  );
};

export default RegisterForm;
