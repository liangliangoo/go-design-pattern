package test

import (
	"go-design-pattern/createpattern/proto"
	"testing"
)

var manager *proto.PrototypeManager

type Type1 struct {
	name string
}

// Clone 这里用到了指针和地址的转换
func (t *Type1) Clone() proto.Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() proto.Cloneable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	t2 := t1.Clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	} else {
		t.Log("t1 != t2")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()

	t1 := c.(*Type1)
	if t1.name != "type1" {
		t.Fatal("error")
	}

}

// init() golang中的钩子函数
func init() {
	manager = proto.NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}
