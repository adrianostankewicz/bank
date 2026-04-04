package account_test

import (
	"testing"

	"github.com/adrianostankewicz/bank/internal/core/domain/account"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryAccountRepositorySave(t *testing.T) {
	repo := account.NewInMemoryAccountRepository()
	a := account.NewBaseAccount("1")

	err := repo.Save(a)

	assert.Nil(t, err)
}

func TestInMemoryAccountRepositoryFindById(t *testing.T) {
	repo := account.NewInMemoryAccountRepository()
	a := account.NewBaseAccount("1")

	repo.Save(a)

	found, err := repo.FindById("1")

	assert.Nil(t, err)
	assert.Equal(t, "1", found.ID())
}

func TestInMemoryAccountRepositoryFindByIdNotFound(t *testing.T) {
	repo := account.NewInMemoryAccountRepository()
	a := account.NewBaseAccount("1")

	repo.Save(a)

	_, err := repo.FindById("2")

	assert.Error(t, err)
}

func TestInMemoryAccountRepositoryReset(t *testing.T) {
	repo := account.NewInMemoryAccountRepository()
	a := account.NewBaseAccount("1")

	repo.Save(a)

	repo.Reset()

	_, err := repo.FindById(a.ID())

	assert.Error(t, err)
}
