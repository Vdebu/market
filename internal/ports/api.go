// Package ports 定义提供的服务与需要什么服务
package ports

import "github.com/vdebu/microservice/order/internal/application/core/domain"

// APIPort API提供的服务
type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
