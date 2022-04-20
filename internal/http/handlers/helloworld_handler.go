package handler

import (
	"ccc/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//依赖注入
type HelloWorldHandler struct {
	sFactory *service.Factory
}

func NewHelloWorldHandler(factory *service.Factory) *HelloWorldHandler {
	return &HelloWorldHandler{
		sFactory: factory,
	}
}

func (h *HelloWorldHandler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("HelloWorld", h.HelloWorld)
}

func (h *HelloWorldHandler) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "HelloWorld"})
}
