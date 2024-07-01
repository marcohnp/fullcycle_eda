package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateBalace(t *testing.T) {
	balance := NewBalance("1", 100)
	balance.UpdateBalance(50)
	assert.Equal(t, float64(50), balance.Balance)
	assert.NotNil(t, balance.UpdatedAt)
}

func TestAddBalance(t *testing.T) {
	balance := NewBalance("1", 100)
	balance.AddBalance(50)
	assert.Equal(t, float64(150), balance.Balance)
	assert.NotNil(t, balance.UpdatedAt)
}
