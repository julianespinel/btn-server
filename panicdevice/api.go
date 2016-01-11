package panicdevice

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	routes := func(context *gin.Context) {
		var device PanicDevice
		bindingError := context.Bind(&device)
		if bindingError == nil {
			business := api.panicBusiness
			createdPanicDevice := business.createPanicDevice(device)
			context.JSON(http.StatusCreated, createdPanicDevice)
		} else {
			handleApiError(context, bindingError)
		}
	}
	return routes
}
