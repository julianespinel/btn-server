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
				inf.HandleApiError(context, err)
			}
		} else {
			inf.HandleApiError(context, bindingError)
		}
	}
	return handlerFunction
}

func (api PanicAPI) AttachElderToPanicDevice() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		serial := context.Param("serial")
		elderId := context.Param("elderId")
		business := api.panicBusiness
		_, err := business.attachElderToPanicDevice(serial, elderId)
		if err == nil {
			stringMessage := inf.GetStringMessage("message", "Elder has been attached")
			context.JSON(http.StatusOK, stringMessage)
		} else {
			inf.HandleApiError(context, err)
		}
	}
	return handlerFunction
}

func (api PanicAPI) DetachElderFromPanicDevice() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		serial := context.Param("serial")
		elderId := context.Param("elderId")
		business := api.panicBusiness
		_, err := business.detachElderFromPanicDevice(serial, elderId)
		if err == nil {
			stringMessage := inf.GetStringMessage("message", "Elder has been detached")
			context.JSON(http.StatusOK, stringMessage)
		} else {
			inf.HandleApiError(context, err)
		}
	}
	return handlerFunction
}
