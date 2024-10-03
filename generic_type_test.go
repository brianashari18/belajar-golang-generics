package golang_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, item := range bag {
		fmt.Println(item)
	}
}

func TestGenericType(t *testing.T) {
	PrintBag[string](Bag[string]{"Brian", "Anashari"})
	PrintBag[int](Bag[int]{1, 2, 3, 4, 5})

}
