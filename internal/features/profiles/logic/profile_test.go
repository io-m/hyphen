package profile_logic

import (
	"context"
	"testing"

	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/models"
)

func TestRegisterProfile(t *testing.T) {
	// Create a mock profile and context
	mockProfile := &models.Profile{
		Email:    "test@example.com",
		Password: "passworD123!",
	}
	mockContext := context.TODO()

	// Create a mock profile repository with predefined behavior
	mockRepository := profile_repository.NewInMemoryProfileRepository()

	// Create a profile logic instance with the mock repository
	profileLogic := NewProfileLogic(mockRepository, nil, nil)

	// Call the RegisterProfile function
	createdProfileID, err := profileLogic.RegisterProfile(mockContext, mockProfile)

	// Perform assertions based on expected behavior
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if createdProfileID.ID != len(profile_repository.DummyProfiles) {
		t.Errorf("Expected createdProfileID to be %d, got %v", len(profile_repository.DummyProfiles)+1, createdProfileID.ID)
	}
}
