package orders

import (
	"ecommerce-backend/internal/constant"
	"time"

	"ecommerce-backend/internal/model"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	Validate   Validate
}

func NewService(dbconn *gorm.DB) Service {
	return Service{
		Repository: NewRepository(dbconn),
	}
}

// สร้างคำสั่งซื้อ
func (service Service) Create(userID uint, items []model.OrderItem, totalAmount float64) (model.Order, error) {
	order := model.Order{
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      constant.OrderPendingStatus,
		Items:       items, // ใส่ items ที่ส่งเข้ามาเข้าไปใน order
	}

	// บันทึก order ลงฐานข้อมูล
	if err := service.Repository.CreateOrder(&order); err != nil {
		return model.Order{}, err
	}
	return order, nil
}
func (service Service) GetAllOrders(orders *[]model.Order) error {
	return service.Repository.GetAllOrders(orders)
}

// GetOrderByID ดึงข้อมูลคำสั่งซื้อจาก repository โดยใช้ ID
func (service Service) GetOrderByID(orderID uint) (*model.Order, error) {
	return service.Repository.GetOrderByID(orderID)
}

// UpdateOrder อัปเดตข้อมูลคำสั่งซื้อใน repository
func (service Service) UpdateOrder(orderID uint, req model.RequestUpdateOrder) (*model.Order, error) {
	return service.Repository.UpdateOrder(orderID, req)
}
func (service Service) DeleteOrder(orderID string) error {
	// เรียกใช้ repository เพื่อลบคำสั่งซื้อ
	return service.Repository.DeleteOrder(orderID)
}
func (service Service) UpdateStatus(id uint, status constant.OrderStatus) (model.Order, error) {
	// Find item
	order, err := service.Repository.GetOrderByID(id)
	if err != nil {
		return model.Order{}, err
	}

	// Check status
	if err := service.Validate.OrderStatusFlow(order.Status, status); err != nil {
		return model.Order{}, err
	}

	order.Status = status
	order.UpdatedAt = time.Now()

	// Replace
	if err := service.Repository.ReplaceStatus(order); err != nil {
		return model.Order{}, err
	}

	return *order, nil // Dereference เพื่อคืนค่าชนิด model.Order แทน *model.Order
}
