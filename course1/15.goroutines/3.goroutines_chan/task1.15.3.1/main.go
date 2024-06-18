package main

import (
	"fmt"

	"math/rand"

	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []*Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTimes chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count)
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)

	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)
	wg.Wait()
	sinceProgramStarted = time.Since(programStart)

	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()

	checkTimeDifference(limitCount)
}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)

	for _, order := range orders {
		wg.Add(1)
		limit <- struct{}{}
		go func(o *Order) {
			defer func() {
				<-limit
				wg.Done()
			}()

			t := time.Now()
			OrderProcessing(o)
			processTimes <- time.Since(t)
		}(order)
	}
}

func OrderProcessing(o *Order) {
	if completeOrders[o.ID] {
		o.Complete = true
		fmt.Printf("Order %d processed\n", o.ID)
	} else {
		fmt.Printf("Order %d is not complete\n", o.ID)
	}
}

func GenerateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:       i + 1,
			Complete: false,
		}
	}
	return orders
}

func GenerateCompleteOrders(count int) map[int]bool {
	rand.Seed(time.Now().UnixNano())
	completeOrders := make(map[int]bool)
	for i := 0; i < count; i++ {
		if rand.Float32() < 0.5 {
			completeOrders[i+1] = true
		}
	}
	return completeOrders
}

func checkTimeDifference(limitCount int) {
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int

	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}

	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}

	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	fmt.Printf("Order process total time: %v\n", orderProcessTotalTime/time.Second)
	fmt.Printf("Average time per order: %v\n", averageTime/time.Second)
	fmt.Printf("Time since program started: %v\n", sinceProgramStarted/time.Second)
	fmt.Printf("Time difference per order: %v\n", (orderProcessTotalTime-sinceProgramStarted)/time.Second)

	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}
}
