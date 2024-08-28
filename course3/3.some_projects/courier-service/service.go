package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var ctx = context.Background()

type Courier struct {
	ID        string
	Latitude  float64
	Longitude float64
	Zoom      int
}

type Order struct {
	ID        string
	Latitude  float64
	Longitude float64
}

const (
	MoveStep = 0.001  // Step size for each move, can vary based on zoom level
	Radius   = 0.0225 // Radius for displaying orders in degrees (~2500 meters)
)

var rdb *redis.Client

func initRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
		DB:   0,
	})

	// Example: Set initial courier location
	rdb.Set(ctx, "courier:1", "55.751244,37.618423", 0)
}

func generateOrder(id string, lat, lon float64) {
	orderKey := fmt.Sprintf("order:%s", id)
	location := fmt.Sprintf("%f,%f", lat, lon)
	rdb.Set(ctx, orderKey, location, 0)
	fmt.Printf("Order %s generated at %f, %f\n", id, lat, lon)
}

func main() {
	initRedis()

	generateOrder("100", 55.751244, 37.618423)
	generateOrder("101", 55.751244, 37.618423)
	generateOrder("102", 55.751244, 37.618423)

	courier := Courier{
		ID:        "1",
		Latitude:  55.751244,
		Longitude: 37.618423,
		Zoom:      10,
	}

	for {
		var input string
		fmt.Println("Enter command (W/A/S/D):")
		fmt.Scanln(&input)

		switch input {
		case "W":
			courier.move("up")
		case "A":
			courier.move("left")
		case "S":
			courier.move("down")
		case "D":
			courier.move("right")
		default:
			fmt.Println("Invalid command!")
		}

		courier.displayOrdersInRadius()
	}
}

func (c *Courier) move(direction string) {
	step := MoveStep / float64(c.Zoom)
	switch direction {
	case "up":
		c.Latitude += step
	case "down":
		c.Latitude -= step
	case "left":
		c.Longitude -= step
	case "right":
		c.Longitude += step
	}

	location := fmt.Sprintf("%f,%f", c.Latitude, c.Longitude)
	rdb.Set(ctx, "courier:"+c.ID, location, 0)

	fmt.Printf("Courier moved to: %f, %f\n", c.Latitude, c.Longitude)
}

func (c *Courier) displayOrdersInRadius() {
	keys := rdb.Keys(ctx, "order:*").Val()
	for _, key := range keys {
		location := rdb.Get(ctx, key).Val()
		latLon := parseLocation(location)

		distance := calculateDistance(c.Latitude, c.Longitude, latLon[0], latLon[1])
		if distance <= Radius {
			fmt.Printf("Order %s is within radius: %f meters\n", key, distance*1000) // Convert to meters
		}
	}
}

func parseLocation(location string) [2]float64 {
	coords := [2]float64{}
	splitLoc := strings.Split(location, ",")
	coords[0], _ = strconv.ParseFloat(splitLoc[0], 64)
	coords[1], _ = strconv.ParseFloat(splitLoc[1], 64)
	return coords
}

func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371.0 // Earth radius in kilometers

	dLat := (lat2 - lat1) * (math.Pi / 180.0)
	dLon := (lon2 - lon1) * (math.Pi / 180.0)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat1*(math.Pi/180.0))*math.Cos(lat2*(math.Pi/180.0))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
