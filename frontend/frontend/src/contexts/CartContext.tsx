// // contexts/CartContext.tsx
// import React, { createContext, useState, useContext, ReactNode, useCallback } from 'react';
// import { Product } from '../types/product';

// interface CartItem extends Product {
//   quantity: number;
// }

// interface CartContextProps {
//   cart: CartItem[];
//   addToCart: (product: Product) => void;
//   removeFromCart: (productId: number) => void;
//   clearCart: () => void;
// }

// const CartContext = createContext<CartContextProps | undefined>(undefined);

// export const CartProvider = ({ children }: { children: ReactNode }) => {
//   const [cart, setCart] = useState<CartItem[]>([]);

//   const addToCart = useCallback((product: Product) => {
//     setCart((prevCart) => {
//       const existingItem = prevCart.find((item) => item.id === product.id);
//       if (existingItem) {
//         return prevCart.map((item) =>
//           item.id === product.id ? { ...item, quantity: item.quantity + 1 } : item
//         );
//       } else {
//         return [...prevCart, { ...product, quantity: 1 }];
//       }
//     });
//   }, []);

//   const removeFromCart = useCallback((productId: number) => {
//     setCart((prevCart) => prevCart.filter((item) => item.id !== productId));
//   }, []);

//   const clearCart = useCallback(() => {
//     setCart([]);
//   }, []);

//   return (
//     <CartContext.Provider value={{ cart, addToCart, removeFromCart, clearCart }}>
//       {children}
//     </CartContext.Provider>
//   );
// };

// export const useCart = () => {
//   const context = useContext(CartContext);
//   if (context === undefined) {
//     throw new Error("useCart must be used within a CartProvider");
//   }
//   return context;
// };
