package profile_repository

import (
	"context"

	"github.com/io-m/hyphen/internal/shared/models"
)

func (pr *postgresProfileRepository) FindAllProfiles(ctx context.Context) ([]models.Profile, error) {
	return []models.Profile{models.EmptyProfile()}, nil
}
func (pr *postgresProfileRepository) FindProfileById(ctx context.Context, profileId uint) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (pr *postgresProfileRepository) FindProfileByEmail(ctx context.Context, email string) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (pr *postgresProfileRepository) CreateProfile(ctx context.Context, profile *models.Profile) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (pr *postgresProfileRepository) UpdateProfileById(ctx context.Context, profileId uint, profileRequest *models.ProfileUpdateRequest) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (r *postgresProfileRepository) DeleteProfileById(ctx context.Context, profileId uint) (bool, error) {
	return false, nil
}
