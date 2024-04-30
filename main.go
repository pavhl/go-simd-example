package main

import (
	"fmt"

	"github.com/klauspost/cpuid/v2"

	"github.com/Crash129/go-simd-example/avo"
	"github.com/Crash129/go-simd-example/cgo"
	"github.com/Crash129/go-simd-example/gocc"
	"github.com/Crash129/go-simd-example/gosimd"
)

func main() {
	a := make([]float32, 32)
	b := make([]float32, 32)
	out := make([]float32, 32)

	for i := range a {
		a[i] = float32(i + 1)
		b[i] = float32(100)
	}

	fmt.Printf("a:\t%+v\n", a)
	fmt.Printf("b:\t%+v\n", b)
	fmt.Println()

	clear(out)
	plainGo(a, b, out)
	fmt.Printf("plain:\t%+v\n", out)

	if cpuid.CPU.Has(cpuid.ASIMD) {
		clear(out)
		cgo.Add(a, b, out)
		fmt.Printf("cgo:\t%+v\n", out)

		clear(out)
		gosimd.Add(a, b, out)
		fmt.Printf("gosimd:\t%+v\n", out)
	}
	if cpuid.CPU.Has(cpuid.AVX2) {
		clear(out)
		cgo.Add(a, b, out)
		fmt.Printf("cgo:\t%+v\n", out)

		clear(out)
		gocc.Add(a, b, out)
		fmt.Printf("gocc:\t%+v\n", out)

		clear(out)
		avo.Add(a, b, out)
		fmt.Printf("avo:\t%+v\n", out)
	}

}

func plainGo(a []float32, b []float32, out []float32) {
	for i := range a {
		out[i] = a[i] + b[i]
	}
}
