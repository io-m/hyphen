package profile_repository

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/io-m/hyphen/internal/shared/models"
)

//go:embed queries/create_profile.psql
var CREATE_PROFILE string

func (pr *profileRepository) FindAllProfiles(ctx context.Context) ([]models.Profile, error) {
	return []models.Profile{models.EmptyProfile()}, nil
}
func (pr *profileRepository) FindProfileById(ctx context.Context, profileId uint) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (pr *profileRepository) FindProfileByEmail(ctx context.Context, email string) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (pr *profileRepository) CreateProfile(ctx context.Context, profile *models.Profile) (int64, error) {
	var id int64
	err := pr.postgres.QueryRow(CREATE_PROFILE, profile.Email, profile.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not insert profile and retrieve ID: %w", err)
	}
	if err != nil {
		return 0, fmt.Errorf("could not insert profile: %w", err)
	}

	return id, nil
}
func (pr *profileRepository) UpdateProfileById(ctx context.Context, profileId uint, profileRequest *models.ProfileUpdateRequest) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (r *profileRepository) DeleteProfileById(ctx context.Context, profileId uint) (bool, error) {
	return false, nil
}
