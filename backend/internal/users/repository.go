package users

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

func (repo Repository) CreateUser(user *model.User) error {
	return repo.Database.Create(user).Error
}

func (repo Repository) FindOneByEmail(email string) (model.User, error) {
	var result model.User

	db := repo.Database
	db = db.Where("email = ?", email)

	if err := db.Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}
func (repo Repository) GetByID(id uint) (model.User, error) {
	var result model.User
	if err := repo.Database.First(&result, id).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (repo Repository) Update(user model.User) error {
	return repo.Database.Model(&user).Updates(user).Error
}
