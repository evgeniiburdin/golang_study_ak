package service

import (
	"context"

	geocoder "github.com/alexliesenfeld/opencage"

	"geo-service/models"
)

type GeoServicer interface {
	SearchAddress(address string) (*models.ResponseAddressInfo, error)
	GeocodeAddress(lat, lng string) (*models.ResponseAddressInfo, error)
}

type GeoService struct {
	client *geocoder.Client
}

func NewGeoService(apiKey string) *GeoService {
	return &GeoService{
		client: geocoder.New(apiKey),
	}
}

func (s *GeoService) SearchAddress(address string) (*models.ResponseAddressInfo, error) {
	response, err := s.client.Geocode(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}

	info := &models.ResponseAddressInfo{
		Info: make([]interface{}, 0),
	}
	for _, result := range response.Results {
		info.Info = append(info.Info, result)
	}
	return info, nil
}

func (s *GeoService) GeocodeAddress(lat, lng string) (*models.ResponseAddressInfo, error) {
	response, err := s.client.Geocode(context.Background(), lat+","+lng, &geocoder.GeocodingParams{
		RoadInfo: true,
	})
	if err != nil {
		return nil, err
	}

	info := &models.ResponseAddressInfo{
		Info: make([]interface{}, 0),
	}
	for _, result := range response.Results {
		info.Info = append(info.Info, result)
	}
	return info, nil
}
