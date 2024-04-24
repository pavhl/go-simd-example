//go:build !noasm && amd64

package avo

//go:generate go run ./generate/main.go -out ./add_amd64.s -stubs ./add_stubs_amd64.go

func Add(a []float32, b []float32, out []float32) {
	avx2_add(a, b, out)
}
