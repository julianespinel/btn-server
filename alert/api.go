package alert

import (
	"errors"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	el "github.com/julianespinel/btn-server/elder"
	inf "github.com/julianespinel/btn-server/infrastructure"
	pa "github.com/julianespinel/btn-server/panicdevice"
)

type AlertAPI struct {
	errorCode     int
	errorMessage  error
	alertBusiness AlertBusiness
	panicBusiness pa.PanicBusiness
	elderBusiness el.ElderBusiness
}

func CreateAlertAPI(alertBusiness AlertBusiness, panicBusiness pa.PanicBusiness, elderBusiness el.ElderBusiness) AlertAPI {
	return AlertAPI{errorCode: -1, errorMessage: nil, alertBusiness: alertBusiness, panicBusiness: panicBusiness, elderBusiness: elderBusiness}
}

var log = logrus.New()

func (api *AlertAPI) handleError(errCode int, err error) {
	if api.errorMessage == nil && err != nil {
		api.errorCode = errCode
		api.errorMessage = err
		log.Error(err)
	}
}

func (api *AlertAPI) error() error {
	apiError := api.errorMessage
	api.errorMessage = nil
	return apiError
}

func (api AlertAPI) CreateAlert() gin.HandlerFunc {
	handlerFunction := func(context *gin.Context) {
		serial := context.Query("serial")
		panicBusiness := api.panicBusiness
		elderId, err := panicBusiness.GetElderIdAssignedToPanicDevice(serial)
		api.handleError(500, err)

		var elder el.Elder
		elderBusiness := api.elderBusiness
		if elderId != "" {
			elder, err = elderBusiness.GetElderById(elderId)
			api.handleError(500, err)
		} else {
			err = errors.New("The panic device serial does not exist or does not have an associated elder.")
			api.handleError(404, err)
		}

		elderRelatives := []el.Relative{}
		if elder.Id != "" {
			elderRelatives, err = elderBusiness.GetElderRelatives(elderId)
			api.handleError(500, err)
		} else {
			err = errors.New("The elder with id " + elderId + " does not exist.")
			api.handleError(404, err)
		}

		var sendingResults []SendingResult
		if len(elderRelatives) > 0 {
			alert := Alert{Serial: serial, ElderId: elderId, Date: time.Now()}
			alertBusiness := api.alertBusiness
			sendingResults, err = alertBusiness.processAlert(alert, elder, elderRelatives)
			api.handleError(500, err)
		} else {
			err = errors.New("The elder with id " + elderId + " does not have any relatives.")
			api.handleError(404, err)
		}

		apiError := api.error()
		if apiError == nil {
			context.JSON(http.StatusCreated, sendingResults)
		} else {
			inf.HandleApiErrorWithStatusCode(context, api.errorCode, apiError)
		}
	}
	return handlerFunction
}
