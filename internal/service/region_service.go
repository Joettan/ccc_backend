package service

import (
	"ccc/couchdb"
	"ccc/global"
	"ccc/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type RegionService struct {
}

func NewRegionService() *RegionService {
	return &RegionService{}
}

func (r *RegionService) GetWeatherData() {

}

func (r *RegionService) GetSports(locationPid string) {

}

func (r *RegionService) GetFoods(locationPid string) interface{} {
	dBAddress := global.DBSetting.DBAddress
	view := "_design/bars/_view/data"
	barQueryString := couchdb.GetQueryString(dBAddress, "aurin-bars", view)
	log.Default().Printf(barQueryString)
	resp, error := http.Get(barQueryString)
	if error != nil {
		fmt.Printf(error.Error())
	}
	body, error := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var regionRowsDO model.RegionRowsDO
	err := json.Unmarshal(body, &regionRowsDO)
	if err != nil {
		log.Default().Printf("error unmarshal json: %v", err.Error())
		return nil
	}
	resultMap := make(map[string]*model.FoodsVO, 0)
	for _, row := range regionRowsDO.Rows {
		value := row.Value
		location := row.Key[0]
		vo := &model.FoodsVO{
			BarsScore:   value.Max,
			LocationPid: location,
		}
		resultMap[location] = vo
	}
	view = "_design/cafes/_view/data"
	cafeQueryString := couchdb.GetQueryString(dBAddress, "aurin-cafes", view)
	log.Default().Printf(cafeQueryString)
	respV2, error := http.Get(cafeQueryString)
	if error != nil {
		fmt.Printf(error.Error())
	}
	bodyV2, error := ioutil.ReadAll(respV2.Body)
	defer respV2.Body.Close()
	var regionRowsDOV2 model.RegionRowsDO
	err = json.Unmarshal(bodyV2, &regionRowsDOV2)
	if err != nil {
		log.Default().Printf("error unmarshal json: %v", err.Error())
		return nil
	}
	for _, row := range regionRowsDOV2.Rows {
		value := row.Value
		location := row.Key[0]
		vo, ok := resultMap[location]
		if !ok {
			vo = &model.FoodsVO{
				CafesScore:  value.Max,
				LocationPid: location,
			}
			resultMap[location] = vo
		} else {
			vo.CafesScore = value.Max
		}
	}
	resultSlice := make([]*model.FoodsVO, 0)
	for _, vo := range resultMap {
		if vo.BarsScore*vo.CafesScore != 0 {
			vo.Ratio = vo.BarsScore / vo.CafesScore
		} else {
			vo.Ratio = 0
		}
		resultSlice = append(resultSlice, vo)
	}

	if locationPid != "" {
		filterResultSlice := make([]*model.FoodsVO, 0)
		for _, vo := range filterResultSlice {
			if vo.LocationPid == locationPid {
				filterResultSlice = append(filterResultSlice, vo)
			}
		}
		return filterResultSlice
	}
	return resultSlice
}
