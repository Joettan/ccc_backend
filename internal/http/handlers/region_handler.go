package handler

import (
	"ccc/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//依赖注入
type RegionHandler struct {
	sFactory *service.Factory
}

func NewRegionHandler(factory *service.Factory) *RegionHandler {
	return &RegionHandler{
		sFactory: factory,
	}
}

func (h *RegionHandler) RegisterRoutes(group *gin.RouterGroup) {
	group.GET("weather", h.getWeatherData)
	group.GET("sports", h.getSportsData)
	group.GET("foods", h.getFoodsData)
}

func (h *RegionHandler) getWeatherData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello"})
}

func (h *RegionHandler) getSportsData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello"})
}

func (h *RegionHandler) getFoodsData(c *gin.Context) {
	r := h.sFactory.RegionService
	location := c.Query("location")
	result := r.GetFoods(location)
	c.JSON(http.StatusOK, gin.H{"data": result})
}
