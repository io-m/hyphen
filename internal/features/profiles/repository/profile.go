package profile_repository

import (
	"context"

	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/jmoiron/sqlx"
)

type postgresProfileRepository struct {
	postgres *sqlx.DB
	// redis    *redis.Client
}

type IProfileRepository interface {
	FindAllProfiles(ctx context.Context) ([]models.Profile, error)
	FindProfileById(ctx context.Context, profileId uint) (models.Profile, error)
	FindProfileByEmail(ctx context.Context, email string) (models.Profile, error)
	CreateProfile(ctx context.Context, profile *models.Profile) (models.Profile, error)
	UpdateProfileById(ctx context.Context, profileId uint, profileRequest *models.ProfileUpdateRequest) (models.Profile, error)
	DeleteProfileById(ctx context.Context, profileId uint) (bool, error)
}
