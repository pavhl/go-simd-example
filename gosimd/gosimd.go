//go:build noasm || !arm64

package gosimd

func Add(a []float32, b []float32, out []float32) {
	panic("wrong GOARCH")
}
