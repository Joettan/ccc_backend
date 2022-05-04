package service

import (
	"ccc/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SportService struct {
}

func NewSportService() *SportService {
	return &SportService{}
}

func (h *SportService) GetSports(ctx *gin.Context, s string) model.SportsSceneVO {
	resp, error := http.Get("http://admin:1234@116.62.214.19:5984/tweets/_design/sports/_view/new-view?group=true")
	if error != nil {
		fmt.Printf(error.Error())
	}
	body, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Print(string(body))
	var sports model.SportsSceneBO
	error = json.Unmarshal(body, &sports)
	sportsVOSlice := make([]*model.SportVO, 0, len(sports.Rows))
	for i, row := range sports.Rows {
		year, _ := strconv.Atoi(row.Key[0])
		sportsVO := &model.SportVO{
			Id:       i,
			Location: row.Key[1],
			Year:     year,
			Scores:   row.Value,
		}
		sportsVOSlice = append(sportsVOSlice, sportsVO)
	}
	sportsSceneVO := model.SportsSceneVO{
		Metrics: sportsVOSlice,
	}
	return sportsSceneVO
}
