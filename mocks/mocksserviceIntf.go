package mocks

import (
	"context"
	"userservice/models"
)

type Service interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}
