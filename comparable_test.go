package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](param1, param2 T) bool {
	return param1 == param2
}

func TestComparable(t *testing.T) {
	assert.True(t, IsSame[string]("brian", "brian"))
	assert.True(t, IsSame[int](1, 1))
	assert.True(t, IsSame[bool](true, true))
}
