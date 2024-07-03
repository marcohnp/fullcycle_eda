package get_balance

import (
	"errors"
	"fmt"
	"github.com/marcohnp/fullcycle_eda/internal/gateway"
	"time"
)

type GetBalanceInputDto struct {
	AccountId string `json:"account_id"`
}

type GetBalanceOutputDto struct {
	AccountId string    `json:"account_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetBalanceUsecase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUsecase(balanceGateway gateway.BalanceGateway) *GetBalanceUsecase {
	return &GetBalanceUsecase{
		BalanceGateway: balanceGateway,
	}
}

func (u *GetBalanceUsecase) Execute(input GetBalanceInputDto) (*GetBalanceOutputDto, error) {
	balance, err := u.BalanceGateway.FindByID(input.AccountId)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("balance not found")
	}

	return &GetBalanceOutputDto{
		AccountId: balance.AccountID,
		Balance:   balance.Balance,
		CreatedAt: balance.CreatedAt,
		UpdatedAt: balance.UpdatedAt,
	}, nil
}
