package infrastructure

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.com/jorgeAM/kata-transactions/domain"
)

var _ domain.NotificationRepository = (*PostgresNotificationRepository)(nil)

type postgresNotification struct {
	ID      int    `db:"id"`
	UserID  int    `db:"user_id"`
	Message string `db:"message"`
}

type PostgresNotificationRepository struct {
	db *sqlx.DB
}

func NewPostgresNotificationRepository(db *sqlx.DB) *PostgresNotificationRepository {
	return &PostgresNotificationRepository{
		db: db,
	}
}

func (p *PostgresNotificationRepository) Save(ctx context.Context, notification *domain.Notification) error {
	dto := &postgresNotification{
		ID:      notification.ID,
		UserID:  notification.UserID,
		Message: notification.Message,
	}

	ds := goqu.
		Insert("kata_transactions.notifications").
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
