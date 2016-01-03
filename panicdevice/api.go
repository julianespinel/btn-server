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

func (api PanicAPI) PanicRoutes() gin.HandlerFunc {
	routes := func(context *gin.Context) {
		var device PanicDevice
		bindingError := context.Bind(&device)
		if bindingError == nil {
			business := api.panicBusiness
			createdPanicDevice := business.createPanicDevice(device)
			context.JSON(http.StatusCreated, createdPanicDevice)
		} else {
			context.JSON(http.StatusBadRequest, bindingError)
		}
	}
	return routes
}
