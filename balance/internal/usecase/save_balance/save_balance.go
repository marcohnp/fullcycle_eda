package save_balance

import (
	"github.com/marcohnp/fullcycle_eda/internal/entity"
	"github.com/marcohnp/fullcycle_eda/internal/gateway"
	"time"
)

type SaveBalanceInputDto struct {
	AccountId string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type SaveBalanceOutputDto struct {
	AccountId string  `json:"account_id"`
	Balance   float64 `json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SaveBalanceUsecase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewSaveUseCase(balanceGateway gateway.BalanceGateway) *SaveBalanceUsecase {
	return &SaveBalanceUsecase{
		BalanceGateway: balanceGateway,
	}
}

func (u *SaveBalanceUsecase) Execute(input SaveBalanceInputDto) (*SaveBalanceOutputDto, error) {
	balance := entity.NewBalance(input.AccountId, input.Balance)
	err := u.BalanceGateway.Save(balance)
	if err != nil {
		return nil, err
	}
	return &SaveBalanceOutputDto{
		AccountId: balance.AccountID,
		Balance:   balance.Balance,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}, nil
}
