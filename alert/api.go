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
		if err == nil {
			elderBusiness := api.elderBusiness
			elder, err := elderBusiness.GetElderById(elderId)
			if err == nil {
				elderBusiness := api.elderBusiness
				elderRelatives, err := elderBusiness.GetElderRelatives(elderId)
				if err == nil {
					alertBusiness := api.alertBusiness
					alert := Alert{Serial: serial, ElderId: elderId, Date: time.Now()}
					sendingResults, err := alertBusiness.processAlert(alert, elder, elderRelatives)
					if err == nil {
						context.JSON(http.StatusCreated, sendingResults)
					} else {
						inf.HandleApiError(context, err)
					}
				} else {
					inf.HandleApiError(context, err)
				}
			} else {
				inf.HandleApiError(context, err)
			}
		} else {
			inf.HandleApiError(context, err)
		}
	}
	return handlerFunction
}
