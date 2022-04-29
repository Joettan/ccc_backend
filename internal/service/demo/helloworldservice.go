package demo

import (
	"ccc/configs"
	"ccc/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kivik/kivik/v3"
	"io/ioutil"
	"net/http"
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
	resp, error := http.Get("http://admin:1234@116.62.214.19:5984/tweets/_design/sports/_view/new-view?group=true")
	if error != nil {
		fmt.Printf(error.Error())
	}
	body, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Print(string(body))
	var sports model.SportsSceneBO
	error = json.Unmarshal(body, &sports)

	return r
}
