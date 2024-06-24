package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://httpbin.org/get"
	requestsAtOneTimeLimit := 5
	requestCount := 50

	result := benchRequest(url, requestsAtOneTimeLimit, requestCount)

	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}

	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, requestsAtOneTimeLimit int, requestCount int) <-chan int {
	limit := make(chan struct{}, requestsAtOneTimeLimit)
	results := make(chan int, requestCount)

	for i := 0; i < requestCount; i++ {
		limit <- struct{}{}
		go func(url string) {
			statusCode, err := httpRequest(url)
			if err != nil {
				log.Printf("Error performing http-get request to %s: %v", url, err.Error())
			}
			<-limit
			results <- statusCode
		}(url)
	}

	return results
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
