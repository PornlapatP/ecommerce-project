// services/productService.ts
import axios from 'axios';
import { Product } from '../types/product';

interface ProductResponse {
  data: {
    count: number;
    results: Product[];
  };
}

// ฟังก์ชันดึงข้อมูลผลิตภัณฑ์ทั้งหมด
const getAllProducts = async (): Promise<Product[]> => {
  const token = localStorage.getItem('token'); // ดึง token จาก localStorage

  const response = await axios.get<ProductResponse>('http://localhost:2027/items/products', {
    headers: {
      Authorization: `Bearer ${token}`, // เพิ่ม Authorization header
    },
  });

  return response.data.data.results;
};

// ฟังก์ชันเพิ่มผลิตภัณฑ์ใหม่
const createProduct = async (product: Product): Promise<Product> => {
  const token = localStorage.getItem('token');

  const response = await axios.post<Product>('http://localhost:2025/items/products', product, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  return response.data;
};

// ฟังก์ชันอัพเดตข้อมูลผลิตภัณฑ์
const updateProduct = async (id: string, product: Product): Promise<Product> => {
  const token = localStorage.getItem('token');

  const response = await axios.put<Product>(`http://localhost:2025/items/products/${id}`, product, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  return response.data;
};

// ฟังก์ชันลบผลิตภัณฑ์
const deleteProduct = async (id: string): Promise<void> => {
  const token = localStorage.getItem('token');

  await axios.delete(`http://localhost:2025/items/products/${id}`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });
};

// ฟังก์ชันอัพเดตสถานะผลิตภัณฑ์
const updateStatusProduct = async (id: string, status: string): Promise<Product> => {
  const token = localStorage.getItem('token');

  const response = await axios.patch<Product>(
    `http://localhost:2025/items/products/${id}`,
    { status },
    {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  );

  return response.data;
};

// ฟังก์ชันอัพโหลดรูปภาพ
const uploadImage = async (formData: FormData): Promise<string> => {
  const token = localStorage.getItem('token');

  const response = await axios.post<{ image_url: string }>(
    'http://localhost:2025/items/upload-image',
    formData,
    {
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'multipart/form-data',
      },
    }
  );

  return response.data.image_url; // Returning the URL of the uploaded image
};

export default {
  getAllProducts,
  createProduct,
  updateProduct,
  deleteProduct,
  updateStatusProduct,
  uploadImage,
};
