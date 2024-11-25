import React, { useState } from 'react';
import { useRouter } from 'next/router';
import authService from '../../services/authService';
import styles from '../../style/Login.module.css'; // Import the CSS module

const LoginForm: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState(''); // State for error message
  const [isLoading, setIsLoading] = useState(false); // State for loading
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // Validate input fields
    if (!email || !password) {
      setError('Please fill in all fields.');
      return;
    }

    // Validate email format
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      setError('Invalid email format.');
      return;
    }

    setIsLoading(true); // Set loading state
    setError(''); // Clear error

    try {
      await authService.login({ email, password });
      router.push('/home'); // Redirect to home page
    } catch (error: any) {
      if (error.response) {
        setError(error.response.data.message || 'Login failed.');
        console.error('Axios error:', error.response?.data || error.message);
      } else {
        setError('An unknown error occurred.');
        console.error('Unknown error:', error);
      }
    } finally {
      setIsLoading(false); // Clear loading state
    }
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.heading}>Login</h2>
      {error && <p className={styles.error}>{error}</p>} {/* Display error message */}
      <form className={styles.form} onSubmit={handleSubmit}>
        <input
          className={styles.input}
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
          required
        />
        <input
          className={styles.input}
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Password"
          required
        />
        <button className={styles.button} type="submit" disabled={isLoading}>
          {isLoading ? 'Logging in...' : 'Login'}
        </button>
      </form>
    </div>
  );
};

export default LoginForm;
