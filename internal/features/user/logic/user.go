package user_logic

import (
	"context"
	"errors"
	"fmt"

	user_repository "github.com/io-m/hyphen/internal/features/user/repository"
	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/io-m/hyphen/internal/shared/tokens"
	"github.com/io-m/hyphen/pkg/constants"
	"github.com/io-m/hyphen/pkg/helpers"
)

type IUserLogic interface {
	GetUserWithEmail(ctx context.Context, userId uint) (models.User, error)
	GetUserWithId(ctx context.Context, userId uint) (models.User, error)
	RegisterUser(ctx context.Context, user *models.User) (models.User, error)
	Login(ctx context.Context, user *models.User) (models.User, *string, *string, error)
	OAuth(ctx context.Context, user *models.User) (models.User, error)
	UpdateUserWithId(ctx context.Context, userId uint, userRequestOptional *models.UserUpdateRequest) (models.User, error)
	DeleteUserWithId(ctx context.Context, userId uint) (bool, error)
}

type userLogic struct {
	userRepository user_repository.IUserRepository
	tokens         tokens.ITokens
	protector      tokens.IProtector
}

func NewUserLogic(userRepository user_repository.IUserRepository, tokens tokens.ITokens, protector tokens.IProtector) *userLogic {
	return &userLogic{
		userRepository: userRepository,
		tokens:         tokens,
		protector:      protector,
	}
}

func (ul *userLogic) GetUserWithEmail(ctx context.Context, userId uint) (models.User, error) {
	return models.EmptyUser(), nil
}

func (ul *userLogic) GetUserWithId(ctx context.Context, userId uint) (models.User, error) {
	return models.EmptyUser(), nil
}

func (ul *userLogic) UpdateUserWithId(ctx context.Context, userId uint, userRequestOptional *models.UserUpdateRequest) (models.User, error) {
	return models.EmptyUser(), nil
}

func (ul *userLogic) DeleteUserWithId(ctx context.Context, userId uint) (bool, error) {
	return false, nil
}

func (ul *userLogic) RegisterUser(ctx context.Context, user *models.User) (models.User, error) {
	if err := helpers.ValidatePassword(user.Password); err != nil {
		return models.EmptyUser(), errors.New("password is invalid")
	}
	if err := helpers.ValidateEmail(user.Email); err != nil {
		return models.EmptyUser(), errors.New("email is invalid")
	}
	hashedPassword, err := helpers.HashPassword(user.Password)
	if err != nil {
		return models.EmptyUser(), fmt.Errorf("error while hashing password: %w", err)
	}
	user.Password = hashedPassword
	createdPerson, err := ul.userRepository.CreateUser(ctx, user)

	if err != nil {
		return models.EmptyUser(), err
	}
	return createdPerson, nil
}

func (ul *userLogic) Login(ctx context.Context, user *models.User) (models.User, *string, *string, error) {
	foundUser, err := ul.userRepository.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return models.EmptyUser(), nil, nil, err
	}
	if err := helpers.CheckPassword(user.Password, foundUser.Password); err != nil {
		return models.EmptyUser(), nil, nil, fmt.Errorf("wrong password: %w", err)
	}
	claims, err := tokens.NewClaims(user.ID, constants.ACCESS_TOKEN_DURATION)
	if err != nil {
		return models.EmptyUser(), nil, nil, err
	}
	accessToken, refreshToken, err := ul.protector.GenerateTokens(claims)
	if err != nil {
		return models.EmptyUser(), nil, nil, err
	}
	// Here we need to save refresh token in Redis
	if err := ul.tokens.SaveRefreshToken(ctx, foundUser.ID, *refreshToken); err != nil {
		return models.EmptyUser(), nil, nil, err
	}
	return foundUser, accessToken, refreshToken, nil
}

func (ul *userLogic) OAuth(ctx context.Context, user *models.User) (models.User, error) {
	return models.EmptyUser(), nil
}
