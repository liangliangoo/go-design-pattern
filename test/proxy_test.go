package test

import (
	"go-design-pattern/structure/proxy"
	"testing"
)

func TestProxy(t *testing.T) {
	var sub proxy.Subject
	sub = &proxy.Proxy{}

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}
