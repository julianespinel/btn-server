package alert

import (
	"database/sql"

	twilio "github.com/carlosdp/twiliogo"
	el "github.com/julianespinel/btn-server/elder"
	inf "github.com/julianespinel/btn-server/infrastructure"
)

type AlertBusiness struct {
	smsConfig inf.SmsConfig
	dbConfig  inf.DBConfig
	alertDAO  AlertDAO
}

func CreateAlertBusiness(smsConfig inf.SmsConfig, dbConfig inf.DBConfig, alertDAO AlertDAO) AlertBusiness {
	return AlertBusiness{smsConfig: smsConfig, dbConfig: dbConfig, alertDAO: alertDAO}
}

func getDatabase(dbConfig inf.DBConfig) *sql.DB {
	database, err := sql.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@/"+dbConfig.DbName)
	inf.HandleDBError(err)
	return database
}

func (business AlertBusiness) registerAlert(alert Alert) (Alert, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.alertDAO
	return dao.registerAlert(database, alert)
}

func (business AlertBusiness) saveSendingResult(sendingResult SendingResult) (SendingResult, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.alertDAO
	return dao.saveSendingResult(database, sendingResult)
}

func (business AlertBusiness) sendAlert(alert Alert, elder el.Elder, elderRelatives []el.Relative) ([]SendingResult, error) {
	sendingResults := []SendingResult{}
	var resultError error
	accountSID := business.smsConfig.AccountSID // "AC21b7beef6dcc64ba200ac17e4b5f00b4"
	authToken := business.smsConfig.AuthToken   //"3bd1d774aca57afc183d2ee5dfd174b9"
	twClient := twilio.NewClient(accountSID, authToken)
	from := business.smsConfig.FromNumber // "+12016270931"
	for _, relative := range elderRelatives {
		to := relative.Cellphone
		message := twilio.Body("btn-project: hello, your " + relative.Relationship + " " + elder.Name + " " +
			elder.LastName + " just pressed the panic button.")
		_, resultError := twilio.NewMessage(twClient, from, to, message)
		wasSuccessful := (resultError == nil)
		resultMessage := ""
		if !wasSuccessful {
			resultMessage = resultError.Error()
		}
		notSavedSendingResult := SendingResult{AlertId: alert.Id, RelativeId: relative.Id,
			WasSuccessful: wasSuccessful, ResultMessage: resultMessage}
		savedSendingResult, resultError := business.saveSendingResult(notSavedSendingResult)
		if resultError == nil {
			sendingResults = append(sendingResults, savedSendingResult)
		} else {
			sendingResults = append(sendingResults, notSavedSendingResult)
		}
	}

	return sendingResults, resultError
}

func (business AlertBusiness) processAlert(alert Alert, elder el.Elder, elderRelatives []el.Relative) ([]SendingResult, error) {
	var sendingResults []SendingResult
	createdAlert, err := business.registerAlert(alert)
	if err == nil {
		sendingResults, err = business.sendAlert(createdAlert, elder, elderRelatives)
	}
	return sendingResults, err
}
