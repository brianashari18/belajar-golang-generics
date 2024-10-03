package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"testing"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	result := ExperimentalMin(1, 2)
	assert.Equal(t, 1, result)
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"name": "Brian",
	}

	second := map[string]string{
		"name": "Brian",
	}

	assert.True(t, maps.Equal(first, second))
}
