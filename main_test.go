package main

import (
	"testing"

	"github.com/klauspost/cpuid/v2"

	"github.com/Crash129/go-simd-example/avo"
	"github.com/Crash129/go-simd-example/gocc"
	"github.com/Crash129/go-simd-example/gosimd"
)

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
	if !cpuid.CPU.Has(cpuid.AVX2) && !cpuid.CPU.Has(cpuid.ASIMD) {
		b.Skip("no AVX2 or ARM SIMD support")
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
