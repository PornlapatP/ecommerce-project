# Stage 1: Build stage
FROM golang:1.23.3-alpine3.19 AS builder

# ตั้งค่าตัวแปร environment
ENV GO111MODULE=on  
ENV CGO_ENABLED=0  
ENV GOOS=linux  
ENV GOARCH=amd64  

# ตั้งค่าโฟลเดอร์ทำงานใน container
WORKDIR /app  

# คัดลอก go.mod และ go.sum เพื่อตรวจสอบและดาวน์โหลด dependencies
COPY go.mod go.sum ./ 

# ดาวน์โหลด dependencies
RUN go mod tidy && go mod verify && go mod download

# คัดลอกไฟล์โค้ดทั้งหมด
COPY . .  

# ตรวจสอบว่าไฟล์ Go ถูกคัดลอกไปแล้ว
RUN ls -al /app  # ตรวจสอบว่าไฟล์ Go ถูกคัดลอกเข้ามาแล้ว

# คอมไพล์แอปพลิเคชัน Go
RUN go build -o ecommerce-app ./cmd/ecommerce-api/main.go  # ระบุพาธที่ถูกต้องไปยัง main.go

# Stage 2: Final stage
FROM alpine:latest  

# ตั้งค่าโฟลเดอร์ทำงาน
WORKDIR /root/

# คัดลอกไฟล์ binary จาก builder
COPY --from=builder /app/ecommerce-app ./ 

# คัดลอกไฟล์ .env
COPY .env ./ 

# ตั้งสิทธิ์ในการรันไฟล์ binary
RUN chmod +x ecommerce-app

# เปิดพอร์ต 2025 (ปรับให้ตรงกับ docker-compose.yml)
EXPOSE 2027  

# ใช้ JSON array แทน string
CMD ["./ecommerce-app"]
