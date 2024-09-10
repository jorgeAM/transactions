package infrastructure

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/jorgeAM/kata-transactions/domain"
)

var _ domain.UserRepository = (*PostgresUserRepository)(nil)

type postgresUser struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type PostgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (p *PostgresUserRepository) Save(ctx context.Context, user *domain.User) error {
	dto := &postgresUser{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	ds := goqu.
		Insert("kata_transactions.users").
		Rows(dto)

	sql, _, err := ds.ToSQL()
	if err != nil {
		return err
	}

	_, err = p.db.ExecContext(ctx, sql)
	if err != nil {
		return err
	}

	return nil
}
