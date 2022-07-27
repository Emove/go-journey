package code

import "fmt"

const (
	n1 = iota
	n2
	n3
)

const (
	m1 = iota
	m2
	_
	m4
)

const (
	z1 = iota
	z2 = 100
	z3 = iota
	z4
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
)

func PrintDemo1() {
	fmt.Println(n1, n2, n3)
}

func PrintSkip1() {
	fmt.Println(m1, m2, m4)
}

func PrintSkip2() {
	fmt.Println(z1, z2, z3, z4)
}

func PrintByteUnits() {
	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("PB: ", PB)
}

func PrintMultiAssignment() {
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
}
