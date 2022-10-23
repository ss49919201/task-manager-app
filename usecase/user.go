package usecase

import (
	"context"

	"github.com/s-beats/rest-todo/domain"
	"github.com/s-beats/rest-todo/domain/repository"
	"github.com/s-beats/rest-todo/util"
)

type User interface {
	Create(ctx context.Context, name string) (*domain.User, error)
}

type user struct {
	userRepository repository.User
}

func NewUser(userRepo repository.User) User {
	return &user{
		userRepository: userRepo,
	}
}

func (u *user) Create(ctx context.Context, name string) (*domain.User, error) {
	wrappedSave := util.ConvMapperWithCtx(ctx, u.userRepository.Save)

	result := domain.NewUser(domain.NewUserID(util.NewUUID()), name).FlatMap(wrappedSave)
	if result.Error() != nil {
		return nil, result.Error()
	}

	return result.MustGet(), nil
}
