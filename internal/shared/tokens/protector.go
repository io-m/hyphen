package tokens

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/io-m/hyphen/pkg/constants"
)

type IProtector interface {
	GenerateTokens(claims *Claims) (*string, *string, error)
	VerifyToken(token string) (*Claims, error)
	VerifyRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) (bool, error)
}

// Based on running environment we select protector
func NewProtector() IProtector {
	if os.Getenv(constants.RUNNING_ENV) == constants.DEVELOPMENT {
		return NewPasetoProtector()
	}
	return NewJWTProtector()
}

type Claims struct {
	ClaimID   uuid.UUID `json:"jti"`
	SubjectID int64     `json:"sub"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"exp"`
	// Roles     []entities.AuthorizationLevel `json:"roles,omitempty"`
}

func NewClaims(subjectID int64, duration time.Duration) (*Claims, error) {
	claimID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.New("could not generate claims")
	}
	claims := &Claims{
		ClaimID:   claimID,
		SubjectID: subjectID,
		IssuedAt:  time.Now().UTC(),
		ExpiredAt: time.Now().Add(constants.ACCESS_TOKEN_DURATION).UTC(),
	}
	return claims, nil
}
