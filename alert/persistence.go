package alert

import "database/sql"

type AlertDAO struct {
	err error
}

func CreateAlertDAO() AlertDAO {
	return AlertDAO{err: nil}
}

// Keep the first error.
func (dao *AlertDAO) handleError(newError error) {
	if dao.err == nil && newError != nil {
		dao.err = newError
		log.Error(dao.err)
	}
}

// Clean dao.err and return the latest error if any.
func (dao *AlertDAO) error() error {
	daoError := dao.err
	dao.err = nil
	return daoError
}

func (dao AlertDAO) registerAlert(database *sql.DB, alert Alert) (Alert, error) {
	stmt, err := database.Prepare("INSERT INTO alerts (serial, elder_id, latitude, longitude, date) VALUES (?, ?, ?, ?, ?);")
	dao.handleError(err)
	defer stmt.Close()
	result, err := stmt.Exec(alert.Serial, alert.ElderId, alert.Latitude, alert.Longitude, alert.Date)
	dao.handleError(err)
	var createdAlert Alert
	if dao.err == nil {
		generatedId, err := result.LastInsertId()
		dao.handleError(err)
		if dao.err == nil {
			createdAlert = Alert{float64(generatedId), alert.Serial, alert.ElderId, alert.Latitude, alert.Longitude, alert.Date}
		}
	}
	return createdAlert, dao.error()
}

func (dao AlertDAO) saveSendingResult(database *sql.DB, sendingResult SendingResult) (SendingResult, error) {
	stmt, err := database.Prepare("INSERT INTO sending_results (alert_id, relative_id, was_successful, result_message) VALUES (?, ?, ?, ?);")
	dao.handleError(err)
	defer stmt.Close()
	result, err := stmt.Exec(sendingResult.AlertId, sendingResult.RelativeId, sendingResult.WasSuccessful, sendingResult.ResultMessage)
	dao.handleError(err)
	var savedSendingResult SendingResult
	if dao.err == nil {
		generatedId, err := result.LastInsertId()
		dao.handleError(err)
		if dao.err == nil {
			savedSendingResult = SendingResult{float64(generatedId), sendingResult.AlertId, sendingResult.RelativeId,
				sendingResult.WasSuccessful, sendingResult.ResultMessage}
		}
	}
	return savedSendingResult, dao.error()
}
