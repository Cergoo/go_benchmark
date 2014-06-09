package main

import (
	"fmt"
	"strconv"
	"testing"
)

var (
	m map[string]int
	c []int
)

func init() {
	m = make(map[string]int, 100)
	c = make([]int, 0, 200)
	for i := 0; i < 200; i++ {
		c = append(c, i)
		m[strconv.Itoa(i)] = i
	}
}

func main() {
	operation_name := "range"
	r1 := testing.Benchmark(Benchmark_slice)
	r2 := testing.Benchmark(Benchmark_map)
	fmt.Print(operation_name, "\n", "slice:", r1, r1.MemString(), "\n", "map:", r2, r2.MemString(), "\n")

}

func Benchmark_slice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, i := range c {
			_ = i
		}
	}
}

func Benchmark_map(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, i := range m {
			_ = i
		}
	}
}
