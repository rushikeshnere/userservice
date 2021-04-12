package service 

import (
	"userservice/repository"
	"userservice/models"
	"context"
	//"fmt"
	"errors"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) Service {
	return Service {
		repository: repo, 
	}
}

func (s *Service) CreateUser(ctx context.Context, user models.User) (string, error) {
	result := s.repository.CreateUser(ctx, user)
	if(result.Error != nil) {
		return "", result.Error
	}
	return user.Id, nil
}

func (s *Service) GetUser(ctx context.Context, id string) (*models.User, error) {
	user, result := s.repository.GetUser(ctx, id)
	if(result.Error != nil) {
		return nil, result.Error
	}
	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	result, err := s.repository.UpdateUser(ctx, user)
	if(err != nil) {
		return nil, err
	} else if(result.Error != nil) {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func (s *Service) DeleteUser(ctx context.Context, id string) (string, error) {
	result := s.repository.DeleteUser(ctx, id)
	if(result.Error != nil) {
		return "Error occurred while deleting a user", result.Error
	} else {
		if(result.RowsAffected == 0) {
			return "User not present", errors.New("User not present")
		} else {
			return "User deleted successfully", nil
		}
	}
}
	