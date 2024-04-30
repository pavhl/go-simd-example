//go:build !cgo || (!amd64 && !arm64)

package cgo

func Add(a []float32, b []float32, out []float32) {
	panic("target build platform was not supported or cgo was disabled")
}
