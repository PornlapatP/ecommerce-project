'use client'; // ทำให้ไฟล์นี้เป็น Client Component
import localFont from "next/font/local";
import "./globals.css";
// import Navbar from '../components/layout/Navbar'; // Uncomment หากต้องการใช้
 
const geistSans = localFont({
  src: "./fonts/GeistVF.woff",
  variable: "--font-geist-sans",
  weight: "100 900",
});

const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff",
  variable: "--font-geist-mono",
  weight: "100 900",
});

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}>
        {children}
        {/* Uncomment หากต้องการ Navbar */}
        {/* <Navbar /> */}
        
      </body>
    </html>
  );
}
