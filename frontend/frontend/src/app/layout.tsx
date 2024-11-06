import type { Metadata } from "next"; // นำเข้า Metadata สำหรับการจัดการข้อมูลเมตา
import localFont from "next/font/local"; // นำเข้า localFont เพื่อใช้งานฟอนต์ในโปรเจ็กต์
import "./globals.css"; // นำเข้าไฟล์ CSS สำหรับสไตล์ทั่วทั้งแอป

// กำหนดฟอนต์ Geist Sans โดยระบุเส้นทางและน้ำหนัก
const geistSans = localFont({
  src: "./fonts/GeistVF.woff", // เส้นทางไปยังไฟล์ฟอนต์
  variable: "--font-geist-sans", // ชื่อ CSS Variable ที่จะใช้
  weight: "100 900", // น้ำหนักฟอนต์ที่รองรับ
});

// กำหนดฟอนต์ Geist Mono โดยระบุเส้นทางและน้ำหนัก
const geistMono = localFont({
  src: "./fonts/GeistMonoVF.woff", // เส้นทางไปยังไฟล์ฟอนต์
  variable: "--font-geist-mono", // ชื่อ CSS Variable ที่จะใช้
  weight: "100 900", // น้ำหนักฟอนต์ที่รองรับ
});

// กำหนด metadata ของแอปพลิเคชัน เช่น ชื่อและคำอธิบาย
export const metadata: Metadata = {
  title: "Your App Title", // ชื่อของแอปพลิเคชัน
  description: "Your App Description", // คำอธิบายของแอปพลิเคชัน
  // สามารถเพิ่มฟิลด์ metadata อื่นๆ ได้ เช่น keywords
};

// ฟังก์ชัน RootLayout ที่ใช้เป็นเลเอาต์หลักของแอป
export default function RootLayout({
  children, // children คือส่วนที่อยู่ภายใน RootLayout
}: Readonly<{
  children: React.ReactNode; // กำหนดประเภทของ children เป็น React Node
}>) {
  return (
    <html lang="en">
      <body className={`${geistSans.variable} ${geistMono.variable} antialiased`}> {/* ใช้งานฟอนต์และเพิ่มการป้องกันการกระพริบ */}
        {children} {/* แสดงเนื้อหาภายใน RootLayout */}
      </body>
    </html>
  );
}
