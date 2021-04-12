package service

import (
	"context"
	"userservice/models"
)

type MockService interface {
	CreateUser(ctx context.Context, user models.User) (string, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) (string, error)
}
