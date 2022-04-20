package configs

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-kivik/couchdb/v3"
	kivik "github.com/go-kivik/kivik/v3"
)

type couchdb struct {
}

func NewCouchDB() *couchdb {
	return &couchdb{}
}

func (c *couchdb) Hello(ctx context.Context) {
	client, err := kivik.New("couch", "http://admin:1234@9.135.78.172:5984/")
	if err != nil {
	}
	db := client.DB(context.TODO(), "helloworld")
	db.Put(ctx, "hello", gin.H{"messge": "helloworld"})
}
