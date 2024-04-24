#include <stdint.h>
#include <immintrin.h>

void avx2_add(const float *a, const float *b, float *c, const uint64_t size) {
    for (uint64_t i = 0; i < size; i += 8) {
        __m256i va = _mm256_loadu_si256((__m256i*)&a[i]);
        __m256i vb = _mm256_loadu_si256((__m256i*)&b[i]);
        __m256i vc = _mm256_add_ps(va, vb);
        _mm256_storeu_si256((__m256i*)&c[i], vc);
    }
}
