package demo

import (
	"ccc/configs"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kivik/kivik/v3"
)

type IHelloWorldService interface {
	HelloWorld(s string)
}

type HelloWorldService struct {
	db *kivik.DB
}

func NewHelloWorldService() *HelloWorldService {
	couchdb := configs.NewCouchDB()
	db := couchdb.DB()
	return &HelloWorldService{
		db: db,
	}
}

func (h *HelloWorldService) HelloWorld(ctx *gin.Context, s string) string {
	r := fmt.Sprintf("This is a fucking ccc service %s", s)
	h.db.Put(context.Background(), "test", gin.H{
		"message": r,
	})
	return r
}
