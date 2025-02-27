package products

import (
	"ecommerce-backend/internal/constant"
	"strings"

	"github.com/pkg/errors"
)

type Validate struct {
}

func NewValidate() Validate {
	return Validate{}
}

func (validata Validate) ProductStatus(current, next constant.ProductsStatus) error {
	switch current {
	case constant.ProductActiveStatus:
		if next == constant.ProductActiveStatus {
			return errors.New("cannot change active status to active status")
		}
	default:
		return errors.Errorf(
			"cannot change %s status to %s status",
			strings.ToLower(string(current)),
			strings.ToLower(string(next)),
		)
	}
	return nil
}
func (validator Validate) DeleteProduct(current constant.ProductsStatus) error {
	switch current {
	case constant.ProductActiveStatus:
		return errors.Errorf("cannot delete item when status is %s", strings.ToLower(string(current)))

	}

	return nil
}
func (validator Validate) UpdateProduct(current constant.ProductsStatus) error {
	switch current {
	case constant.ProductActiveStatus:
		return errors.Errorf("cannot update item when status is %s", strings.ToLower(string(current)))
	}

	return nil
}
