package main

import (
	"fmt"
	"geo-service/internal/entity"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("error dialing rpc at :1234 - %v", err)
	}

	args := entity.Geocode{
		Lat: "47.6062",
		Lng: "-122.3321",
	}
	var reply entity.Address

	err = client.Call("RPCServer.GeocodeToAddress", args, &reply)
	if err != nil {
		log.Fatalf("rpc error - %v", err)
	}

	fmt.Printf("%#v\n", reply)
}
