package application

import (
	"context"

	"github.com/jorgeAM/kata-transactions/domain"
)

type CreateUser struct {
	userRepository         domain.UserRepository
	notificationRepository domain.NotificationRepository
}

func NewCreateUser(userRepository domain.UserRepository, notificationRepository domain.NotificationRepository) *CreateUser {
	return &CreateUser{
		userRepository:         userRepository,
		notificationRepository: notificationRepository,
	}
}

func (c *CreateUser) Exec(ctx context.Context) error {
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
		UserID:  1,
		Message: "welcome!",
	}

	if err := c.notificationRepository.Save(ctx, notification); err != nil {
		return err
	}

	return nil
}
