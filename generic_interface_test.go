package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(param T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (m *MyData[T]) GetValue() T {
	return m.Value
}

func (m *MyData[T]) SetValue(param T) {
	m.Value = param
}

func TestGenericInterface(t *testing.T) {
	myData := MyData[string]{Value: "Brian"}
	result := ChangeValue[string](&myData, "Anashari")

	assert.Equal(t, "Anashari", result)
}
