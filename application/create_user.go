package application

import (
	"context"

	"github.com/jorgeAM/kata-transactions/common"
	"github.com/jorgeAM/kata-transactions/domain"
)

type CreateUser struct {
	userRepository         domain.UserRepository
	notificationRepository domain.NotificationRepository
	transactionManager     common.Transactional
}

func NewCreateUser(
	userRepository domain.UserRepository,
	notificationRepository domain.NotificationRepository,
	transactionManager common.Transactional,
) *CreateUser {
	return &CreateUser{
		userRepository:         userRepository,
		notificationRepository: notificationRepository,
		transactionManager:     transactionManager,
	}
}

func (c *CreateUser) Exec(ctx context.Context) error {
	return c.transactionManager.WithinTransaction(ctx, func(ctx context.Context) error {
		user := &domain.User{
			ID:    1,
			Name:  "Jorge",
			Email: "jorge@gmail.com",
		}

		if err := c.userRepository.Save(ctx, user); err != nil {
			return err
		}

		notification := &domain.Notification{
			ID:      1,
			UserID:  user.ID,
			Message: "welcome!",
		}

		if err := c.notificationRepository.Save(ctx, notification); err != nil {
			return err
		}

		return nil
	})
}
