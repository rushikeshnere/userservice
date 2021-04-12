package mocks 

import (
	"userservice/repository"
	"userservice/models"
	"context"
	//"fmt"
    "errors"
)

var userData = make(map[string]models.User)

type Mocksservice struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) Mocksservice {
	userData["1"] = models.User{Id: "1", Name: "Rushikesh", Dob: "1-1-1994", Address: "Pune", Description: "Engineer"}
	userData["2"] = models.User{Id: "2", Name: "Mayur", Dob: "1-1-1993", Address: "Pune", Description: "Person"}
	
	return Mocksservice {
		repository: repo, 
	}
}

func (s *Mocksservice) CreateUser(ctx context.Context, user models.User) (string, error) {
	if _, ok := userData[user.Id]; ok {
		return "", errors.New("ERROR: duplicate key value violates unique constraint")
	}
	userData[user.Id] = user
	return user.Id, nil
}

func (s *Mocksservice) GetUser(ctx context.Context, id string) (*models.User, error) {
	if val, ok := userData[id]; ok {
		return &val, nil
	}
	return nil, errors.New("record not present")
}

func (s *Mocksservice) UpdateUser(ctx context.Context, user models.User) (*models.User, error) {
	if _, ok := userData[user.Id]; ok {
		userData[user.Id] = user
		return &user, nil
	}
	return &user, errors.New("record not present")
}

func (s *Mocksservice) DeleteUser(ctx context.Context, id string) (string, error) {
	if _, ok := userData[id]; ok {
		delete(userData, id)
		return "User deleted successfully", nil
	}
	return "", errors.New("record not present")
}
	