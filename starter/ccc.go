package main

import (
	"ccc/conf"
	"ccc/global"
	handler "ccc/internal/http/handlers"
	"ccc/internal/http/middleware"
	"ccc/internal/http/routers"
	"ccc/internal/service"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func setupSetting() error {
	setting, err := conf.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("DB", &global.DBSetting)
	if err != nil {
		return err
	}
	return nil
}

func initGin() *gin.Engine {
	global.CitySuburb = ReadCitySuburbConfig()
	s := global.CitySuburb.Get("452")
	fmt.Println(s)
	global.CityConfig = ReadCityConfig()
	err := setupSetting()
	for err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(middleware.Cors())
	factory := service.NewFactory()
	v1Api := routers.NewAPI(handler.NewHelloWorldHandler(factory), handler.NewSceneHandler(factory), handler.NewRegionHandler(factory))
	v1Api.RegisterRouter(server)
	return server
}

func ReadCityConfig() []string {
	var cityDO global.CityDO
	filePtr, err := os.Open("./conf/city_name.json")
	if err != nil {
		log.Default().Printf("can't read city_name.json")
	}
	defer filePtr.Close()
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&cityDO)
	if err != nil {
		log.Default().Printf("can't decode city_name.json")
	}
	return cityDO.City
}

func ReadCitySuburbConfig() *simplejson.Json {
	var suburbJson *simplejson.Json
	filePtr, err := os.Open("./conf/city_suburb_map.json")
	if err != nil {
		log.Default().Printf("can't read city_suburb_map.json")
	}
	defer filePtr.Close()
	suburbJson, err = simplejson.NewFromReader(filePtr)
	if err != nil {
		log.Default().Printf("can't read city_suburb_map.json")
	}
	return suburbJson
}

func main() {
	server := initGin()
	go func() {
		_ = http.ListenAndServe("0.0.0.0:8000", nil)
	}()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        server,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
