//go:build cgo && amd64

package cgo

/*
   #include <stdint.h>
   #include <immintrin.h>

   void avx2_add(const float *a, const float *b, float *c, const uint64_t size) {
       for (uint64_t i = 0; i < size; i += 8) {
           __m256 va = _mm256_loadu_ps(&a[i]);
           __m256 vb = _mm256_loadu_ps(&b[i]);
           __m256 vc = _mm256_add_ps(va, vb);
           _mm256_storeu_ps(&c[i], vc);
       }
   }
*/
import "C"

func Add(a []float32, b []float32, out []float32) {
	C.avx2_add((*C.float)(&a[0]), (*C.float)(&b[0]), (*C.float)(&out[0]), C.uint64_t(len(a)))
}
