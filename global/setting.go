package global

import (
	"ccc/conf"
	"github.com/bitly/go-simplejson"
)

var (
	DBSetting  *conf.DBSetting
	CityConfig []string
	CitySuburb *simplejson.Json
)

type CityDO struct {
	City []string `json:"city"`
}
