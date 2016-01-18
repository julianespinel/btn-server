package main

import (
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	el "github.com/julianespinel/btn-server/elder"
	inf "github.com/julianespinel/btn-server/infrastructure"
	pd "github.com/julianespinel/btn-server/panicdevice"
)

var log = logrus.New()

func getSystemConfig(fileName string) inf.Config {
	var config inf.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		// handle error
		log.Error(err)
	}
	return config
}

func getPanicAPI(dbConfig inf.DBConfig) pd.PanicAPI {
	panicBusiness := pd.CreatePanicBusiness(dbConfig, pd.PanicDAO{})
	panicAPI := pd.CreatePanicAPI(panicBusiness)
	return panicAPI
}

func getElderAPI(dbConfig inf.DBConfig) el.ElderAPI {
	elderBusiness := el.CreateElderBusiness(dbConfig, el.ElderDAO{})
	elderAPI := el.CreateElderAPI(elderBusiness)
	return elderAPI
}

// run: go run main.go config.toml
func main() {

	fileName := os.Args[1]
	log.Formatter = new(logrus.JSONFormatter)
	config := getSystemConfig(fileName)
	log.Info("config", config)

	dbConfig := config.Database
	panicAPI := getPanicAPI(dbConfig)
	elderAPI := getElderAPI(dbConfig)

	router := gin.Default()
	btn := router.Group("/btn")
	api := btn.Group("/api")
	{
		// panic devices routes
		api.POST("/panic-devices", panicAPI.CreatePanicDevice())
		api.PUT("/panic-devices/:serial/elders/:elderId", panicAPI.AttachElderToPanicDevice())
		api.DELETE("/panic-devices/:serial/elders/:elderId", panicAPI.DetachElderFromPanicDevice())
		// elders routes
		api.POST("/elders", elderAPI.CreateElder())
	}

	serverConfig := config.Server
	router.Run(":" + strconv.Itoa(serverConfig.Port))
}
