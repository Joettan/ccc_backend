package handler

import (
	"ccc/internal/model"
	"ccc/internal/service"
	"github.com/gin-gonic/gin"
	"log"
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
	group.GET("scene", s.getScene)
}

func (h *SceneHandler) getScene(c *gin.Context) {
	s := h.sFactory.SportService
	var request model.SceneRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Default().Printf("Error Request %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"data": "Error Request Wrong Value Type"})
		return
	}
	result := s.GetMetrics(c, &request)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
