package main

import (
	"fmt"
	"geo-service/internal/entity"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatalf("Error dialing rpc at :1234 - %v", err)
	}

	args := entity.Geocode{
		Lat: "47.6062",
		Lng: "-122.3321",
	}
	var reply entity.Address

	err = client.Call("JSONRPCServer.GeocodeToAddress", args, &reply)
	if err != nil {
		log.Fatalf("JSONRPC error - %v", err)
	}

	fmt.Printf("%#v", reply)
}
