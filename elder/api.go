package elder

import (
	"net/http"

	"github.com/gin-gonic/gin"
	inf "github.com/julianespinel/btn-server/infrastructure"
)

type ElderAPI struct {
	elderBusiness ElderBusiness
}

func CreateElderAPI(elderBusiness ElderBusiness) ElderAPI {
	return ElderAPI{elderBusiness: elderBusiness}
}

func (api ElderAPI) CreateElder() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		var elder Elder
		bindingError := context.Bind(&elder)
		if bindingError == nil {
			business := api.elderBusiness
			createdElder, err := business.createElder(elder)
			if err == nil {
				context.JSON(http.StatusCreated, createdElder)
			} else {
				inf.HandleApiError(context, err)
			}
		} else {
			inf.HandleApiError(context, bindingError)
		}
	}
	return handlerFunction
}

func (api ElderAPI) AddRelativeToElder() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		elderId := context.Param("elderId")
		var relative Relative
		bindingError := context.Bind(&relative)
		if bindingError == nil {
			business := api.elderBusiness
			_, err := business.addRelativeToElder(elderId, relative)
			if err == nil {
				stringMessage := inf.GetStringMessage("message", "Relative was attached to elder.")
				context.JSON(http.StatusCreated, stringMessage)
			} else {
				inf.HandleApiError(context, err)
			}
		} else {
			inf.HandleApiError(context, bindingError)
		}
	}
	return handlerFunction
}
