package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Age int

type Number interface {
	~int | int8 | int16 | int32 | int64 | float32 | float64
}

func GetNumber[T Number](num1, num2 T) T {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func TestTypeSet(t *testing.T) {
	assert.Equal(t, int32(100), GetNumber[int32](100, 50))
	assert.Equal(t, float64(100), GetNumber[float64](100, 50))
	assert.Equal(t, int64(100), GetNumber[int64](100, 50))
	assert.Equal(t, Age(100), GetNumber(Age(100), Age(50)))
}
