# ใช้ Node.js base image
FROM node:18

# กำหนด working directory ใน container
WORKDIR /app

# คัดลอกไฟล์ package.json และติดตั้ง dependencies
COPY package*.json ./
RUN npm install

# คัดลอก source code ทั้งหมด
COPY . .

# สร้าง production build
RUN npm run build

# ใช้คำสั่ง start เพื่อรันแอป
CMD ["npm", "run", "start"]

# เปิดพอร์ตที่ต้องการ
EXPOSE 3000
