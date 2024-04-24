//go:build !noasm && arm64

package gosimd

import (
	"github.com/alivanz/go-simd/arm"
	"github.com/alivanz/go-simd/arm/neon"
)

func Add(a []float32, b []float32, out []float32) {
	var va, vb, vc arm.Float32X4

	for i := 0; i < len(a); i += 4 {
		va[0] = arm.Float32(a[i])
		va[1] = arm.Float32(a[i+1])
		va[2] = arm.Float32(a[i+2])
		va[3] = arm.Float32(a[i+3])

		vb[0] = arm.Float32(b[i])
		vb[1] = arm.Float32(b[i+1])
		vb[2] = arm.Float32(b[i+2])
		vb[3] = arm.Float32(b[i+3])

		neon.VaddqF32(vc, va, vb)

		out[i] = float32(vc[0])
		out[i+1] = float32(vc[1])
		out[i+2] = float32(vc[2])
		out[i+3] = float32(vc[3])
	}
}
