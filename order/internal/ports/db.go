package ports

import "github.com/vdebu/market/order/internal/application/core/domain"

// DBPort 定义数据操纵规范
type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
}
