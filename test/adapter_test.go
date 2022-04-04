package test

import (
	"go-design-pattern/structure/adapter"
	"testing"
)

var expect = "adaptee method"

func TestAdapter(t *testing.T) {
	adaptee := adapter.NewAdaptee()
	target := adapter.NewAdapter(adaptee)
	res := target.Request()
	if res != expect {
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
