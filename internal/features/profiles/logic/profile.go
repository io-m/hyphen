package profile_logic

import (
	"context"
	"errors"
	"fmt"

	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/io-m/hyphen/internal/shared/tokens"
	"github.com/io-m/hyphen/pkg/constants"
	"github.com/io-m/hyphen/pkg/helpers"
)

type IProfileLogic interface {
	GetProfileWithEmail(ctx context.Context, ProfileId uint) (models.Profile, error)
	GetProfileWithId(ctx context.Context, ProfileId uint) (models.Profile, error)
	RegisterProfile(ctx context.Context, Profile *models.Profile) (models.Profile, error)
	Login(ctx context.Context, Profile *models.Profile) (models.Profile, *string, *string, error)
	OAuth(ctx context.Context, Profile *models.Profile) (models.Profile, error)
	UpdateProfileWithId(ctx context.Context, ProfileId uint, ProfileRequestOptional *models.ProfileUpdateRequest) (models.Profile, error)
	DeleteProfileWithId(ctx context.Context, ProfileId uint) (bool, error)
}

type profileLogic struct {
	profileRepository profile_repository.IProfileRepository
	tokens            tokens.ITokens
	protector         tokens.IProtector
}

func NewProfileLogic(profileRepository profile_repository.IProfileRepository, tokens tokens.ITokens, protector tokens.IProtector) *profileLogic {
	return &profileLogic{
		profileRepository: profileRepository,
		tokens:            tokens,
		protector:         protector,
	}
}

func (ul *profileLogic) GetProfileWithEmail(ctx context.Context, profileId uint) (models.Profile, error) {
	return models.EmptyProfile(), nil
}

func (ul *profileLogic) GetProfileWithId(ctx context.Context, profileId uint) (models.Profile, error) {
	return models.EmptyProfile(), nil
}

func (ul *profileLogic) UpdateProfileWithId(ctx context.Context, profileId uint, profileRequestOptional *models.ProfileUpdateRequest) (models.Profile, error) {
	return models.EmptyProfile(), nil
}

func (ul *profileLogic) DeleteProfileWithId(ctx context.Context, profileId uint) (bool, error) {
	return false, nil
}

func (ul *profileLogic) RegisterProfile(ctx context.Context, profile *models.Profile) (models.Profile, error) {
	if err := helpers.ValidatePassword(profile.Password); err != nil {
		return models.EmptyProfile(), errors.New("password is invalid")
	}
	if err := helpers.ValidateEmail(profile.Email); err != nil {
		return models.EmptyProfile(), errors.New("email is invalid")
	}
	hashedPassword, err := helpers.HashPassword(profile.Password)
	if err != nil {
		return models.EmptyProfile(), fmt.Errorf("error while hashing password: %w", err)
	}
	profile.Password = hashedPassword
	createdPerson, err := ul.profileRepository.CreateProfile(ctx, profile)

	if err != nil {
		return models.EmptyProfile(), err
	}
	return createdPerson, nil
}

func (ul *profileLogic) Login(ctx context.Context, profile *models.Profile) (models.Profile, *string, *string, error) {
	foundprofile, err := ul.profileRepository.FindProfileByEmail(ctx, profile.Email)
	if err != nil {
		return models.EmptyProfile(), nil, nil, err
	}
	if err := helpers.CheckPassword(profile.Password, foundprofile.Password); err != nil {
		return models.EmptyProfile(), nil, nil, fmt.Errorf("wrong password: %w", err)
	}
	claims, err := tokens.NewClaims(profile.ID, constants.ACCESS_TOKEN_DURATION)
	if err != nil {
		return models.EmptyProfile(), nil, nil, err
	}
	accessToken, refreshToken, err := ul.protector.GenerateTokens(claims)
	if err != nil {
		return models.EmptyProfile(), nil, nil, err
	}
	// Here we need to save refresh token in Redis
	if err := ul.tokens.SaveRefreshToken(ctx, foundprofile.ID, *refreshToken); err != nil {
		return models.EmptyProfile(), nil, nil, err
	}
	return foundprofile, accessToken, refreshToken, nil
}

func (ul *profileLogic) OAuth(ctx context.Context, profile *models.Profile) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
