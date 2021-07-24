package code

import "testing"

func Test_convertDoubleInt2float32(t *testing.T) {
	convertDoubleInt2float32(2, 1)
	convertDoubleInt2float32(2, 2)
	convertDoubleInt2float32(2, 3)
	convertDoubleInt2float32(2, 4)
	convertDoubleInt2float32(2, 5)
	convertDoubleInt2float32(3, 1)
	convertDoubleInt2float32(3, 2)
}

func Test_convertFloat322DoubleInt(t *testing.T) {
	convertFloat322DoubleInt(2.01)
	convertFloat322DoubleInt(2.02)
	convertFloat322DoubleInt(2.03)
	convertFloat322DoubleInt(2.04)
	convertFloat322DoubleInt(2.05)
	convertFloat322DoubleInt(3.01)
	convertFloat322DoubleInt(3.02)
}

func TestFormatFloat(t *testing.T) {
	FormatFloat(5000)
}
