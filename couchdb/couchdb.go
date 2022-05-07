package couchdb

import (
	"ccc/internal/model"
	"fmt"
)

type Couchdb struct {
}

func GetQueryString(db string, r *model.SceneRequest) string {
	var view string
	var yearCondition string
	if r.Year == 0 {
		view = "_design/scene/_view/location"
		yearCondition = ""
	} else {
		view = "_design/scene/_view/yearlocation"
		yearCondition = fmt.Sprintf("&startkey=[\"%v\"]&endkey=[\"%v\"]", r.Year, r.Year+1)
	}
	q := fmt.Sprintf("%s/%s/%s?group=true%s", db, "tweets", view, yearCondition)
	return q
}
