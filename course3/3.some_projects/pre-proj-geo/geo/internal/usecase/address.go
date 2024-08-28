package usecase

import (
	"context"
	"fmt"

	"geo-service/internal/entity"
)

// AddressUseCase -.
type AddressUseCase struct {
	webAPI AddressWebAPI
}

// New -.
func New(w AddressWebAPI) *AddressUseCase {
	return &AddressUseCase{
		webAPI: w,
	}
}

// GeocodeToAddress - method to get address info based on provided geo coordinates
func (uc AddressUseCase) GeocodeToAddress(ctx context.Context, r entity.Geocode) (*entity.Address, error) {
	geocode, err := uc.webAPI.GeocodeToAddress(r)
	if err != nil {
		return &entity.Address{}, fmt.Errorf("AddressUseCase - GeocodeToAddress: %w", err)
	}

	return geocode, nil
}
