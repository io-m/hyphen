package address_logic

import (
	"context"

	address_repository "github.com/io-m/hyphen/internal/features/addresses/repository"
	"github.com/io-m/hyphen/internal/shared"
	"github.com/io-m/hyphen/internal/shared/models"
)

type IAddressLogic interface {
	GetAddressWithId(ctx context.Context, addressId uint) (models.Address, error)
	GenerateAddress(ctx context.Context, address *models.AddressRequest) (shared.ItemID, error)
	UpdateAddressWithId(ctx context.Context, addressId uint, addressRequestOptional *models.AddressRequest) (models.Address, error)
	DeleteAddressWithId(ctx context.Context, addressId uint) (bool, error)
}

type addressLogic struct {
	addressRepository address_repository.IAddressRepository
}

func NewAddressLogic(addressRepository address_repository.IAddressRepository) *addressLogic {
	return &addressLogic{
		addressRepository: addressRepository,
	}
}

func (al *addressLogic) GetAddressWithId(ctx context.Context, addressId uint) (models.Address, error) {
	return models.EmptyAddress(), nil
}
func (al *addressLogic) GenerateAddress(ctx context.Context, address *models.AddressRequest) (shared.ItemID, error) {
	addressItemId, err := al.addressRepository.CreateAddress(ctx, address)
	if err != nil {
		return shared.ItemID{ID: addressItemId}, err
	}
	return shared.ItemID{ID: addressItemId}, nil
}
func (al *addressLogic) UpdateAddressWithId(ctx context.Context, addressId uint, addressRequestOptional *models.AddressRequest) (models.Address, error) {
	return models.EmptyAddress(), nil
}
func (al *addressLogic) DeleteAddressWithId(ctx context.Context, addressId uint) (bool, error) {
	return false, nil
}
