package persistence

import (
	"context"
	"time"

	"github.com/s-beats/rest-todo/domain"
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

func (u *user) Save(ctx context.Context, user *domain.User) error {
	session := u.db.NewSession().Context(ctx)
	return u.save(session, user)
}

func (u *user) SaveWithTx(ctx context.Context, user *domain.User) error {
	// TODO: トランザクションの実装
	return nil
}

func (u *user) save(session *xorm.Session, user *domain.User) error {
	userDTO := &UserDTO{
		ID:        user.ID().String(),
		Name:      user.Name(),
		UpdatedAt: time.Now(),
	}

	exists, err := session.Table(userDTO.TableName()).Where("id = ?", userDTO.ID).Exist()
	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	if exists {
		_, err := session.ID(userDTO.ID).Update(userDTO)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	} else {
		userDTO.CreatedAt = userDTO.UpdatedAt
		_, err := session.Insert(userDTO)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	}

	return nil
}

func (u *user) GetOne(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	session := u.db.NewSession().Context(ctx)
	userDTO := UserDTO{ID: userID.String()}
	has, err := session.Get(&userDTO)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, xerrors.New("user not found")
	}
	return domain.NewUser(
		domain.NewUserID(userDTO.ID),
		userDTO.Name,
	), nil
}
