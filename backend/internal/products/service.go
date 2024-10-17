package products

import (
	"ecommerce-backend/internal/constant"
	"ecommerce-backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	Validata   Validate
}

func NewService(dbconn *gorm.DB) Service {
	return Service{
		Repository: NewRepository(dbconn),
	}
}

func (service Service) Create(req model.RequestCreateProduct, imageURL string) (model.Product, error) {
	now := time.Now()

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    imageURL, // URL ของรูปที่ได้รับมาจาก Controller
		Status:      constant.ProductActiveStatus,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := service.Repository.CreateProduct(&product); err != nil {
		return model.Product{}, err
	}
	return product, nil
}
