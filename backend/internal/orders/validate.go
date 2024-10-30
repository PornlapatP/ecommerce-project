package orders

import (
	"ecommerce-backend/internal/constant"
	"ecommerce-backend/internal/model"
	"strings"

	"github.com/pkg/errors"
)

type Validate struct {
}

func NewValidate() Validate {
	return Validate{}
}

// แปลงจาก OrderProductRequest เป็น OrderItem
func ConvertOrderItemRequestToOrderItem(requestItems []model.OrderProductRequest) []model.OrderItem {
	var orderItems []model.OrderItem
	for _, item := range requestItems {
		orderItem := model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems
}
func (validator Validate) OrderStatusFlow(current, next constant.OrderStatus) error {
	switch current {
	case constant.OrderPendingStatus:
		if next == constant.OrderPendingStatus {
			return errors.New("cannot change pending status to pending status")
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
