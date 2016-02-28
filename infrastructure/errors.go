package infrastructure

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var log = logrus.New()

func HandleApiError(context *gin.Context, err error) {
	context.JSON(-1, context.Error(err)) // -1 == not override the current error code
}

func HandleOptionalApiError(context *gin.Context, err error) {
	if err != nil {
		context.JSON(-1, context.Error(err)) // -1 == not override the current error code
	}
}

func HandleDBError(err error) {
	log.Error(err)
}
