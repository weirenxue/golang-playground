# Benchmark Set Slice

This project tries to find a more efficient way to set the slice value in a for loop.

## Result

```sh
make
go test -v -bench=. -run=none . -benchmem
goos: linux
goarch: amd64
pkg: slice-benchmark
cpu: 11th Gen Intel(R) Core(TM) i9-11900 @ 2.50GHz
BenchmarkSetSlice01
BenchmarkSetSlice01-16            175076              6171 ns/op           81920 B/op          1 allocs/op
BenchmarkSetSlice02
BenchmarkSetSlice02-16            186766              6156 ns/op           81920 B/op          1 allocs/op
BenchmarkSetSlice03
BenchmarkSetSlice03-16            164215              7291 ns/op           81920 B/op          1 allocs/op
BenchmarkSetSlice04
BenchmarkSetSlice04-16             43023             27342 ns/op          357625 B/op         19 allocs/op
BenchmarkSetSlice05
BenchmarkSetSlice05-16             10000            238242 ns/op          161920 B/op      10001 allocs/op
```

The performance of `SetSlice01` and `SetSlice02` are almost the same and the best, they are initialized with the appropriate length and the set values are accessed through the slice index.

The second ranked method is `SetSlice03`, which is initilized with enough capacity but zero length and sets the value by `append`. This makes sense. We can imagine more actions in the `append` function than just accessing through the slice index.

The third ranked method is `SetSlice04`, which is initilized with zero capacity and sets the value by `append`. To increase the capacity of a slice one must create a new, larger slice and copy the contents of the original slice into it.

The worst method is `SetSlice05`, which is almost the same as `SetSlice01` and `SetSlice02`. The only difference is that it is a slice of struct's pointer. This makes memory access discontinuous.
