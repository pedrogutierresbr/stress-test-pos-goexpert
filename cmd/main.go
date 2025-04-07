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
	requests := flag.Int("requests", 100, "Número total de requests (mínimo 1)")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas (mínimo 1)")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Uso: %s [flags]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: a URL deve ser fornecida")
		flag.Usage()
		os.Exit(1)
	}
	if *requests < 1 {
		fmt.Println("Erro: o número de requests deve ser maior que 0")
		flag.Usage()
		os.Exit(1)
	}
	if *concurrency < 1 {
		fmt.Println("Erro: o número de chamadas simultâneas deve ser maior que 0")
		flag.Usage()
		os.Exit(1)
	}

	startTime := time.Now()
	results := stress_test.RunLoadTest(*url, *requests, *concurrency)
	duration := time.Since(startTime)
	stress_test.GenerateReport(results, duration)
}
