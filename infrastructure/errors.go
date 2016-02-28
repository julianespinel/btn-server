package infrastructure

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

var log = logrus.New()

func HandleApiErrorWithStatusCode(context *gin.Context, errorCode int, err error) {
	context.JSON(errorCode, context.Error(err)) // -1 == not override the current error code
}

func HandleApiError(context *gin.Context, err error) {
	context.JSON(500, context.Error(err)) // -1 == not override the current error code
}

func HandleDBError(err error) {
	log.Error(err)
}
