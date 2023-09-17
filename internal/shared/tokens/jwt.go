package tokens

import (
	"context"

	"github.com/google/uuid"
	"github.com/io-m/hyphen/pkg/constants"
)

type jwtAuthenticator struct {
	accessTokenSecretKey  []byte
	refreshTokenSecretKey []byte
}

func NewJWTProtector() IProtector {
	return &jwtAuthenticator{
		accessTokenSecretKey:  []byte(constants.ACCESS_TOKEN_SECRET_KEY),
		refreshTokenSecretKey: []byte(constants.REFRESH_TOKEN_SECRET_KEY),
	}
}

// TODO: Implement NewJWTProtector to be IAuthenticator
func (protector *jwtAuthenticator) GenerateTokens(claims *Claims) (*string, *string, error) {

	return nil, nil, nil
}

// TODO: Implement NewJWTProtector to be IAuthenticator
func (protector *jwtAuthenticator) VerifyToken(stringifiedToken string) (*Claims, error) {

	return nil, nil
}

func (protector *jwtAuthenticator) VerifyRefreshToken(ctx context.Context, customerId uuid.UUID, refreshToken string) (bool, error) {
	return false, nil
}
