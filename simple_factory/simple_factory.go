package simple_factory

import "fmt"

/*
简单的工厂模式
go 中没有构造函数，通过NewXXX函数来初始化类对象。
封装了实现细节，仅仅暴露了API接口以及NewAPI函数。
*/
type API interface {
	Say(name string) string
}

type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("hi, %s", name)
}

type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("hello, %s", name)
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	}
	if t == 2 {
		return &helloAPI{}
	}
	return nil
}
