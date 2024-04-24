//go:build !noasm && amd64

package gocc

import (
	"unsafe"
)

//go:generate gocc ./generate/x86.c --arch avx2 -O3 -o . --package gocc --local

func Add(a []float32, b []float32, out []float32) {
	avx2_add(unsafe.Pointer(&a[0]), unsafe.Pointer(&b[0]), unsafe.Pointer(&out[0]), uint64(len(a)))
}
