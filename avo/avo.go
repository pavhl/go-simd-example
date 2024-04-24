//go:build noasm || !amd64

package avo

func Add(a []float32, b []float32, out []float32) {
	panic("wrong GOARCH")
}
