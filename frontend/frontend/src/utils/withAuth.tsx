// utils/withAuth.tsx
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';

const withAuth = (WrappedComponent: React.FC) => {
  const AuthenticatedComponent = (props: any) => {
    const router = useRouter();
    const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);

    useEffect(() => {
      const token = localStorage.getItem('token'); // ตรวจสอบ token ใน localStorage

      if (!token) {
        router.replace('/login'); // รีไดเรกต์ไปที่หน้า login ทันที
      } else {
        setIsAuthenticated(true); // ถ้า token ถูกต้องให้ตั้งค่า isAuthenticated เป็น true
      }
    }, [router]);

    // รอให้ตรวจสอบสิทธิ์เสร็จก่อนแสดงหน้า
    if (isAuthenticated === null) {
      return null; // หรือแสดง Loading Component
    }

    return <WrappedComponent {...props} />;
  };

  return AuthenticatedComponent;
};

export default withAuth;
