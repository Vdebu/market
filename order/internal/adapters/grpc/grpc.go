package grpc

import (
	"context"
	"github.com/vdebu/market-proto/golang/order"
	"github.com/vdebu/market/order/internal/application/core/domain"
)

// Create 接收订单创建请求并进行领域模型的转换，对接实际的业务逻辑
func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	// 对数据结构进行转换
	var orderItems []domain.OrderItem
	for _, orderItem := range request.OrderItems {
		// 转换成本地域的模型
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	// 使用实际的API依赖方法创建订单
	res, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}
	// 返回对应的响应
	return &order.CreateOrderResponse{OrderId: res.ID}, nil
}
