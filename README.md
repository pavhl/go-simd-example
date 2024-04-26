# Examples of methods to use SIMD in Go

Demonstrated is an example which calculates the sum of two slices of float32 using SIMD instructions.

```plain
result[i] = a[i] + b[i]
```

The example is implemented in three different ways:

* code generated Assembler by [avo](https://github.com/mmcloughlin/avo) (only x86 possible)
* compile and translate to Go Assembler by [gocc](https://github.com/kelindar/gocc) (any architecture supported by clang)
* call SIMD instructions directly in Go code with [go-simd](https://github.com/alivanz/go-simd) (only ARM possible)

## how to install gocc

To install see [gocc README](https://github.com/kelindar/gocc/blob/main/README.md)

## other notable examples

See [kelindar/bitmap](https://github.com/kelindar/bitmap/tree/649a9ff772be765e8cfc9d59cc7daa8723044a01/codegen) for an example using auto-vectorization in combination with gocc.

## Benchmark results

### x86 (on AMD Ryzen 7 5700X)

```text
goos: linux
goarch: amd64
pkg: github.com/Crash129/go-simd-example
cpu: AMD Ryzen 7 5700X 8-Core Processor
BenchmarkAvo
BenchmarkAvo-16        	 3100104	       383.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGocc
BenchmarkGocc-16       	 3770598	       318.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkPlainGo
BenchmarkPlainGo-16    	  666050	      1797 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoSIMD
    main_test.go:49: no ARM SIMD support
--- SKIP: BenchmarkGoSIMD
PASS
```

Note that the assembler code of the `avo` implementation has no optimizations.
`avo` is able to have the same performance as `gocc` but the generated assembler code of `gocc` is already optimized by clang.

### ARM (on Apple M2 Pro)

```text
goos: darwin
goarch: arm64
pkg: github.com/Crash129/go-simd-example
BenchmarkAvo
    main_test.go:28: no AVX2 support
--- SKIP: BenchmarkAvo
BenchmarkGocc
BenchmarkGocc-12        15680464               374.5 ns/op
BenchmarkPlainGo
BenchmarkPlainGo-12      3339758              1798 ns/op
BenchmarkGoSIMD
BenchmarkGoSIMD-12       4593692              1309 ns/op
PASS
```