// components/layout/ProductCard.tsx
import React from 'react';
import { Product } from '../../types/product';
import styles from '../../style/ProductCard.module.css'; // Import the CSS module

interface ProductCardProps {
  product: Product;
  onDelete: (id: string) => void;
  onUpdateStatus: (id: number, status: string) => void;
  onEdit: (product: Product) => void;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onDelete, onUpdateStatus, onEdit }) => {
  const handleDelete = () => {
    onDelete(product.id.toString());
  };

  const handleStatusChange = (status: string) => {
    onUpdateStatus(product.id, status);
  };

  return (
    <div className={styles.productCard}>
      <h4 className={styles.productName}>{product.name}</h4>
      {product.image_url && (
        <img
          src={product.image_url}
          alt={product.name}
          className={styles.productImage}
        />
      )}
      <p className={styles.productDescription}>{product.description}</p>
      <p className={styles.productPrice}>Price: ${product.price}</p>
      <p className={styles.productStock}>Stock: {product.stock}</p>
      <p className={styles.productStatus}>Status: {product.status}</p>
      <div className={styles.productButtons}>
        <button onClick={handleDelete} className={styles.deleteButton}>Delete</button>
        <button onClick={() => onEdit(product)} className={styles.editButton}>Edit</button>
      </div>
      <select
        className={styles.statusDropdown}
        aria-label="Product Status"
        value={product.status}
        onChange={(e) => handleStatusChange(e.target.value)}
      >
        <option value="active">Active</option>
        <option value="inactive">Inactive</option>
      </select>
    </div>
  );
};

export default ProductCard;
