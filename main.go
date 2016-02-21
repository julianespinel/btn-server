package main

import (
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	al "github.com/julianespinel/btn-server/alert"
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
	panicBusiness := pd.CreatePanicBusiness(dbConfig, pd.CreatePanicDAO())
	panicAPI := pd.CreatePanicAPI(panicBusiness)
	return panicAPI
}

func getElderAPI(dbConfig inf.DBConfig) el.ElderAPI {
	elderBusiness := el.CreateElderBusiness(dbConfig, el.CreateElderDAO())
	elderAPI := el.CreateElderAPI(elderBusiness)
	return elderAPI
}

func getAlertAPI(smsConfig inf.SmsConfig, dbConfig inf.DBConfig) al.AlertAPI {
	alertBusiness := al.CreateAlertBusiness(smsConfig, dbConfig, al.CreateAlertDAO())
	panicBusiness := pd.CreatePanicBusiness(dbConfig, pd.CreatePanicDAO())
	elderBusiness := el.CreateElderBusiness(dbConfig, el.CreateElderDAO())
	alertAPI := al.CreateAlertAPI(alertBusiness, panicBusiness, elderBusiness)
	return alertAPI
}

// run: go run main.go config.toml
func main() {

	fileName := os.Args[1]
	log.Formatter = new(logrus.JSONFormatter)
	config := getSystemConfig(fileName)
	log.Info("config", config)

	dbConfig := config.Database
	smsConfig := config.Sms
	panicAPI := getPanicAPI(dbConfig)
	elderAPI := getElderAPI(dbConfig)
	alertAPI := getAlertAPI(smsConfig, dbConfig)

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
		api.POST("/elders/:elderId/relatives", elderAPI.AddRelativeToElder())
		api.DELETE("/elders/:elderId/relatives/:relativeId", elderAPI.RemoveRelativeFromElder())
		// alert routes
		api.POST("/alerts", alertAPI.CreateAlert())
	}

	serverConfig := config.Server
	router.Run(":" + strconv.Itoa(serverConfig.Port))
}
