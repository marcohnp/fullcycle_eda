package entity

import "time"

type Balance struct {
	AccountID string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBalance(accountId string, balance float64) *Balance {
	now := time.Now()
	return &Balance{
		AccountID: accountId,
		Balance:   balance,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (b *Balance) AddBalance(amount float64) {
	b.Balance += amount
	b.UpdatedAt = time.Now()
}

func (b *Balance) UpdateBalance(amount float64) {
	b.Balance = amount
	b.UpdatedAt = time.Now()
}
