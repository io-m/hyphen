package profile_repository

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/io-m/hyphen/internal/shared/models"
)

type inMemoryProfileRepository struct{}

func NewInMemoryProfileRepository() *inMemoryProfileRepository {
	return &inMemoryProfileRepository{}
}

var DummyProfiles = []models.Profile{
	{
		ID:        1,
		Email:     "user1@example.com",
		Password:  "passworD123!",
		CreatedAt: CreateTime(2023, 9, 28, 10, 0, 0),
		UpdatedAt: CreateTime(2023, 9, 28, 10, 0, 0),
	},
	{
		ID:        2,
		Email:     "user2@example.com",
		Password:  "passworD223!",
		CreatedAt: CreateTime(2023, 9, 27, 14, 30, 0),
		UpdatedAt: CreateTime(2023, 9, 28, 11, 15, 0),
	},
	{
		ID:        3,
		Email:     "user3@example.com",
		Password:  "passworD323!",
		CreatedAt: CreateTime(2023, 9, 26, 9, 45, 0),
		UpdatedAt: CreateTime(2023, 9, 27, 16, 20, 0),
	},
	{
		ID:        4,
		Email:     "user4@example.com",
		Password:  "passworD423!",
		CreatedAt: CreateTime(2023, 9, 25, 12, 15, 0),
		UpdatedAt: CreateTime(2023, 9, 26, 13, 40, 0),
	},
	{
		ID:        5,
		Email:     "user5@example.com",
		Password:  "passworD523!",
		CreatedAt: CreateTime(2023, 9, 24, 16, 50, 0),
		UpdatedAt: CreateTime(2023, 9, 25, 10, 5, 0),
	},
}

func CreateTime(year, month, day, hour, minute, second int) *time.Time {
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	return &t
}

func (*inMemoryProfileRepository) FindAllProfiles(ctx context.Context) ([]models.Profile, error) {
	return DummyProfiles, nil
}

func (*inMemoryProfileRepository) FindProfileById(ctx context.Context, profileId int) (models.Profile, error) {
	for _, profile := range DummyProfiles {
		if profile.ID == profileId {
			return profile, nil
		}
	}
	return models.Profile{}, errors.New("profile not found")
}

func (*inMemoryProfileRepository) FindProfileByEmail(ctx context.Context, email string) (models.Profile, error) {
	for _, profile := range DummyProfiles {
		if profile.Email == email {
			return profile, nil
		}
	}
	return models.Profile{}, errors.New("profile not found")
}

func (*inMemoryProfileRepository) CreateProfile(ctx context.Context, profile *models.Profile) (int, error) {
	// Generate a new ID (you may use a different method to generate unique IDs)
	newID := len(DummyProfiles) + 1
	profile.ID = newID
	DummyProfiles = append(DummyProfiles, *profile)
	slog.Debug("NED DUMMY LIST --->		", DummyProfiles)
	return newID, nil
}

func (*inMemoryProfileRepository) UpdateProfileById(ctx context.Context, profileId int, profileRequest *models.ProfileUpdateRequest) (models.Profile, error) {
	now := time.Now()
	for i, profile := range DummyProfiles {
		if profile.ID == profileId {
			// Update the profile fields if the corresponding fields in the request are not nil
			if profileRequest.Email != nil {
				profile.Email = *profileRequest.Email
			}
			if profileRequest.NewPassword != nil {
				profile.Password = *profileRequest.NewPassword
			}
			profile.UpdatedAt = &now
			DummyProfiles[i] = profile
			return profile, nil
		}
	}
	return models.Profile{}, errors.New("profile not found")
}

func (*inMemoryProfileRepository) DeleteProfileById(ctx context.Context, profileId int) (bool, error) {
	for i, profile := range DummyProfiles {
		if profile.ID == profileId {
			// Remove the profile from the DummyProfiles slice
			DummyProfiles = append(DummyProfiles[:i], DummyProfiles[i+1:]...)
			return true, nil
		}
	}
	return false, errors.New("profile not found")
}
