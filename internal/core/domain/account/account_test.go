package account_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	a := account.NewBaseAccount("1")

	assert.NotEmpty(t, a.ID())
	assert.Equal(t, "1", a.ID())
	assert.NotZero(t, a.CreatedAt())
	assert.NotZero(t, a.UpdatedAt())
	assert.Nil(t, a.DeletedAt())
}
