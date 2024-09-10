package domain

import "context"

type User struct {
	ID    int
	Name  string
	Email string
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
}
