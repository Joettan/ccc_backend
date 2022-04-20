package demo

import (
	"ccc/configs"
	"fmt"
	"github.com/gin-gonic/gin"
)

type IHelloWorldService interface {
	HelloWorld(s string)
}

type HelloWorldService struct {
}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{}
}

func (h *HelloWorldService) HelloWorld(ctx *gin.Context, s string) string {
	r := fmt.Sprintf("This is a fucking ccc service %s", s)
	couchdb := configs.NewCouchDB()
	couchdb.Hello(ctx)
	return r
}
