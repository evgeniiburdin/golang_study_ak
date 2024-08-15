package usecase

import (
	"context"

	"geo-service/internal/entity"
)

type (
	// Addresser -.
	Addresser interface {
		GeocodeToAddress(context.Context, entity.Geocode) (*entity.Address, error)
	}

	// AddressWebAPI -.
	AddressWebAPI interface {
		GeocodeToAddress(entity.Geocode) (*entity.Address, error)
	}
)
