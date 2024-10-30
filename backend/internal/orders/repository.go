package orders

import (
	"ecommerce-backend/internal/model"

	"gorm.io/gorm"
)

type Repository struct {
	Database *gorm.DB
}

func NewRepository(dbconn *gorm.DB) Repository {
	return Repository{
		Database: dbconn,
	}
}

func (repo Repository) CreateOrder(order *model.Order) error {
	// บันทึก order พร้อมกับ order items
	return repo.Database.Create(order).Error
}

//	func (repo Repository) GetOrderByID(orderID uint) (*model.Order, error) {
//		var order model.Order
//		err := repo.Database.Preload("Product").First(&order, orderID).Error
//		return &order, err
//	}
func (repo Repository) GetAllOrders(orders *[]model.Order) error {
	return repo.Database.Preload("Items.Product").Find(orders).Error // Preload Product ใน Items
}

// GetOrderByID ดึงข้อมูลคำสั่งซื้อจากฐานข้อมูลโดยใช้ ID
func (repo Repository) GetOrderByID(orderID uint) (*model.Order, error) {
	var order model.Order
	err := repo.Database.Preload("Items.Product").First(&order, orderID).Error // Preload เพื่อดึงข้อมูล Items และ Product
	return &order, err
}

// UpdateOrder อัปเดตข้อมูลคำสั่งซื้อในฐานข้อมูล
func (repo Repository) UpdateOrder(orderID uint, req model.RequestUpdateOrder) (*model.Order, error) {
	var order model.Order
	// ค้นหาคำสั่งซื้อที่ต้องการอัปเดต
	if err := repo.Database.First(&order, orderID).Error; err != nil {
		return nil, err // ไม่พบคำสั่งซื้อ
	}

	// อัปเดตข้อมูลคำสั่งซื้อ
	order.TotalAmount = req.TotalAmount
	order.Status = req.Status
	// คุณอาจจะต้องอัปเดตข้อมูล Items ถ้าจำเป็น

	// บันทึกการอัปเดตลงฐานข้อมูล
	if err := repo.Database.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
func (repo Repository) DeleteOrder(orderID string) error {
	// ทำการลบคำสั่งซื้อโดยใช้ ID
	return repo.Database.Delete(&model.Order{}, orderID).Error
}
func (repo Repository) ReplaceStatus(data *model.Order) error {
	return repo.Database.Model(data).Updates(data).Error
}
