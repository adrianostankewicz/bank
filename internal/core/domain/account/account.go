package account

import (
	"time"
)

type Account interface {
	ID() string
	CreatedAt() time.Time
	UpdatedAt() time.Time
	DeletedAt() *time.Time
}

type BaseAccount struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time
}

func NewBaseAccount(id string) *BaseAccount {
	account := &BaseAccount{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
	return account
}

func (a BaseAccount) ID() string {
	return a.id
}

func (a BaseAccount) CreatedAt() time.Time {
	return a.createdAt
}

func (a BaseAccount) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a BaseAccount) DeletedAt() *time.Time {
	return a.deletedAt
}
