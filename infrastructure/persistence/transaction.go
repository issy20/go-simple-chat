package persistence

import (
	"context"
	"database/sql"

	"github.com/issy20/go-simple-chat/infrastructure/database"
	"github.com/issy20/go-simple-chat/transaction"
)

var txKey = struct{}{}

type tx struct {
	conn *database.Conn
}

func NewTransaction(conn *database.Conn) transaction.Transaction {
	return &tx{
		conn: conn,
	}
}

func (t *tx) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := t.conn.DB.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, &txKey, tx)
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}
