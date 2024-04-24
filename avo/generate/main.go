package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
)

func main() {
	TEXT("avx2_add", NOSPLIT, "func(a, b, out []float32)")
	Doc("avx2_add adds a and b into out using AVX2 instructions.")
	a := Load(Param("a").Base(), GP64())
	b := Load(Param("b").Base(), GP64())
	out := Load(Param("out").Base(), GP64())

	// Get the number of elements to process
	n := Load(Param("a").Len(), GP64())
	i := GP64()
	XORQ(i, i) // i = 0

	Label("loop")

	// Load 8 floats from a and b
	va := YMM()
	vb := YMM()
	VMOVDQU(Mem{Base: a, Index: i, Scale: 4}, va)
	VMOVDQU(Mem{Base: b, Index: i, Scale: 4}, vb)

	// Add them together
	vc := YMM()
	VADDPS(va, vb, vc)

	// Store the result
	VMOVDQU(vc, Mem{Base: out, Index: i, Scale: 4})

	// increase loop variable
	ADDQ(U8(8), i)

	// Jump back if i < n
	CMPQ(i, n)
	JB(LabelRef("loop"))

	RET()
	Generate()
}
