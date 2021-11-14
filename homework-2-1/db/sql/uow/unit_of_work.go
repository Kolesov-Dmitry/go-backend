package uow

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime/debug"
)

type UnitOfWork struct {
	db *sql.DB
	tx *sql.Tx
}

func NewUnitOfWork(db *sql.DB) *UnitOfWork {
	return &UnitOfWork{
		db: db,
		tx: nil,
	}
}

func (uow UnitOfWork) WithTx(ctx context.Context, f func(uow UnitOfWork) error) (err error) {
	if uow.tx != nil {
		return f(uow)
	}

	return uow.withBeginTx(ctx, f)
}

func (uow UnitOfWork) Tx() *sql.Tx {
	return uow.tx
}

func (uow UnitOfWork) withBeginTx(ctx context.Context, f func(uow UnitOfWork) error) (err error) {
	var newTx *sql.Tx

	if newTx, err = uow.db.BeginTx(ctx, nil); err != nil {
		return err
	}

	nuow := uow
	nuow.tx = newTx

	commit := false
	defer func() {
		if r := recover(); r != nil || !commit {
			if r != nil {
				log.Printf("!!! TRANSACTION PANIC !!! : %s\n%s", r, string(debug.Stack()))
			}

			if e := newTx.Rollback(); e != nil {
				err = e
			} else if r != nil {
				err = fmt.Errorf("transaction panic: %s", r)
			}
		} else if commit {
			err = newTx.Commit()
		}
	}()

	if err = f(nuow); err != nil {
		return err
	}

	commit = true
	return nil
}
