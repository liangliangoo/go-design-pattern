package simplefactory

import "fmt"

// 简单工厂模式

type Api interface {
	Say(str string)
}

func NewApi(i interface{}) Api {
	switch i.(type) {
	case HelloApi:
		return &HelloApi{}
	case HiApi:
		return &HiApi{}
	//	这里只是简单的测试没有写那么多结构体
	// ...	return  other type,
	default:
		return nil
	}
}

type HelloApi struct {
}

type HiApi struct {
}

func (*HelloApi) Say(str string) {
	fmt.Println("hello, " + str)
}

func (*HiApi) Say(str string) {
	fmt.Println("hi, " + str)
}
