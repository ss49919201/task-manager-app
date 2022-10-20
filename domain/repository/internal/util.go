package internal

import (
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type Option func(session *xorm.Session)

func Upsert(session *xorm.Session, bean any, opts ...Option) error {
	for _, opt := range opts {
		opt(session)
	}

	affected, err := session.Update(bean)
	if err != nil {
		return xerrors.Errorf("%v", err)
	}
	if affected == 0 {
		_, err := session.InsertOne(bean)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	}
	return nil
}
