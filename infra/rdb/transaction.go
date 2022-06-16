package rdb

import (
	"context"

	"xorm.io/xorm"
)

type Transaction interface {
	Do(ctx context.Context, f func()) error
}

type transaction struct {
	db *xorm.Engine
}

func (t *transaction) Do(ctx context.Context, f func() error) error {
	session := t.db.NewSession().Context(ctx)

	if err := session.Begin(); err != nil {
		return err
	}

	if err := f(); err != nil {
		if err := session.Rollback(); err != nil {
			return err
		}
	}

	if err := session.Commit(); err != nil {
		return err
	}

	return nil
}
