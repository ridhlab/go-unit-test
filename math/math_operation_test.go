package mathOperation

import "testing"

func TestSum(t *testing.T) {
	res := Sum(1, 2)
	if res != 3 {
		panic("Test is not valid")
	}
}

func TestSubstraction(t *testing.T) {
	res := Substraction(1, 2)
	if res != -1 {
		panic("Test is not valid")
	}
}

func TestMultiple(t *testing.T) {
	res := Multiple(1, 2)
	if res != 2 {
		panic("Test is not valid")
	}
}

func TestDivide(t *testing.T) {
	res1 := Divide(1, 2)
	res2 := Divide(2, 2)
	if res1 != 0 || res2 != 1 {
		panic("Test is not valid")
	}
}
