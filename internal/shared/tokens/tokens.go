package tokens

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/io-m/hyphen/pkg/constants"
)

type AuthorizationLevel string

const (
	CUSTOMER AuthorizationLevel = "CUSTOMER"
	PUBLIC   AuthorizationLevel = "PUBLIC"
	PROVIDER AuthorizationLevel = "PROVIDER"
	SUDO     AuthorizationLevel = "SUDO"
)

type ITokens interface {
	SaveRefreshToken(ctx context.Context, profileId int, refreshToken string) error
	DeleteRefreshToken(ctx context.Context, profileId int, refreshToken string) error
	RetrieveRefreshToken(ctx context.Context, profileId int, refreshToken string) (string, error)
}

type tokens struct {
	redis *redis.Client
}

func NewTokens(redis *redis.Client) *tokens {
	return &tokens{
		redis: redis,
	}
}

func (t *tokens) SaveRefreshToken(ctx context.Context, profileId int, refreshToken string) error {
	err := t.redis.HSet(ctx, fmt.Sprintf("%d", profileId), "refreshToken", refreshToken)

	if err != nil {
		return err.Err()
	}
	return nil
}

func (t *tokens) RetrieveRefreshToken(ctx context.Context, profileId int, refreshToken string) (string, error) {
	rt, err := t.redis.Get(ctx, constants.REFRESH_TOKEN_KEY).Result()
	if err != redis.Nil {
		return "", fmt.Errorf("%s key does not exist: %w", constants.REFRESH_TOKEN_KEY, err)
	} else if err != nil {
		return "", err
	}
	return rt, nil
}

func (t *tokens) DeleteRefreshToken(ctx context.Context, profileId int, refreshToken string) error {
	return nil
}
