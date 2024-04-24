package main

import (
	"testing"

	"github.com/Crash129/go-simd-example/avo"
	"github.com/Crash129/go-simd-example/gocc"
	"github.com/Crash129/go-simd-example/gosimd"
	"github.com/klauspost/cpuid/v2"
)

/*
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
    go-simd-example/main_test.go:49: no ARM SIMD support
--- SKIP: BenchmarkGoSIMD
PASS
ok  	github.com/Crash129/go-simd-example	4.325s
*/

var xs, xy, out []float32

func init() {
	xs = make([]float32, 4096)
	xy = make([]float32, 4096)
	out = make([]float32, 4096)

	for i := range xs {
		xs[i] = float32(i)
		xy[i] = float32(i)
	}
}

func BenchmarkAvo(b *testing.B) {
	if !cpuid.CPU.Has(cpuid.AVX2) {
		b.Skip("no AVX2 support")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		avo.Add(xs, xy, out)
	}
}

func BenchmarkGocc(b *testing.B) {
	if !cpuid.CPU.Has(cpuid.AVX2) {
		b.Skip("no AVX2 support")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gocc.Add(xs, xy, out)
	}
}

func BenchmarkPlainGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plainGo(xs, xy, out)
	}
}

func BenchmarkGoSIMD(b *testing.B) {
	if !cpuid.CPU.Has(cpuid.ASIMD) {
		b.Skip("no ARM SIMD support")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		gosimd.Add(xs, xy, out)
	}
}
