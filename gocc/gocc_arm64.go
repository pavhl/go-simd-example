//go:build !noasm && arm64 && !darwin

package gocc

import "unsafe"

//go:generate gocc ./generate/arm.c --arch neon -O3 -o . --package gocc --local

func Add(a []float32, b []float32, out []float32) {
	neon_add(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&out[0]), uint64(len(a)))
}
