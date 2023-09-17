package user_repository

import (
	"context"

	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/jmoiron/sqlx"
)

type postgresUserRepository struct {
	postgres *sqlx.DB
	// redis    *redis.Client
}

type IUserRepository interface {
	FindAllUsers(ctx context.Context) ([]models.User, error)
	FindUserById(ctx context.Context, userId uint) (models.User, error)
	FindUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) (models.User, error)
	UpdateUserById(ctx context.Context, userId uint, userRequest *models.UserUpdateRequest) (models.User, error)
	DeleteUserById(ctx context.Context, userId uint) (bool, error)
}
