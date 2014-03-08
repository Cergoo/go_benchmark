package main

import (
	"fmt"
	"sync"
	"testing"
)

var (
	wg     sync.WaitGroup
	ch1    chan []byte
	r1, r2 testing.BenchmarkResult
	p      = []byte("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqq")
)

func f1(n []byte) {
	wg.Done()
}

func f2() {
	for {
		<-ch1
	}
}

func main() {
	ch1 = make(chan []byte)
	go f2()
	operation_name := "work"
	r1 = testing.Benchmark(Benchmark_go)
	r2 = testing.Benchmark(Benchmark_ch)
	close(ch1)
	fmt.Print(operation_name, "\n", "go:", r1, r1.MemString(), "\n", "chan:", r2, r2.MemString(), "\n")

}

func Benchmark_go(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go f1(p)
	}
}

func Benchmark_ch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ch1 <- p
	}
}
