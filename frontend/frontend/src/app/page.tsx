import Link from 'next/link';
import styles from '../style/Home.module.css'; // Import the CSS module

const Home = () => {
  return (
    <div className={styles.container}>
      <h1 className='h1'>Welcome to the E-commerce App</h1>
      <div>
        <Link href="/register" className={styles.link}>Register</Link>
        | 
        <Link href="/login" className={styles.link}>Login</Link>
      </div>
    </div>
  );
};

export default Home;
