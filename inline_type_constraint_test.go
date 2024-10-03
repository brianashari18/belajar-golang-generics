package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindMin[T interface{ int | float64 }](first, second T) T {
	if first > second {
		return second
	} else {
		return first
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 1, FindMin(1, 2))
	assert.Equal(t, 2, FindMin(2, 3))
}

func GetFirst[T []E, E string](slice T) E {
	return slice[0]
}

func TestGetFirst(t *testing.T) {
	slice := []string{"Brian", "Anashari"}
	assert.Equal(t, "Brian", GetFirst[[]string, string](slice))
}
