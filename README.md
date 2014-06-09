go_benchmark test
------------

reflectionObj_vs_map
в два раза быстрее получить значение нативно из хеш таблицы чем из среза средствами рефлексии
bnative:	        16.6 ns/op
breflect:	        30.9 ns/op 
