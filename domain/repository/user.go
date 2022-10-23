package repository

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/samber/mo"
)

type User interface {
	Save(ctx context.Context, user *domain.User) mo.Result[*domain.User]
	SaveWithTx(ctx context.Context, user *domain.User) mo.Result[*domain.User]

	GetOne(ctx context.Context, userID domain.UserID) mo.Result[*domain.User]
}
