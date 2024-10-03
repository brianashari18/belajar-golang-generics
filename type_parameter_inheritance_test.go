package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetNamee[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	Employee
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	Employee
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func TestName(t *testing.T) {
	assert.Equal(t, "Brian", GetNamee[Manager](&MyManager{Name: "Brian"}))
	assert.Equal(t, "Brian", GetNamee[VicePresident](&MyVicePresident{Name: "Brian"}))
}
