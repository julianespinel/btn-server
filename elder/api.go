package elder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ElderAPI struct {
	elderBusiness ElderBusiness
}

func CreateElderAPI(elderBusiness ElderBusiness) ElderAPI {
	return ElderAPI{elderBusiness: elderBusiness}
}

func handleApiError(context *gin.Context, err error) {
	context.JSON(-1, context.Error(err)) // -1 == not override the current error code
}

func (api ElderAPI) CreateElder() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		var elder Elder
		bindingError := context.Bind(&elder)
		if bindingError == nil {
			business := api.elderBusiness
			createdElder := business.createElder(elder)
			context.JSON(http.StatusCreated, createdElder)
		} else {
			handleApiError(context, bindingError)
		}
	}
	return handlerFunction
}
