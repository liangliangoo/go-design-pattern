package test

import (
	"go-design-pattern/createpattern/factorymethod"
	"testing"
)

func compute(factory factorymethod.OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {

	var (
		factory factorymethod.OperatorFactory
	)

	factory = factorymethod.PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}

	factory = factorymethod.MinusOperatorFactory{}
	if compute(factory, 4, 2) != 2 {
		t.Fatal("error with factory method pattern")
	}
}
