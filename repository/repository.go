package repository

import (
	"context"
	"gorm.io/gorm"
	"userservice/models"
	"errors"
	//"fmt"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repository, error) {
	return &Repository{
		db: db,
	}, nil
}

func(repo *Repository) CreateUser(ctx context.Context, user models.User) *gorm.DB {
	result := repo.db.Create(&user)
	return result
}

func(repo *Repository) GetUser(ctx context.Context, id string)  (*models.User, *gorm.DB) {
	var user models.User
	result := repo.db.First(&user, id)
	return &user, result
}

func (repo *Repository) UpdateUser(ctx context.Context, user models.User) (*gorm.DB, error) {
	_, result := repo.GetUser(ctx, user.Id)
	if(result.Error == nil) {
		result := repo.db.Save(user)
		if(result.Error != nil) {
			return nil, result.Error
		}
		return result, nil
	} else {
		return nil, errors.New("record not found")
	}
}

func (repo *Repository) DeleteUser(ctx context.Context, id string) (*gorm.DB) {
	result := repo.db.Delete(&models.User{}, id)
	return result
}
