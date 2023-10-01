package profile_repository

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/io-m/hyphen/internal/shared/models"
)

type profileRepository struct {
	postgres *sql.DB
	redis    *redis.Client
}

func NewProfileRepository(postgres *sql.DB, redis *redis.Client) *profileRepository {
	return &profileRepository{
		postgres: postgres,
		redis:    redis,
	}
}

type IProfileRepository interface {
	FindAllProfiles(ctx context.Context) ([]models.Profile, error)
	FindProfileById(ctx context.Context, profileId int) (models.Profile, error)
	FindProfileByEmail(ctx context.Context, email string) (models.Profile, error)
	CreateProfile(ctx context.Context, profile *models.Profile) (int, error)
	UpdateProfileById(ctx context.Context, profileId int, profileRequest *models.ProfileUpdateRequest) (models.Profile, error)
	DeleteProfileById(ctx context.Context, profileId int) (bool, error)
}
