package domain

import "context"

type Notification struct {
	ID      int
	UserID  int
	Message string
}

type NotificationRepository interface {
	Save(ctx context.Context, notification *Notification) error
}
