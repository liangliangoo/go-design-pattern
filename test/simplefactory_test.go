package test

import (
	"go-design-pattern/createpattern/simplefactory"
	"testing"
)

func TestNewApi(t *testing.T) {
	api := simplefactory.NewApi(simplefactory.HelloApi{})
	api.Say("simple factory!")
}
