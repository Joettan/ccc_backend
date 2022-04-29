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
	client, err := kivik.New("couch", "http://admin:1234@116.62.214.19:5984")
	if err != nil {
	}
	db := client.DB(context.TODO(), "tweets")
	return db
}
