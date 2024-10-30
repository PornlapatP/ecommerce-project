package cart

import (
	"ecommerce-backend/internal/model"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	// Validate   Validate
}

func NewService(dbconn *gorm.DB) Service {
	return Service{
		Repository: NewRepository(dbconn),
	}
}
func (s Service) CreateCart(request model.RequestCreateCart) (model.Cart, error) {
	return s.Repository.CreateCart(request)
}

func (s Service) GetAllCart() ([]model.Cart, error) {
	return s.Repository.GetAllCart()
}

func (s Service) GetCartById(id uint) (model.Cart, error) {
	return s.Repository.GetCartById(id)
}

func (s Service) UpdateCart(id uint, request model.RequestCreateCart) (model.Cart, error) {
	return s.Repository.UpdateCart(id, request)
}

func (s Service) DeleteCart(id uint) error {
	return s.Repository.DeleteCart(id)
}
