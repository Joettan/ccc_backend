package main

import (
	handler "ccc/internal/http/handlers"
	"ccc/internal/http/routers"
	"ccc/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initGin() *gin.Engine {
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	factory := service.NewFactory()
	v1Api := routers.NewAPI(handler.NewHelloWorldHandler(factory))
	v1Api.RegisterRouter(server)
	return server
}

func main() {
	server := initGin()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
