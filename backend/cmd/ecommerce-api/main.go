package main

import (
	"log"
	"os"

	"ecommerce-backend/internal/cart"
	"ecommerce-backend/internal/orders"
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
	controllerorder := orders.NewController(dbconn)
	// controllerorder_item := order_items.NewController(dbconn)
	controllercart := cart.NewController(dbconn)

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
		items.GET("/products/:id", controllerproduct.GetProductById)
		items.PUT("/products/:id", controllerproduct.UpdateProduct)
		items.DELETE("/products/:id", controllerproduct.DeleteProduct)
		items.PATCH("/products/:id", controllerproduct.UpdateStatusProduct)
	}

	orders := r.Group("/orders")
	{
		orders.POST("/products", controllerorder.CreateOrder)
		orders.GET("/products", controllerorder.GetAllOrder)
		orders.GET("/products/:id", controllerorder.GetOrderById)
		orders.PUT("/products/:id", controllerorder.UpdateOrder)
		orders.DELETE("/products/:id", controllerorder.DeleteOrder)
		orders.PATCH("/products/:id", controllerorder.UpdateStatusOrder)
	}
	//build and cd and wed server
	cart := r.Group("/cart")
	{
		cart.POST("/", controllercart.CreateCart)
		cart.GET("/", controllercart.GetAllCart)
		cart.GET("/:id", controllercart.GetCartById)
		cart.PUT("/:id", controllercart.UpdateCart)
		cart.DELETE("/:id", controllercart.DeleteCart)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
