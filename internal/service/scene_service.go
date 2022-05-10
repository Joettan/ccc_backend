package service

import (
	"ccc/couchdb"
	"ccc/global"
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

func (h *SceneService) GetMetrics(ctx *gin.Context, r *model.SceneRequest) *model.SceneMetricsVO {
	DBAddress := global.DBSetting.DBAddress
	QueryString := couchdb.GetQueryString(DBAddress, r)
	log.Default().Printf(QueryString)
	resp, error := http.Get(QueryString)
	if error != nil {
		fmt.Printf(error.Error())
	}
	body, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Print(string(body))
	var sports model.SceneRowsDO
	err := json.Unmarshal(body, &sports)
	if err != nil {
		log.Default().Printf("error unmarshal json: %v", err.Error())
		return nil
	}
	sceneSlice := make([]*model.SceneVO, 0, len(sports.Rows))
	sceneBOMap := make(map[string]*model.SceneBO, 0)
	for _, row := range sports.Rows {
		yearString := row.Key[0]
		year, _ := strconv.Atoi(row.Key[0])
		location := row.Key[1]
		key := yearString + location
		sceneBO, ok := sceneBOMap[key]
		if !ok {
			sceneBO = &model.SceneBO{
				Location: location,
				Year:     year,
			}
			sceneBOMap[key] = sceneBO
		}
		sentiment := row.Key[2]
		switch sentiment {
		case "neg":
			sceneBO.NegativeScore = row.Value
		case "neu":
			sceneBO.NeutralScore = row.Value
		case "pos":
			sceneBO.PositiveScore = row.Value
		}
	}
	var i = 0
	for _, sceneBO := range sceneBOMap {
		sceneVO := &model.SceneVO{
			Year:          sceneBO.Year,
			Location:      sceneBO.Location,
			Id:            i,
			NegativeScore: sceneBO.NegativeScore,
			NeutralScore:  sceneBO.NeutralScore,
			PositiveScore: sceneBO.PositiveScore,
		}
		i++
		sceneVO.Scores = (4*sceneBO.PositiveScore + 2*sceneBO.NeutralScore - 4*sceneBO.NegativeScore) / 10
		sceneSlice = append(sceneSlice, sceneVO)
	}
	sportsSceneVO := model.SceneMetricsVO{
		Metrics: sceneSlice,
	}
	return &sportsSceneVO
}
