//go:build noasm || (!amd64 && !arm64)

package gocc

func Add(a []float32, b []float32, out []float32) {
	panic("wrong GOARCH")
}
