// pages/home.tsx
import Navbar from '../components/layout/Navbar';
import withAuth from '../components/layout/withAuth';

const Home = () => {
  return (
    <div>
      <Navbar />
      <h1>Welcome to the Home Page</h1>
      {/* เนื้อหาหลักของหน้า Home */}
    </div>
  );
};

export default withAuth(Home);
