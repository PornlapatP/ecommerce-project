// pages/products.tsx
import { useEffect, useState } from 'react';
import productService from '../services/productService';
import { Product } from '../types/product';
import ProductItem from '../components/layout/ProductCard';
import styles from '../style/ProductsPage.module.css';
import withAuth from '../utils/withAuth';

const ProductsPage = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [newProduct, setNewProduct] = useState<Partial<Product>>({
    name: '',
    description: '',
    price: 0,
    stock: 0,
    image_url: '',
    status: 'active',
  });

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const data = await productService.getAllProducts();
        setProducts(data);
      } catch (err: any) {
        setError('Error fetching products');
        console.error('Error fetching products:', err);
      }
    };

    fetchProducts();
  }, []);

  const handleCreateOrUpdateProduct = async () => {
    // ตรวจสอบข้อมูลก่อนการบันทึก
    if (!newProduct.name || !newProduct.price || !newProduct.stock) {
      setError('Name, Price, and Stock are required fields.');
      return;
    }

    try {
      if (newProduct.id) {
        // หากมี id แสดงว่าเป็นการอัปเดตสินค้า
        const updatedProduct = await productService.updateProduct(newProduct.id.toString(), newProduct as Product);
        setProducts(
          products.map((product) =>
            product.id === updatedProduct.id ? updatedProduct : product
          )
        );
      } else {
        // หากไม่มี id แสดงว่าเป็นการสร้างสินค้าใหม่
        const createdProduct = await productService.createProduct(newProduct as Product);
        setProducts([...products, createdProduct]);
      }

      setNewProduct({
        name: '',
        description: '',
        price: 0,
        stock: 0,
        image_url: '',
        status: 'active',
      });
      setError(null);
    } catch (err: any) {
      setError('Error creating or updating product');
      console.error('Error creating/updating product:', err);
    }
  };

  const handleDeleteProduct = async (id: number) => {
    try {
      await productService.deleteProduct(id.toString());
      setProducts(products.filter((product) => product.id !== id));
    } catch (err: any) {
      setError('Error deleting product');
      console.error('Error deleting product:', err);
    }
  };

  const handleStatusUpdate = async (id: number, status: string) => {
    try {
      const updatedProduct = await productService.updateStatusProduct(id.toString(), status);
      setProducts(
        products.map((product) =>
          product.id === id ? { ...product, status: updatedProduct.status } : product
        )
      );
    } catch (err: any) {
      setError('Error updating product status');
      console.error('Error updating product status:', err);
    }
  };

  const handleEditProduct = (product: Product) => {
    setNewProduct(product);
  };

  const handleImageUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const formData = new FormData();
      formData.append('image', file);

      try {
        const uploadedImageUrl = await productService.uploadImage(formData);
        setNewProduct({ ...newProduct, image_url: uploadedImageUrl });
      } catch (err) {
        console.error('Error uploading image:', err);
      }
    }
  };

  return (
    <div className={styles.container}>
      <h2 className={styles.heading}>Products</h2>
      {error && <p className={styles.error}>{error}</p>}
      {products.length === 0 ? (
        <p className={styles.noProducts}>No products available</p>
      ) : (
        <ul className={styles.productList}>
          {products.map((product) => (
            <ProductItem
              key={product.id}
              product={product}
              onDelete={(id: string) => handleDeleteProduct(Number(id))}
              onUpdateStatus={handleStatusUpdate}
              onEdit={handleEditProduct}
            />
          ))}
        </ul>
      )}

      <div className={styles.createProductForm}>
        <h3>{newProduct.id ? 'Edit Product' : 'Add New Product'}</h3>

        {/* Product Name */}
        <div>
          <label htmlFor="name">Product Name</label>
          <input
            id="name"
            type="text"
            placeholder="Enter Product Name"
            value={newProduct.name}
            onChange={(e) => setNewProduct({ ...newProduct, name: e.target.value })}
          />
        </div>

        {/* Description */}
        <div>
          <label htmlFor="description">Description</label>
          <input
            id="description"
            type="text"
            placeholder="Enter Product Description"
            value={newProduct.description}
            onChange={(e) => setNewProduct({ ...newProduct, description: e.target.value })}
          />
        </div>

        {/* Price */}
        <div>
          <label htmlFor="price">Price</label>
          <input
            id="price"
            type="number"
            placeholder="Enter Product Price"
            value={newProduct.price}
            onChange={(e) => setNewProduct({ ...newProduct, price: parseFloat(e.target.value) })}
          />
        </div>

        {/* Stock */}
        <div>
          <label htmlFor="stock">Stock</label>
          <input
            id="stock"
            type="number"
            placeholder="Enter Stock Quantity"
            value={newProduct.stock}
            onChange={(e) => setNewProduct({ ...newProduct, stock: parseInt(e.target.value) })}
          />
        </div>

        {/* Image Upload */}
        <div>
          <label htmlFor="image_url">Image</label>
          <input
            id="image_url"
            type="file"
            accept="image/*"
            onChange={handleImageUpload}
          />
          {newProduct.image_url && <p>Current image: {newProduct.image_url}</p>}
        </div>

        <button onClick={handleCreateOrUpdateProduct}>
          {newProduct.id ? 'Update Product' : 'Create Product'}
        </button>
      </div>
    </div>
  );
};

export default withAuth(ProductsPage);
