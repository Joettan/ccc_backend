package handler

import (
	"ccc/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//依赖注入
type SceneHandler struct {
	sFactory *service.Factory
}

func NewSceneHandler(factory *service.Factory) *SceneHandler {
	return &SceneHandler{
		sFactory: factory,
	}
}

func (s *SceneHandler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("sports", s.getSports)
}

func (h *SceneHandler) getSports(c *gin.Context) {
	s := h.sFactory.SportService
	result := s.GetSports(c, "test")
	c.JSON(http.StatusOK, gin.H{"data": result})
}
