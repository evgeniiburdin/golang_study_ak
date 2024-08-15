package webapi

import (
	"context"
	"fmt"

	"github.com/alexliesenfeld/opencage"

	"geo-service/internal/entity"
)

// AddressWebAPI -.
type AddressWebAPI struct {
	client *opencage.Client
}

// New -.
func New(APIKey string) AddressWebAPI {
	return AddressWebAPI{
		client: opencage.New(APIKey),
	}
}

// GeocodeToAddress -.
func (w AddressWebAPI) GeocodeToAddress(geocode entity.Geocode) (*entity.Address, error) {
	webAPIResp, err := w.client.Geocode(context.Background(), geocode.Lat+","+geocode.Lng, &opencage.GeocodingParams{})
	if err != nil {
		return nil, fmt.Errorf("AddressWebAPI - GeocodeToAddress: %w", err)
	}

	resp := &entity.Address{
		Country: webAPIResp.Results[0].Components.Country,
		City:    webAPIResp.Results[0].Components.City,
	}

	return resp, nil
}
