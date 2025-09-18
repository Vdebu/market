// Package db 实现外部具体的数据库技术 -> 实际的数据库状态
package db

import (
	"fmt"
	"github.com/vdebu/market/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Order 基础订单信息
type Order struct {
	gorm.Model // 嵌入默认字段
	CustomerID int64
	Status     string
	OrderItems []domain.OrderItem
}

// OrderItem 基础物品信息
type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint
}

// 向适配器注入依赖
type Adapter struct {
	db *gorm.DB
}

// NewAdapter 初始化适配器
func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}
	// 开启自动迁移确保表结构被正常创建
	err = db.AutoMigrate(Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration err: %v", err)
	}
	return &Adapter{db: db}, nil
}

// 实现ports满足需要的服务

// Get 查询订单
func (a Adapter) Get(id string) (domain.Order, error) {
	// 存储结果
	var orderEntity domain.Order
	res := a.db.First(&orderEntity, id)
	var orderItems []domain.OrderItem
	// 数据转换
	for _, orderItem := range orderEntity.OrderItems {
		// ?
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	// 转换成最终的order进行返回
	order := domain.Order{
		ID:         int64(orderEntity.ID),
		CustomerID: orderEntity.CustomerID,
		Status:     orderEntity.Status,
		OrderItems: orderItems,
		CreatedAt:  orderEntity.CreatedAt,
	}
	return order, res.Error
}

// Save 保存订单
func (a Adapter) Save(order *domain.Order) error {
	var orderItems []domain.OrderItem
	// 将数据转换成实际需要的数据
	for _, orderItem := range orderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderItem.ProductCode,
			UnitPrice:   orderItem.UnitPrice,
			Quantity:    orderItem.Quantity,
		})
	}
	// 构建表模型
	orderModel := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	// 持久化到数据库
	res := a.db.Create(&orderModel)
	// 若未发生错误则返回创建的订单ID?
	if res.Error == nil {
		order.ID = int64(orderModel.ID)
	}
	return res.Error
}
