//go:build !noasm && arm64

package gocc

import "unsafe"

//go:generate gocc ./generate/arm.c --arch neon -O3 -o . --package gocc --local
//go:generate gocc ./generate/apple.c --arch apple -O3 -o . --package gocc --local

func Add(a []float32, b []float32, out []float32) {
	arm_add(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&out[0]), uint64(len(a)))
}
