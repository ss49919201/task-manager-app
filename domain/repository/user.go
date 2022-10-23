package repository

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
)

type User interface {
	Save(ctx context.Context, user *domain.User) error
	SaveWithTx(ctx context.Context, user *domain.User) error

	GetOne(ctx context.Context, userID domain.UserID) (*domain.User, error)
}
