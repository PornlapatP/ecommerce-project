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

func (service Service) GetAllProduct(query model.RequestGetProduct) ([]model.Product, error) {
	return service.Repository.GetallProducts(query)
}

func (service Service) GetProductById(id uint) (model.Product, error) {
	product, err := service.Repository.GetproductById(id)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
func (service Service) DeleteProduct(id uint) error {
	// Find item
	product, err := service.Repository.GetproductById(id)
	if err != nil {
		return err
	}

	// Check status
	if err := service.Validata.DeleteProduct(product.Status); err != nil {
		return err
	}

	return service.Repository.DeleteProduct(id)
}
