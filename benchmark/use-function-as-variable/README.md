# Benchmark Use Function as a Variable

This project tries to understand the performance impact if we use functions declared in functions as variables instead of functions defined in the global scope.

## Result

```sh
make
go test -v -bench=. -run=none . -benchmem
goos: linux
goarch: amd64
pkg: benchmark
cpu: 11th Gen Intel(R) Core(TM) i9-11900 @ 2.50GHz
BenchmarkStringCombine1
BenchmarkStringCombine1-16      99380000                11.63 ns/op            0 B/op          0 allocs/op
BenchmarkStringCombine2
BenchmarkStringCombine2-16      45309477                31.21 ns/op           16 B/op          1 allocs/op
PASS
ok      benchmark       2.613s
```

The perfomace of `StringCombine1` is much better than that of `StringCombine2`. This is obvious because a new space is needed in `StringCombine2` to store the pointer to the function.

If you add `fmt.Println(&c)` to `StringCombine2`, you will find that each c is different and cannot be reused.
