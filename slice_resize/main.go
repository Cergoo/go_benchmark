package main

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	v1, v2 []string
)

func main() {
	v1 = make([]string, 1000)
	v2 = make([]string, 1000)

	bnative := testing.Benchmark(BenchmarkNative)
	breflect := testing.Benchmark(BenchmarkReflect)

	fmt.Print("bnative:", bnative, "\n", "breflect:", breflect, "\n")
}

func SliceResize(pointToSlice interface{}, newCap int) {
	slice := reflect.ValueOf(pointToSlice).Elem()
	newslice := reflect.MakeSlice(slice.Type(), newCap, newCap)
	reflect.Copy(newslice, slice)
	slice.Set(newslice)
}

func nativeResize(slice *[]string, newCap int) {
	newslice := make([]string, newCap)
	copy(newslice, *slice)
	*slice = newslice
}

func BenchmarkReflect(b *testing.B) {
	var newcap int
	for n := 0; n < b.N; n++ {
		if cap(v1) > 0 {
			newcap = cap(v1) - 1
		}
		SliceResize(&v1, newcap)
	}
}

func BenchmarkNative(b *testing.B) {
	var newcap int
	for n := 0; n < b.N; n++ {
		if cap(v2) > 0 {
			newcap = cap(v2) - 1
		}
		nativeResize(&v2, newcap)
	}
}
