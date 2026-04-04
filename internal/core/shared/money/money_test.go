package money_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/shared/money"
	"github.com/stretchr/testify/assert"
)

func TestAddMoney(t *testing.T) {
	a := money.NewMoney(100)
	b := money.NewMoney(200)
	v := a.Add(b)

	assert.Equal(t, int64(300), v.Amount())
}

func TestSubMoney(t *testing.T) {
	a := money.NewMoney(200)
	b := money.NewMoney(100)
	v := a.Sub(b)

	assert.Equal(t, int64(100), v.Amount())
}

func TestIsZero(t *testing.T) {
	a := money.NewMoney(0)
	assert.True(t, a.IsZero())
}

func TestIsNegative(t *testing.T) {
	a := money.NewMoney(-100)
	assert.True(t, a.IsNegative())
}

func TestMoneyIsGreaterThan(t *testing.T) {
	a := money.NewMoney(100)
	b := money.NewMoney(50)
	assert.True(t, a.IsGreaterThan(b))
}

func TestMoneyIsNotGreaterThan(t *testing.T) {
	a := money.NewMoney(50)
	b := money.NewMoney(100)
	assert.False(t, a.IsGreaterThan(b))
}

func TestMoneyIsGreaterThanEqual(t *testing.T) {
	a := money.NewMoney(100)
	b := money.NewMoney(100)
	assert.False(t, a.IsGreaterThan(b))
}
