package cart

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
func (r Repository) CreateCart(request model.RequestCreateCart) (model.Cart, error) {
	cart := model.Cart{
		UserID: request.UserID,
	}
	// บันทึก cart ลงฐานข้อมูล
	if err := r.Database.Create(&cart).Error; err != nil {
		return model.Cart{}, err
	}

	// เพิ่มรายการสินค้าใน cart
	for _, product := range request.Products {
		cartItem := model.CartItem{
			CartID:    cart.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}
		if err := r.Database.Create(&cartItem).Error; err != nil {
			return model.Cart{}, err
		}
	}
	return cart, nil
}

func (r Repository) GetAllCart() ([]model.Cart, error) {
	var carts []model.Cart
	err := r.Database.Preload("Products").Find(&carts).Error
	return carts, err
}

func (r Repository) GetCartById(id uint) (model.Cart, error) {
	var cart model.Cart
	err := r.Database.Preload("Products").First(&cart, id).Error
	return cart, err
}

func (r Repository) UpdateCart(id uint, request model.RequestCreateCart) (model.Cart, error) {
	var cart model.Cart
	if err := r.Database.First(&cart, id).Error; err != nil {
		return model.Cart{}, err
	}

	// อัปเดตข้อมูล cart ตามคำขอ
	cart.UserID = request.UserID
	if err := r.Database.Save(&cart).Error; err != nil {
		return model.Cart{}, err
	}

	// ลบสินค้าเก่าใน cart และเพิ่มสินค้าใหม่
	r.Database.Where("cart_id = ?", cart.ID).Delete(&model.CartItem{})
	for _, product := range request.Products {
		cartItem := model.CartItem{
			CartID:    cart.ID,
			ProductID: product.ProductID,
			Quantity:  product.Quantity,
		}
		if err := r.Database.Create(&cartItem).Error; err != nil {
			return model.Cart{}, err
		}
	}
	return cart, nil
}

func (r Repository) DeleteCart(id uint) error {
	return r.Database.Delete(&model.Cart{}, id).Error
}
