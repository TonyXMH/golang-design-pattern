package facade

import "fmt"

/*
外观模式：
外观接口目的简化对facade类的访问
facade模块同事暴露了  a和b模块，其他代码需要使用细节可以直接调用
*/

type API interface {
	Test() string
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

type aModuleAPI struct {
}

func (a aModuleAPI) TestA() string {
	return "A module running"
}

type bModuleAPI struct {
}

func (b bModuleAPI) TestB() string {
	return "B module running"
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleAPI{}
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleAPI{}
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAPI() API {
	return apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}
