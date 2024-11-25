package products

import (
	"ecommerce-backend/internal/constant"
	"ecommerce-backend/internal/model"
	"time"

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

func (service Service) Create(req model.RequestCreateProduct, imageURL string) (model.Product, error) {
	now := time.Now()

	product := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    imageURL,
		Status:      constant.ProductActiveStatus, // ใช้ค่า Status จาก Request
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
	if err := service.Validate.DeleteProduct(product.Status); err != nil {
		return err
	}

	return service.Repository.DeleteProduct(id)
}

func (service Service) UpdateStatusProduct(id uint, status constant.ProductsStatus) (model.Product, error) {

	// Find item
	data, err := service.Repository.GetproductById(id)

	if err != nil {
		return model.Product{}, err
	}

	// Check status
	if err := service.Validate.ProductStatus(data.Status, status); err != nil {
		return model.Product{}, err
	}

	data.Status = status
	data.UpdatedAt = time.Now()

	// Replace
	if err := service.Repository.UpdateProduct(data); err != nil {
		return model.Product{}, err
	}

	return data, nil
}
func (service Service) UpdateProduct(id uint, req model.RequestCreateProduct) (model.Product, error) {
	result, err := service.Repository.GetproductById(id)

	if err := service.Validate.UpdateProduct(result.Status); err != nil {
		return model.Product{}, err
	}

	data := model.Product{
		ID:          result.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		ImageURL:    req.ImageURL, // URL ของรูปที่ได้รับมาจาก Controller
		Status:      constant.ProductActiveStatus,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   time.Now(),
	}
	if err := service.Repository.UpdateProduct(data); err != nil {
		return model.Product{}, err
	}

	return data, err
}
