// type assertion benchmark test
package main

import (
	"fmt"
	"testing"
)

var (
	fmap map[string]func(string, string, string, string) string
	rmap map[string]func(interface{}, interface{}, interface{}, interface{}) string
)

func Init() {

	testFunct := func(a, b, c, d string) string {
		var result string
		for i := 0; i < 100; i++ {
			as, bs, cs, ds := a, b, c, d
			result += as + bs + cs + ds

		}
		return result
	}
	testFunct1 := func(a, b, c, d interface{}) string {
		var result string
		for i := 0; i < 100; i++ {
			as, bs, cs, ds := a.(string), b.(string), c.(string), d.(string)
			result += as + bs + cs + ds
		}
		return result
	}

	fmap = make(map[string]func(string, string, string, string) string)
	rmap = make(map[string]func(interface{}, interface{}, interface{}, interface{}) string)
	fmap["f1"] = testFunct
	rmap["f1"] = testFunct1

}

func main() {
	var pr float32

	Init()
	t1 := testing.B{}
	t1.N = 10

	bf := testing.Benchmark(BenchmarkF)
	br := testing.Benchmark(BenchmarkR)
	pr = float32(br.NsPerOp()) / float32(bf.NsPerOp()) * 100
	fmt.Print("originu", bf, "\n", "type assert:", br, "\n", pr, "%", "\n")
}

func BenchmarkF(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fmap["f1"]("n", "m", "n", "m")
	}
}

func BenchmarkR(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rmap["f1"]("n", "m", "n", "m")
	}
}
