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

func (repo Repository) GetallProducts(query model.RequestGetProduct) ([]model.Product, error) {
	var results []model.Product

	tx := repo.Database

	if status := query.Status; len(status) > 0 {
		tx = tx.Where("status IN ?", query.Status)
	}

	if err := tx.Find(&results).Error; err != nil {
		return results, err
	}

	return results, nil
}

func (repo Repository) GetproductById(id uint) (model.Product, error) {
	var result model.Product

	if err := repo.Database.First(&result, id).Error; err != nil {
		return result, err
	}

	return result, nil

}

func (repo Repository) DeleteProduct(id uint) error {
	return repo.Database.Delete(&model.Product{}, id).Error
}
