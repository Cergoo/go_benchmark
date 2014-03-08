package main

import (
	"fmt"
	"strconv"
	"testing"
)

func main() {

	bnative := testing.Benchmark(BenchmarkNative)
	batoi := testing.Benchmark(BenchmarkAtoi)

	fmt.Print("bnative:", bnative, "\n", "batoi:", batoi, "\n")
}

func sigmastr(a, b string) int {
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)
	return ai + bi
}

func sigmaint(a, b int) int {
	return a + b
}

func BenchmarkAtoi(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sigmastr("10", "20")
	}
}

func BenchmarkNative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sigmaint(10, 20)
	}
}
