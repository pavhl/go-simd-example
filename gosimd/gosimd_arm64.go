//go:build !noasm && arm64

package gosimd

import (
	"unsafe"

	"github.com/alivanz/go-simd/arm"
	"github.com/alivanz/go-simd/arm/neon"
)

func Add(a []float32, b []float32, out []float32) {
	for i := 0; i < len(a); i += 4 {
		neon.VaddqF32(
			(*arm.Float32X4)(unsafe.Pointer(&out[i])),
			(*arm.Float32X4)(unsafe.Pointer(&a[i])),
			(*arm.Float32X4)(unsafe.Pointer(&b[i])),
		)
	}
}
