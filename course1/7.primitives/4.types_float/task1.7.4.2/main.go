package main

import (
	"fmt"
	"github.com/mattevans/dinero"
	"time"
)

func main() {
	fmt.Println(currencyPairRate("USD", "EUR", 100.0))
}

func currencyPairRate(from string, to string, amount float64) float64 {
	client := dinero.NewClient("c91c02817bc54e62b1fa61117e0e3c12", from, 20*time.Minute)
	rsp, err := client.Rates.Get(to)
	if err != nil {
		fmt.Println(fmt.Errorf("get exchange rates: %v", err))
	}
	return *rsp
}
