package common

import (
	"context"

	"github.com/jmoiron/sqlx"
)

var _ Transactional = (*PostgresTransactionalManager)(nil)

type PostgresTransactionalManager struct {
	db *sqlx.DB
}

func NewPostgresTransactionalManager(db *sqlx.DB) *PostgresTransactionalManager {
	return &PostgresTransactionalManager{
		db: db,
	}
}

func (p *PostgresTransactionalManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := p.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(context.WithValue(ctx, TxKey("tx"), tx)); err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}
