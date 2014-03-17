package main

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	obj reflect.Value
	m   map[int]reflect.Value
)

func init() {
	var n interface{}
	sl := []int{1, 2, 3, 4, 5}
	n = sl
	obj = reflect.ValueOf(n)
	m = make(map[int]reflect.Value)
	for key, val := range sl {
		n = val
		m[key] = reflect.ValueOf(n)
	}
}

func main() {
	bnative := testing.Benchmark(BenchmarkNative)
	breflect := testing.Benchmark(BenchmarkReflect)

	fmt.Print("\n", "bnative:", bnative, "\n", "breflect:", breflect, "\n")
}

func BenchmarkReflect(b *testing.B) {
	var v reflect.Value
	for n := 0; n < b.N; n++ {
		v = obj.Index(4)
	}
	fmt.Print(v)
}

func BenchmarkNative(b *testing.B) {
	var v reflect.Value
	for n := 0; n < b.N; n++ {
		v = m[4]
	}
	fmt.Print(v)
}
