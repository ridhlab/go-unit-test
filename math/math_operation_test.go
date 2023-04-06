package mathOperation

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("Test started")
	m.Run()
	fmt.Println("Test finished")
}

func TestSum(t *testing.T) {
	assert.Equal(t, 4, Sum(2, 2), "Sum result must be equal")
	assert.Equal(t, 22, Sum(10, 12), "Sum result must be equal")
	assert.Equal(t, 27, Sum(13, 14), "Sum result must be equal")

	assert.NotEqual(t, 5, Sum(2, 2), "Sum result must not equal")
	assert.NotEqual(t, 11, Sum(10, 2), "Sum result must not equal")
	assert.NotEqual(t, 9, Sum(8, 2), "Sum result must not equal")
}

func TestSubstraction(t *testing.T) {
	assert.Equal(t, 0, Substraction(2, 2), "Substraction result must be equal")
	assert.Equal(t, -2, Substraction(10, 12), "Substraction result must be equal")
	assert.Equal(t, -1, Substraction(13, 14), "Substraction result must be equal")

	assert.NotEqual(t, 1, Substraction(2, 2), "Substraction result must not equal")
	assert.NotEqual(t, 11, Substraction(10, 2), "Substraction result must not equal")
	assert.NotEqual(t, 9, Substraction(8, 2), "Substraction result must not equal")
}

func TestMultiple(t *testing.T) {
	assert.Equal(t, 4, Multiple(2, 2), "Multiple result must be equal")
	assert.Equal(t, 120, Multiple(10, 12), "Multiple result must be equal")
	assert.Equal(t, 182, Multiple(13, 14), "Multiple result must be equal")

	assert.NotEqual(t, 5, Multiple(2, 2), "Multiple result must not equal")
	assert.NotEqual(t, 11, Multiple(10, 2), "Multiple result must not equal")
	assert.NotEqual(t, 9, Multiple(8, 2), "Multiple result must not equal")
}

func TestDivide(t *testing.T) {
	assert.Equal(t, 1, Divide(2, 2), "Divide result must be equal")
	assert.Equal(t, 4, Divide(4, 1), "Divide result must be equal")
	assert.Equal(t, 3, Divide(40, 12), "Divide result must be equal")

	assert.NotEqual(t, 5, Divide(2, 2), "Divide result must not equal")
	assert.NotEqual(t, 11, Divide(10, 2), "Divide result must not equal")
	assert.NotEqual(t, 9, Divide(8, 2), "Divide result must not equal")
}
