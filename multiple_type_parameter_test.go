package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Multiple[T1 any, T2 any](param1 T1, param2 T2) (T1, T2) {
	return param1, param2
}

func TestMultipleTypeParameter(t *testing.T) {
	a, b := Multiple[int, string](100, "100")
	assert.Equal(t, a, 100)
	assert.Equal(t, b, "100")
}
