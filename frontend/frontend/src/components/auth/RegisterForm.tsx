import { useState, ChangeEvent, FormEvent } from 'react';
import { useRouter } from 'next/router';
import { RegisterRequest } from '../../types/auth';
import authService from '../../services/authService';
import styles from '../../style/Register.module.css'; // Import the CSS module

const RegisterForm = () => {
  const [formData, setFormData] = useState<RegisterRequest>({ username: '', email: '', password: '' });
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false); // State for loading
  const router = useRouter();

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setError(null);
    setIsLoading(true); // Set loading state
    try {
      await authService.register(formData);
      router.push('/login'); // Redirect to login page after successful registration
    } catch (err: any) {
      setError(err.response?.data?.error || 'Registration failed');
      console.error('Registration error:', err); // Log error for debugging
    } finally {
      setIsLoading(false); // Clear loading state
    }
  };

  // Simple email validation
  const isEmailValid = formData.email.match(
    /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  );

  // Simple password validation (at least 6 characters)
  const isPasswordValid = formData.password.length >= 6;

  return (
    <div className={styles.container}>
      <h2 className={styles.heading}>Register</h2>
      {error && <p className={styles.error}>{error}</p>}
      <form className={styles.form} onSubmit={handleSubmit}>
        <input
          className={`${styles.input} ${formData.username && styles.validInput}`} // Added validInput conditionally
          type="text"
          name="username"
          placeholder="Username"
          value={formData.username}
          onChange={handleChange}
          required
        />
        <input
          className={`${styles.input} ${!isEmailValid && formData.email && styles.invalidInput}`} // Added invalidInput conditionally
          type="email"
          name="email"
          placeholder="Email"
          value={formData.email}
          onChange={handleChange}
          required
        />
        <input
          className={`${styles.input} ${!isPasswordValid && formData.password && styles.invalidInput}`} // Added invalidInput conditionally
          type="password"
          name="password"
          placeholder="Password"
          value={formData.password}
          onChange={handleChange}
          required
        />
        <button
          className={styles.button}
          type="submit"
          disabled={isLoading || !isEmailValid || !isPasswordValid}
        >
          {isLoading ? 'Registering...' : 'Register'}
        </button>
      </form>
    </div>
  );
};

export default RegisterForm;
