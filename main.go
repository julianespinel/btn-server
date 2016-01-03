package main

import (
	"os"
	"github.com/gin-gonic/gin"
	pd "github.com/julianespinel/btn-server/panicdevice"
	inf "github.com/julianespinel/btn-server/infrastructure"
	"github.com/BurntSushi/toml"
	"github.com/Sirupsen/logrus"
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

// run: go run main.go config.toml
func main() {

	fileName := os.Args[1]
	log.Formatter = new(logrus.JSONFormatter)
	config := getSystemConfig(fileName)
	log.Info("config", config)

	router := gin.Default()
	btn := router.Group("/btn")
	api := btn.Group("/api")
	{
		dbConfig := config.DbConfig
		panicBusiness := pd.CreatePanicBusiness(dbConfig, pd.PanicDAO{})
		panicAPI := pd.CreatePanicAPI(panicBusiness)
		api.POST("/panic-device", panicAPI.PanicRoutes())
	}
}
