//go:build cgo && arm64

package cgo

/*
#include <stdint.h>
#include <arm_neon.h>

void arm_add(const float *a, const float *b, float *c, const uint64_t size) {
    for (uint64_t i = 0; i < size; i += 4) {
        float32x4_t va = vld1q_f32(&a[i]);
        float32x4_t vb = vld1q_f32(&b[i]);
        float32x4_t vc = vaddq_f32(va, vb);
        vst1q_f32(&c[i], vc);
    }
}
*/
import "C"

func Add(a []float32, b []float32, out []float32) {
	C.arm_add((*C.float)(&a[0]), (*C.float)(&b[0]), (*C.float)(&out[0]), C.uint64_t(len(a)))
}
