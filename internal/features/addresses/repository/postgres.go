package address_repository

import (
	"context"
	"log/slog"

	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/io-m/hyphen/pkg/helpers"
)

func (ar *addressRepository) FindAllAddresses(ctx context.Context) ([]models.Address, error) {
	return []models.Address{models.EmptyAddress()}, nil
}
func (ar *addressRepository) FindAddressById(ctx context.Context, addressId int) (models.Address, error) {
	return models.EmptyAddress(), nil
}
func (ar *addressRepository) FindAddressByEmail(ctx context.Context, email string) (models.Address, error) {
	return models.EmptyAddress(), nil
}
func (ar *addressRepository) CreateAddress(ctx context.Context, addressRequest *models.AddressRequest) (int, error) {
	var id int
	baseTable := "addresses"
	requiredFields := map[string]any{
		"street":    addressRequest.Street,
		"city":      addressRequest.City,
		"post_code": addressRequest.PostCode,
		"country":   addressRequest.Country,
	}
	optionalFields := map[string]any{
		"state":     addressRequest.State,
		"latitude":  addressRequest.Latitude,
		"longitude": addressRequest.Longitude,
	}

	finalStatement, values, err := helpers.ConstructInsertStatement(baseTable, requiredFields, optionalFields)

	if err != nil {
		slog.Debug("Failed to insert address: %v", err)
		return -1, err
	}
	err = ar.postgres.QueryRow(finalStatement, values...).Scan(&id)
	if err != nil {
		slog.Debug("Failed to insert address: %v", err)
		return -1, err
	}

	return id, nil
}
func (ar *addressRepository) UpdateAddressById(ctx context.Context, addressId int, addressRequest *models.AddressRequest) (models.Address, error) {
	return models.EmptyAddress(), nil
}
func (ar *addressRepository) DeleteAddressById(ctx context.Context, addressId int) (bool, error) {
	return false, nil
}
