package user_repository

import (
	"context"

	"github.com/io-m/hyphen/internal/shared/models"
)

func (ur *postgresUserRepository) FindAllUsers(ctx context.Context) ([]models.User, error) {
	return []models.User{models.EmptyUser()}, nil
}
func (ur *postgresUserRepository) FindUserById(ctx context.Context, userId uint) (models.User, error) {
	return models.EmptyUser(), nil
}
func (ur *postgresUserRepository) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	return models.EmptyUser(), nil
}
func (ur *postgresUserRepository) CreateUser(ctx context.Context, user *models.User) (models.User, error) {
	return models.EmptyUser(), nil
}
func (ur *postgresUserRepository) UpdateUserById(ctx context.Context, userId uint, userRequest *models.UserUpdateRequest) (models.User, error) {
	return models.EmptyUser(), nil
}
func (ur *postgresUserRepository) DeleteUserById(ctx context.Context, userId uint) (bool, error) {
	return false, nil
}
