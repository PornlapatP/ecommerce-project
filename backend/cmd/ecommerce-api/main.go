package main

import (
	"log"
	"os"

	"ecommerce-backend/internal/cart"
	"ecommerce-backend/internal/middleware"
	"ecommerce-backend/internal/orders"
	"ecommerce-backend/internal/products"
	"ecommerce-backend/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// โหลดค่าจาก .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File", err)
	}

	// อ่านค่า DATABASE_URL จาก environment
	DATABASE := os.Getenv("DATABASE_URL")
	if DATABASE == "" {
		log.Fatal("DATABASE_URL not set in .env file")
	}

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL
	dbconn, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// ตรวจสอบการเชื่อมต่อฐานข้อมูล
	err = dbconn.Exec("SELECT 1").Error
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// สร้าง Controller สำหรับแต่ละฟังก์ชัน
	controlleruser := users.NewController(dbconn, os.Getenv("JWT_SECRET"), os.Getenv("HASH_SECRET"))
	controllerproduct := products.NewController(dbconn)
	controllerorder := orders.NewController(dbconn)
	controllercart := cart.NewController(dbconn)

	// สร้าง Gin router
	r := gin.Default()

	// เสิร์ฟไฟล์จากโฟลเดอร์ uploads
	r.Static("/uploads", "./uploads")

	// ใช้ middleware สำหรับ log, CORS และ error handling
	r.Use(middleware.RequestLogger())  // Logging middleware
	r.Use(middleware.CORSMiddleware()) // CORS middleware
	r.Use(middleware.ErrorHandler())   // Error handling middleware

	// กำหนด routing สำหรับ users
	users := r.Group("/users")
	{
		users.POST("/register", controlleruser.Createuser)
		users.POST("/login", controlleruser.Login)
		users.GET("/:id", controlleruser.GetUserById)
		users.PUT("/:id", controlleruser.UpdateUser)
		users.PATCH("/:id", controlleruser.UpdateRole)
	}

	// กำหนด routing สำหรับ items (ตรวจสอบ JWT ด้วย Guard middleware)
	items := r.Group("/items")
	items.Use(middleware.Guard(os.Getenv("JWT_SECRET")))
	{
		items.POST("/products", controllerproduct.CreateProduct)
		items.GET("/products", controllerproduct.GetAllProduct)
		items.GET("/products/:id", controllerproduct.GetProductById)
		items.PUT("/products/:id", controllerproduct.UpdateProduct)
		items.DELETE("/products/:id", controllerproduct.DeleteProduct)
		items.PATCH("/products/:id", controllerproduct.UpdateStatusProduct)
	}

	// กำหนด routing สำหรับ orders (ตรวจสอบ JWT ด้วย Guard middleware)
	orders := r.Group("/orders")
	orders.Use(middleware.Guard(os.Getenv("JWT_SECRET")))
	{
		orders.POST("/products", controllerorder.CreateOrder)
		orders.GET("/products", controllerorder.GetAllOrder)
		orders.GET("/products/:id", controllerorder.GetOrderById)
		orders.PUT("/products/:id", controllerorder.UpdateOrder)
		orders.DELETE("/products/:id", controllerorder.DeleteOrder)
		orders.PATCH("/products/:id", controllerorder.UpdateStatusOrder)
	}

	// กำหนด routing สำหรับ cart (ตรวจสอบ JWT ด้วย Guard middleware)
	cart := r.Group("/cart")
	cart.Use(middleware.Guard(os.Getenv("JWT_SECRET")))
	{
		cart.POST("/", controllercart.CreateCart)
		cart.GET("/", controllercart.GetAllCart)
		cart.GET("/:id", controllercart.GetCartById)
		cart.PUT("/:id", controllercart.UpdateCart)
		cart.DELETE("/:id", controllercart.DeleteCart)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// อ่านพอร์ตจาก .env หากไม่พบใช้ค่า default เป็น 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // ค่า default
	}

	// เริ่มเซิร์ฟเวอร์
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
