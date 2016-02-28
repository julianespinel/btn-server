package alert

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	el "github.com/julianespinel/btn-server/elder"
	inf "github.com/julianespinel/btn-server/infrastructure"
	pa "github.com/julianespinel/btn-server/panicdevice"
)

type AlertAPI struct {
	alertBusiness AlertBusiness
	panicBusiness pa.PanicBusiness
	elderBusiness el.ElderBusiness
}

func CreateAlertAPI(alertBusiness AlertBusiness, panicBusiness pa.PanicBusiness, elderBusiness el.ElderBusiness) AlertAPI {
	return AlertAPI{alertBusiness: alertBusiness, panicBusiness: panicBusiness, elderBusiness: elderBusiness}
}

func (api AlertAPI) CreateAlert() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		serial := context.Query("serial")
		panicBusiness := api.panicBusiness
		elderId, err := panicBusiness.GetElderIdAssignedToPanicDevice(serial)
		inf.HandleOptionalApiError(context, err)
		elderBusiness := api.elderBusiness
		elder, err := elderBusiness.GetElderById(elderId)
		inf.HandleOptionalApiError(context, err)
		elderRelatives, err := elderBusiness.GetElderRelatives(elderId)
		inf.HandleOptionalApiError(context, err)
		alert := Alert{Serial: serial, ElderId: elderId, Date: time.Now()}
		alertBusiness := api.alertBusiness
		sendingResults, err := alertBusiness.processAlert(alert, elder, elderRelatives)
		inf.HandleOptionalApiError(context, err)
		context.JSON(http.StatusCreated, sendingResults)
	}
	return handlerFunction
}
