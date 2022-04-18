package service

/**
实现依赖注入的工厂类，配置相应service服务
*/
type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}
