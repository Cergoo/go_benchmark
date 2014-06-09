package main

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	chan1 chan interface{}
	chan2 chan int
	chan3 reflect.Value
	m     reflect.Value
)

func init() {
	chan1 = make(chan interface{}, 1)
	chan2 = make(chan int, 1)
	chan3 = reflect.ValueOf(chan2)
	m = reflect.ValueOf(10)
}

func main() {
	operation_name := "work"
	r1 := testing.Benchmark(Benchmark_chint)
	r2 := testing.Benchmark(Benchmark_chref)
	fmt.Print(operation_name, "\n", "ch int:", r1, r1.MemString(), "\n", "ch ref:", r2, r2.MemString(), "\n")

}

func Benchmark_chint(b *testing.B) {
	for n := 0; n < b.N; n++ {
		chan1 <- n
		_ = (<-chan1).(int)
	}
}

func Benchmark_chref(b *testing.B) {
	for n := 0; n < b.N; n++ {
		chan3.TrySend(m)
		_, _ = chan3.Recv()
	}
}
