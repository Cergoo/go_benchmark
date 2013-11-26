package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type (
	t_obj struct {
		Itm_1 int
		Itm_2 string
		Itm_3 bool
		Itm_4 float32
		Itm_5 byte
		Itm_6 []string
		Itm_7 []int
		Itm_8 []bool
		Itm_9 []float32
		Itm10 map[string]string
		Itm11 map[string]int
		Itm12 map[string]bool
		Itm13 *t_obj
	}
)

var (
	v, v1 t_obj
	rjson []byte
	e     error
)

func init() {
	v = t_obj{}
	v.Itm_1 = 100
	v.Itm_2 = "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb"
	v.Itm_3 = true
	v.Itm_4 = 17.257
	v.Itm_5 = 117
	v.Itm_6 = []string{
		"aaaaaaaaaaaaaaaabbbbbbbbbbbbbb1", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb2", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb3", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb4", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb5",
		"aaaaaaaaaaaaaaaabbbbbbbbbbbbbb6", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb7", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb8", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb9", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbb10",
	}
	v.Itm_7 = []int{1, 2, 4, 7, 8, 9, 11, 145, 18, 20}
	v.Itm_8 = []bool{true, false, true, false, true, false, true, false, true, false}
	v.Itm_9 = []float32{11.2, 12.1, 17.1, 11.1, 14.21}
	v.Itm10 = map[string]string{
		"key1": "itm1", "key2": "itm2", "key3": "itm3", "key4": "itm4", "key5": "itm5",
		"key6": "itm6", "key7": "itm7", "key8": "itm8", "key9": "itm9", "key10": "itm10",
	}
	v.Itm11 = map[string]int{
		"key1": 1, "key2": 2, "key3": 10, "key4": 12, "key5": 12,
		"key6": 14, "key7": 12, "key8": 12, "key9": 12, "key10": 14,
	}
	v.Itm12 = map[string]bool{
		"key1": true, "key2": false, "key3": true, "key4": false, "key5": true,
		"key6": true, "key7": false, "key8": true, "key9": false, "key10": true,
	}

	v1 = v
	v1.Itm_1 = 2

	v.Itm13 = &v1
}

func Benchmark1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rjson, e = json.Marshal(v)
	}
}

func main() {
	rjson, e = json.Marshal(v)
	if e != nil {
		panic(e)
	}
	e = ioutil.WriteFile("rjson.json", rjson, 0777)
	if e != nil {
		panic(e)
	}
	b1 := testing.Benchmark(Benchmark1)
	fmt.Println(b1, v.Itm_1, v1.Itm_1, "\n")
}
