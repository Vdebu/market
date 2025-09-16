package ports

import "github.com/vdebu/microservice/order/internal/application/core/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(order *domain.Order) error
}
