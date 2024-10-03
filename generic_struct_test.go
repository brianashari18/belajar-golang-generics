package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d Data[T]) ChangeFirst(first T) T {
	d.First = first
	return first
}

func (d Data[_]) SayHello() string {
	return "Hello!"
}

func TestGenericStruct(t *testing.T) {
	data := Data[string]{
		First:  "Brian",
		Second: "Anashari",
	}

	assert.Equal(t, "Hello!", data.SayHello())
	assert.Equal(t, "Celox", data.ChangeFirst("Celox"))
}
