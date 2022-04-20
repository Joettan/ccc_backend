package configs

import (
	"context"
	_ "github.com/go-kivik/couchdb/v3"
	kivik "github.com/go-kivik/kivik/v3"
)

type Couchdb struct {
}

func NewCouchDB() *Couchdb {
	return &Couchdb{}
}

func (c *Couchdb) DB() *kivik.DB {
	client, err := kivik.New("couch", "http://admin:1234@9.135.78.172:5984/")
	if err != nil {
	}
	db := client.DB(context.TODO(), "helloworld")
	return db
}
