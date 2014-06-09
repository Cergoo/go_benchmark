package main

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	obj reflect.Value
	m   map[int]reflect.Value
	sl  []reflect.Value
)

func init() {
	var n interface{}
	sl = []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3), reflect.ValueOf(4), reflect.ValueOf(5)}
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
	bslice := testing.Benchmark(BenchmarkSlice)

	fmt.Println("Get element benchmark:")
	fmt.Print("native_to_map:", bnative, "\n", "reflect_to_slice:", breflect, "\n", "native_to_slice:", bslice)
}

func BenchmarkReflect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 5; i++ {
			_ = obj.Index(i)
		}
	}

}

func BenchmarkNative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 5; i++ {
			_ = m[i]
		}
	}
}

func BenchmarkSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 5; i++ {
			_ = sl[i]
		}
	}
}
