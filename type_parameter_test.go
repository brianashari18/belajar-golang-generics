package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	return param
}

func TestTypeParameter(t *testing.T) {
	assert.Equal(t, 100, Length[int](100))
	assert.Equal(t, "100", Length[string]("100"))
}
