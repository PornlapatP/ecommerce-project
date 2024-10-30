package constant

type OrderStatus string

const (
	OrderPendingStatus  OrderStatus = "pending"
	OrderApprovedStatus OrderStatus = "approved"
	OrderRejectedStatus OrderStatus = "rejected"
)
