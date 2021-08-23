package adapter

/*
适配器模式：
用于转换一种接口适配另一种接口
实际使用中Adaptee一般为接口，并且使用工厂函数生成实例
在Adapter中匿名组合Adaptee接口，所以Adapter类也拥有SpecificRequest实例方法，又因为Go语言中非侵入式接口特征。
Adapter也适配Adaptee接口
*/

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type adapteeImpl struct {
}

func (a adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

type adapter struct {
	Adaptee
}

func (a adapter) Request() string {
	return a.SpecificRequest()
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}
