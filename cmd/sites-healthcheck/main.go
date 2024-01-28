package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"google.com",
		"ya.ru",
		"go.dev",
		"github.com",
	}
	statuses := make([]int, len(urls))
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		statuses = parallelCheck(urls, &wg)
	}()
	wg.Wait()
	for _, status := range statuses {
		fmt.Println(status)
	}
}
func parallelCheck(urls []string, wg *sync.WaitGroup) []int {
	urls_count := len(urls)
	statuses := make([]int, urls_count)
	for i, url := range urls {
		wg.Add(1)
		go func(url string, index int) {
			defer wg.Done()
			statuses[index] = healthCheck(url)
		}(url, i)
	}
	return statuses
}

func healthCheck(url string) int {
	time.Sleep(2 * time.Second)
	return http.StatusCreated
}
