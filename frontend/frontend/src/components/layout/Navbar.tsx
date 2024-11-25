// 'use client'; // คำสั่งนี้ต้องอยู่บรรทัดแรก
'use client';  // ทำให้คอมโพเนนต์นี้เป็น Client-Side

import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import Link from 'next/link';
import jwt from 'jsonwebtoken';
import authService from '../../services/authService';
import styles from '../../style/Navbar.module.css'; // Import the CSS module
import { TokenType } from '@/types/auth'

const Navbar = () => {
  const router = useRouter();
  const [username, setUsername] = useState<string | null>(null);
  const [role, setRole] = useState<string | null>(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      try {
        // Use TypeScript's type assertion to tell the compiler that decodedToken is of type TokenType
        const decodedToken = jwt.decode(token) as TokenType;

        // console.log(decodedToken);  // Check the decoded token data

        if (decodedToken && decodedToken.username && decodedToken.role) {
          setUsername(decodedToken.username);
          setRole(decodedToken.role);
        } else {
          console.error('Invalid token');
          router.push('/login');
        }
      } catch (error) {
        console.error('Invalid token', error);
        router.push('/login');
      }
    } else {
      router.push('/login');
    }
  }, [router]);

  const handleLogout = () => {
    authService.logout();
    router.push('/login');
  };

  return (
    <nav className={styles.navbar}>
      <div className={styles['nav-left']}>
        {username && role && (
          <p>
            Welcome, {username} ({role})
          </p>
        )}
      </div>
      <div className={styles['nav-right']}>
        <Link href="/products">Products</Link>
        <Link href="/orders">Orders</Link>
        <Link href="/cart">Cart</Link>
        <Link href="/profile">Profile</Link>
        <button onClick={handleLogout}>Logout</button>
      </div>
    </nav>
  );
};

export default Navbar;
