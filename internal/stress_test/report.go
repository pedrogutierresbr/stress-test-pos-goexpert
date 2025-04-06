package stress_test

import (
	"fmt"
	"time"
)

func GenerateReport(results *Result, duration time.Duration) {
	fmt.Printf("Tempo total: %s\n", duration)
	fmt.Printf("Total de requests realizados: %d\n", results.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", results.SuccessfulRequests)

	fmt.Println("Distribuição de outros códigos de status:")
	results.mu.Lock()
	for status, count := range results.StatusCounts {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
	results.mu.Unlock()
}
