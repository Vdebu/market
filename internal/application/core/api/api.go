// Package api 纯粹的业务概念
package api

import "github.com/vdebu/microservice/order/internal/application/core/domain"

type Application struct {
	db ports.DBPort
}

// NewApplication 创建应用程序实例
func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

// PlaceOrder 创建新订单
func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	// 持久化存储
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	// 原路返回
	return order, nil
}
