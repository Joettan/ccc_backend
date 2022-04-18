package demo

import "fmt"

type IHelloWorldService interface {
	HelloWorld(s string)
}

type HelloWorldService struct {
}

func (h *HelloWorldService) HelloWorld(s string) string {
	r := fmt.Sprintf("This is a fucking ccc service %s", s)
	return r
}
