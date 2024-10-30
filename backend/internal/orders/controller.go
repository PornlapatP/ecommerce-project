package orders

import (
	"ecommerce-backend/internal/model"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Claims struct {
	ID uint `json:"id"` // หรือฟิลด์อื่น ๆ ที่คุณต้องการ
	jwt.StandardClaims
}

type Controller struct {
	Service Service
}

func NewController(dbconn *gorm.DB) Controller {
	return Controller{
		Service: NewService(dbconn),
	}
}

func parseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
func (ctrl Controller) CreateOrder(ctx *gin.Context) {
	var request model.RequestCreateOrder

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แปลงจาก OrderItemRequest เป็น OrderItem
	items := ConvertOrderItemRequestToOrderItem(request.Items)

	// เรียกใช้ service เพื่อสร้างคำสั่งซื้อ
	order, err := ctrl.Service.Create(request.UserID, items, request.TotalAmount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order created successfully", "orderId": order.ID})
}
func (ctrl Controller) GetAllOrder(ctx *gin.Context) {
	var orders []model.Order

	// ดึงข้อมูลคำสั่งซื้อทั้งหมดจาก repository
	if err := ctrl.Service.GetAllOrders(&orders); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลลัพธ์กลับไปยังผู้ใช้
	ctx.JSON(http.StatusOK, orders)
}

// GetOrderById ดึงข้อมูลคำสั่งซื้อโดยใช้ ID
func (ctrl Controller) GetOrderById(ctx *gin.Context) {
	// รับ ID จาก URL parameters
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// ดึงข้อมูลคำสั่งซื้อจาก service
	order, err := ctrl.Service.GetOrderByID(uint(orderID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// ส่งผลลัพธ์กลับไปยังผู้ใช้
	ctx.JSON(http.StatusOK, order)
}

// UpdateOrder อัปเดตข้อมูลคำสั่งซื้อโดยใช้ ID
func (ctrl Controller) UpdateOrder(ctx *gin.Context) {
	// รับ ID จาก URL parameters
	orderID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var request model.RequestUpdateOrder
	// Bind JSON request body to the request model
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตคำสั่งซื้อใน service
	order, err := ctrl.Service.UpdateOrder(uint(orderID), request)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found or update failed"})
		return
	}

	// ส่งผลลัพธ์กลับไปยังผู้ใช้
	ctx.JSON(http.StatusOK, order)
}
func (ctrl Controller) DeleteOrder(ctx *gin.Context) {
	// ดึง ID ของคำสั่งซื้อจาก URL parameters
	orderID := ctx.Param("id")

	// เรียกใช้บริการเพื่อลบคำสั่งซื้อ
	if err := ctrl.Service.DeleteOrder(orderID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งผลลัพธ์กลับไปยังผู้ใช้
	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
func (ctrl Controller) UpdateStatusOrder(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)
	var (
		request model.RequestUpdateStatusOrder
	)

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Bind body").Error(),
			},
		)
		return
	}

	// Validate
	if err := validator.New().Struct(&request); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Validate").Error(),
			},
		)
		return
	}
	// Replace
	result, err := ctrl.Service.UpdateStatus(uint(id), request.Status)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			model.BaseResponse[any]{
				Message: errors.Wrap(err, "Update item status").Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		model.BaseResponse[model.Order]{
			Data: result,
		},
	)
}
