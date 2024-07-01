package get_balance_test

import (
	"errors"
	"testing"
	"time"

	"github.com/marcohnp/fullcycle_eda/internal/entity"
	"github.com/marcohnp/fullcycle_eda/internal/usecase/get_balance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) Save(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceGatewayMock) Update(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceGatewayMock) FindByID(id string) (*entity.Balance, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Balance), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetBalanceUsecase_Execute(t *testing.T) {
	mockBalanceGateway := new(BalanceGatewayMock)
	usecase := get_balance.NewGetBalanceUsecase(mockBalanceGateway)

	t.Run("should get balance successfully", func(t *testing.T) {
		accountId := "1"
		createdAt := time.Now().Add(-24 * time.Hour)
		updatedAt := time.Now()
		balance := &entity.Balance{
			AccountID: accountId,
			Balance:   100,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		mockBalanceGateway.On("FindByID", accountId).Return(balance, nil)

		input := get_balance.GetBalanceInputDto{AccountId: accountId}
		expectedOutput := &get_balance.GetBalanceOutputDto{
			AccountId: accountId,
			Balance:   100,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		output, err := usecase.Execute(input)

		assert.NoError(t, err)
		assert.Equal(t, expectedOutput, output)
		mockBalanceGateway.AssertExpectations(t)
	})

	t.Run("should return error if balance not found", func(t *testing.T) {
		accountId := "2"

		mockBalanceGateway.On("FindByID", accountId).Return(nil, errors.New("balance not found"))

		input := get_balance.GetBalanceInputDto{AccountId: accountId}

		output, err := usecase.Execute(input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "balance not found", err.Error())
		mockBalanceGateway.AssertExpectations(t)
	})
}
