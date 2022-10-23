package persistence

import (
	"context"
	"time"

	"github.com/s-beats/rest-todo/domain"
	"github.com/samber/mo"
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type user struct {
	db *xorm.Engine
}

func NewUser(db *xorm.Engine) *user {
	return &user{
		db: db,
	}
}

type UserDTO struct {
	ID        string `xorm:"'id'"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *UserDTO) TableName() string {
	return "users"
}

func (u *user) Save(ctx context.Context, user *domain.User) mo.Result[*domain.User] {
	session := u.db.NewSession().Context(ctx)
	return u.save(session, user)
}

func (u *user) SaveWithTx(ctx context.Context, user *domain.User) mo.Result[*domain.User] {
	// TODO: トランザクションの実装
	panic("not implemented")
}

func (u *user) save(session *xorm.Session, user *domain.User) mo.Result[*domain.User] {
	userDTO := &UserDTO{
		ID:        user.ID().String(),
		Name:      user.Name(),
		UpdatedAt: time.Now(),
	}

	exists, err := session.Table(userDTO.TableName()).Where("id = ?", userDTO.ID).Exist()
	if err != nil {
		return domain.ToErrUser(xerrors.Errorf("%v", err))
	}

	if exists {
		_, err := session.ID(userDTO.ID).Update(userDTO)
		if err != nil {
			return domain.ToErrUser(xerrors.Errorf("%v", err))
		}
	} else {
		userDTO.CreatedAt = userDTO.UpdatedAt
		_, err := session.Insert(userDTO)
		if err != nil {
			return domain.ToErrUser(xerrors.Errorf("%v", err))
		}
	}

	return domain.ToOKUser(user)
}

func (u *user) GetOne(ctx context.Context, userID domain.UserID) mo.Result[*domain.User] {
	session := u.db.NewSession().Context(ctx)
	userDTO := UserDTO{ID: userID.String()}
	has, err := session.Get(&userDTO)
	if err != nil {
		return domain.ToErrUser(err)
	}
	if !has {
		return domain.ToErrUser(xerrors.New("user not found"))
	}
	return domain.NewUser(
		domain.NewUserID(userDTO.ID),
		userDTO.Name,
	)
}
