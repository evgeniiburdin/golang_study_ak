package main

import (
	"context"
	"geo-service/api/geo/grpc/gen/geo-service/geo"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestGeocodeToAddress(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err)
	defer conn.Close()

	client := geo.NewGeoServiceClient(conn)
	resp, err := client.GeocodeToAddress(ctx, &geo.Geocode{Lat: "47.6062", Lng: "-122.3321"})
	assert.NoError(t, err)
	assert.Equal(t, "USA", resp.Country)
	assert.Equal(t, "Seattle", resp.City)
}
