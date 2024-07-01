package save_balance_test

import (
	"errors"
	"testing"
	"time"

	"github.com/marcohnp/fullcycle_eda/internal/entity"
	"github.com/marcohnp/fullcycle_eda/internal/usecase/save_balance"
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

func (m *BalanceGatewayMock) FindByID(id string) (*entity.Balance, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Balance), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestSaveBalanceUsecase_Execute(t *testing.T) {
	mockBalanceGateway := new(BalanceGatewayMock)
	usecase := save_balance.NewSaveUseCase(mockBalanceGateway)

	t.Run("should save balance successfully", func(t *testing.T) {
		accountId := "1"
		balanceAmount := 100.0
		createdAt := time.Now()
		updatedAt := time.Now()

		balance := &entity.Balance{
			AccountID: accountId,
			Balance:   balanceAmount,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		mockBalanceGateway.On("Save", mock.MatchedBy(func(b *entity.Balance) bool {
			return b.AccountID == accountId && b.Balance == balanceAmount
		})).Return(nil)

		input := save_balance.SaveBalanceInputDto{
			AccountId: accountId,
			Balance:   balanceAmount,
		}
		expectedOutput := &save_balance.SaveBalanceOutputDto{
			AccountId: accountId,
			Balance:   balanceAmount,
			CreatedAt: balance.CreatedAt, // Usando balance.CreatedAt diretamente
			UpdatedAt: balance.UpdatedAt, // Usando balance.UpdatedAt diretamente
		}

		output, err := usecase.Execute(input)

		assert.NoError(t, err)
		assert.Equal(t, expectedOutput.AccountId, output.AccountId)
		assert.Equal(t, expectedOutput.Balance, output.Balance)
		assert.WithinDuration(t, expectedOutput.CreatedAt, output.CreatedAt, time.Second)
		assert.WithinDuration(t, expectedOutput.UpdatedAt, output.UpdatedAt, time.Second)
		mockBalanceGateway.AssertExpectations(t)
	})

	t.Run("should return error if save fails", func(t *testing.T) {
		accountId := "2"
		balanceAmount := 200.0

		mockBalanceGateway.On("Save", mock.MatchedBy(func(b *entity.Balance) bool {
			return b.AccountID == accountId && b.Balance == balanceAmount
		})).Return(errors.New("some error"))

		input := save_balance.SaveBalanceInputDto{
			AccountId: accountId,
			Balance:   balanceAmount,
		}

		output, err := usecase.Execute(input)

		assert.Error(t, err)
		assert.Nil(t, output)
		assert.Equal(t, "some error", err.Error())
		mockBalanceGateway.AssertExpectations(t)
	})
}
