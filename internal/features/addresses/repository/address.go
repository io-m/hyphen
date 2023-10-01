package address_repository

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/io-m/hyphen/internal/shared/models"
)

type addressRepository struct {
	postgres *sql.DB
	redis    *redis.Client
}

func NewAddressRepository(postgres *sql.DB, redis *redis.Client) *addressRepository {
	return &addressRepository{
		postgres: postgres,
		redis:    redis,
	}
}

type IAddressRepository interface {
	FindAllAddresses(ctx context.Context) ([]models.Address, error)
	FindAddressById(ctx context.Context, addressId int) (models.Address, error)
	FindAddressByEmail(ctx context.Context, email string) (models.Address, error)
	CreateAddress(ctx context.Context, Address *models.AddressRequest) (int, error)
	UpdateAddressById(ctx context.Context, addressId int, addressRequest *models.AddressRequest) (models.Address, error)
	DeleteAddressById(ctx context.Context, addressId int) (bool, error)
}
