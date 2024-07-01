package gateway

import "github.com/marcohnp/fullcycle_eda/internal/entity"

type BalanceGateway interface {
	Save(balance *entity.Balance) error
	FindByID(id string) (*entity.Balance, error)
}
