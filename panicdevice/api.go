package panicdevice

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inf "github.com/julianespinel/btn-server/infrastructure"
)

type PanicAPI struct {
	panicBusiness PanicBusiness
}

func CreatePanicAPI(panicBusiness PanicBusiness) PanicAPI {
	return PanicAPI{panicBusiness: panicBusiness}
}

func handleApiError(context *gin.Context, err error) {
	context.JSON(-1, context.Error(err)) // -1 == not override the current error code
}

func (api PanicAPI) CreatePanicDevice() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		var device PanicDevice
		bindingError := context.Bind(&device)
		if bindingError == nil {
			business := api.panicBusiness
			createdPanicDevice, err := business.createPanicDevice(device)
			if err == nil {
				context.JSON(http.StatusCreated, createdPanicDevice)
			} else {
				handleApiError(context, err)
			}
		} else {
			handleApiError(context, bindingError)
		}
	}
	return handlerFunction
}

func (api PanicAPI) AttachElderToDevice() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		serial := context.Param("serial")
		elderId := context.Param("elderId")
		business := api.panicBusiness
		_, err := business.attachElderToPanicDevice(serial, elderId)
		if err == nil {
			stringMessage := inf.GetStringMessage("message", "The elder has been attached.")
			context.JSON(http.StatusOK, stringMessage)
		} else {
			handleApiError(context, err)
		}
	}
	return handlerFunction
}
