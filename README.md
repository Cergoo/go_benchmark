go_benchmark
------------
json_bench - compare benchmark test a json encode Go vs PHP  
>go  10000 op: 1.419 sec  
>php 10000 op: 0.2714 sec  
>523%

type_assertion_bench - compare benchmark test a Go type assertion operation vs Go native   
>~1% loss time a type assertion operation vs native

reflectionObj_vs_map
в два раза быстрее получить значение нативно из хеш таблицы чем из среза средствами рефлексии
bnative:	        16.6 ns/op
breflect:	        30.9 ns/op 
