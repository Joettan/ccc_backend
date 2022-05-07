package service

import (
	"ccc/couchdb"
	"ccc/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type SceneService struct {
}

func NewSportService() *SceneService {
	return &SceneService{}
}

const db = "http://admin:1234@116.62.214.19:5984"

func (h *SceneService) GetMetrics(ctx *gin.Context, r *model.SceneRequest) *model.SceneMetricsVO {
	QueryString := couchdb.GetQueryString(db, r)
	resp, error := http.Get(QueryString)
	if error != nil {
		fmt.Printf(error.Error())
	}
	body, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Print(string(body))
	var sports model.SceneRowsBO
	err := json.Unmarshal(body, &sports)
	if err != nil {
		log.Default().Printf("error unmarshal json: %v", err.Error())
		return nil
	}
	sceneSlice := make([]*model.SceneVO, 0, len(sports.Rows))
	for i, row := range sports.Rows {
		year, _ := strconv.Atoi(row.Key[0])
		sportsVO := &model.SceneVO{
			Id:       i,
			Location: row.Key[1],
			Scores:   row.Value,
			Year:     year,
		}
		sceneSlice = append(sceneSlice, sportsVO)
	}
	sportsSceneVO := model.SceneMetricsVO{
		Metrics: sceneSlice,
	}
	return &sportsSceneVO
}
