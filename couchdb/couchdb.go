package couchdb

import (
	"fmt"
)

type Couchdb struct {
}

func GetQueryString(dbAddress string, dbName string, view string) string {
	//var yearCondition string
	/*	if r.Year == 0 {
			view = "_design/scene/_view/location"
			yearCondition = ""
		} else {
			view = "_design/scene/_view/yearlocation"
			yearCondition = fmt.Sprintf("&startkey=[\"%v\"]&endkey=[\"%v\"]", r.Year, r.Year+1)
		}*/
	q := fmt.Sprintf("%s/%s/%s?group=true", dbAddress, dbName, view)
	return q
}
