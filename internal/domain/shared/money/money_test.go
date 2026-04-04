package money_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/domain/shared/money"
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
