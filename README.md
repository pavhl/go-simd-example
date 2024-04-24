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
