package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/pedrogutierresbr/stress-test-pos-goexpert/internal/stress_test"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()
	if *url == "" {
		fmt.Println("A URL deve ser fornecida")
		os.Exit(1)
	}
	startTime := time.Now()
	results := stress_test.RunLoadTest(*url, *requests, *concurrency)
	duration := time.Since(startTime)
	stress_test.GenerateReport(results, duration)
}
