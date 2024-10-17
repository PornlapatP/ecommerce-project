package products

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

func (repo Repository) CreateProduct(product *model.Product) error {
	return repo.Database.Create(product).Error
}
