package domain

import "time"

// 维护物品信息
type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

// 维护订单信息
type Order struct {
	ID         int64       `json:"id"`          // 订单id
	CustomerID int64       `json:"customer_id"` // 客户ID
	Status     string      `json:"status"`      // 订单状态
	OrderItems []OrderItem `json:"order_item"`  // 购买的物品(列表)
	CreatedAt  int64       `json:"created_at"`  // 订单创建时间
}

// 创建新订单
func NewOrder(customerID int64, orderItems []OrderItem) Order {
	return Order{
		Status:     "Pending",
		CreatedAt:  time.Now().Unix(),
		CustomerID: customerID,
		OrderItems: orderItems,
	}
}
