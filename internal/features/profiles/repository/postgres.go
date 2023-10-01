package profile_repository

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"

	"github.com/io-m/hyphen/internal/shared/models"
	"golang.org/x/exp/slog"
)

//go:embed queries/create_profile.psql
var CREATE_PROFILE string

//go:embed queries/find_profile_by_email.psql
var FIND_PROFILE_BY_EMAIL string

func (pr *profileRepository) FindAllProfiles(ctx context.Context) ([]models.Profile, error) {
	return []models.Profile{models.EmptyProfile()}, nil
}
func (pr *profileRepository) FindProfileById(ctx context.Context, profileId int) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (pr *profileRepository) FindProfileByEmail(ctx context.Context, email string) (models.Profile, error) {
	var profile models.Profile
	err := pr.postgres.QueryRowContext(ctx, FIND_PROFILE_BY_EMAIL, email).Scan(&profile.ID, &profile.Email, &profile.Password, &profile.CreatedAt, &profile.UpdatedAt)

	// Handling the case where no rows are found
	if err == sql.ErrNoRows {
		return models.Profile{}, errors.New("no profile found with the given email")
	}

	if err != nil {
		slog.Debug("Error querying profile by email: %v", err)
		return models.Profile{}, fmt.Errorf("could not query profile: %w", err)
	}
	return profile, nil
}
func (pr *profileRepository) CreateProfile(ctx context.Context, profile *models.Profile) (int, error) {
	var id int
	err := pr.postgres.QueryRowContext(ctx, CREATE_PROFILE, profile.Email, profile.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not insert profile and retrieve ID: %w", err)
	}
	if err != nil {
		return 0, fmt.Errorf("could not insert profile: %w", err)
	}

	return id, nil
}
func (pr *profileRepository) UpdateProfileById(ctx context.Context, profileId int, profileRequest *models.ProfileUpdateRequest) (models.Profile, error) {
	return models.EmptyProfile(), nil
}
func (r *profileRepository) DeleteProfileById(ctx context.Context, profileId int) (bool, error) {
	return false, nil
}
