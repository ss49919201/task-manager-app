package domain

import (
	"github.com/s-beats/rest-todo/util"
	"github.com/samber/mo"
)

type User struct {
	id   UserID
	name string
}

func NewUser(id UserID, name string) mo.Result[*User] {
	return ToOKUser(
		&User{
			id:   id,
			name: name,
		},
	)
}

func (u User) ID() UserID {
	return u.id
}

type UserID struct {
	id string
}

func (u UserID) String() string {
	return u.id
}

func NewUserID(id string) UserID {
	return UserID{
		id: util.UUIDMustParse(id),
	}
}

func (u *User) Name() string {
	return u.name
}

func ToOKUser(v *User) mo.Result[*User] {
	return mo.Ok(v)
}

func ToErrUser(err error) mo.Result[*User] {
	return mo.Err[*User](err)
}
