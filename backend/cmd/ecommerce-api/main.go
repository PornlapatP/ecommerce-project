package main

import (
	"log"
	"os"

	"ecommerce-backend/internal/products"
	"ecommerce-backend/internal/users"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}
	DATABASE := os.Getenv("DATABASE_URL")

	dbconn, err := gorm.Open(postgres.Open(DATABASE), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connection to the database: %v", err)
	}
	controlleruser := users.NewController(dbconn, os.Getenv("JWT_SECRET"), os.Getenv("HASH_SECRET"))
	controllerproduct := products.NewController(dbconn)
	// controllerorder := orders.NewController(dbconn)
	// controllerorder_item := order_items.NewController(dbconn)
	// controllercart := cart.NewController(dbconn)

	r := gin.Default()
	// เสิร์ฟไฟล์จากโฟลเดอร์ uploads
	r.Static("/uploads", "./uploads")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"http://localhost:3001",
		"http://127.0.0.1:3001",
	}
	r.Use(cors.New(config))

	// router user
	users := r.Group("/users")
	{
		users.POST("/register", controlleruser.Createuser)
		users.POST("/login", controlleruser.Login)
		users.GET("/:id", controlleruser.GetUserById)
		users.PUT("/:id", controlleruser.UpdateUser)
		users.PATCH("/:id", controlleruser.UpdateRole)
		// users.DELETE("/:id", controlleruser.DeleteUser)
	}
	items := r.Group("/items")
	{
		items.POST("/products", controllerproduct.CreateProduct)
		items.GET("/products", controllerproduct.GetAllProduct)
		// items.GET("/products/:id", controllerproduct.CreateProduct)
		// items.PUT("/products/:id", controllerproduct.CreateProduct)
		// items.DELETE("/products/:id", controllerproduct.CreateProduct)
		// items.PATCH("/products/:id", controllerproduct.UpdateStatus)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
