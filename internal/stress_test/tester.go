package stress_test

import (
	"net/http"
	"sync"
)

type Result struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusCounts       map[int]int
	mu                 sync.Mutex
}

func RunLoadTest(url string, totalRequests int, concurrency int) *Result {
	results := &Result{
		TotalRequests: totalRequests,
		StatusCounts:  make(map[int]int),
	}

	semaphore := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for range totalRequests {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			resp, err := http.Get(url)
			if err == nil {
				results.mu.Lock()
				results.StatusCounts[resp.StatusCode]++

				if resp.StatusCode == http.StatusOK {
					results.SuccessfulRequests++
				}

				results.mu.Unlock()
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	return results
}
