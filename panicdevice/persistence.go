package panicdevice

import (
	"database/sql"
	"time"
)

type PanicDAO struct {
	err error
}

func (dao PanicDAO) handleError(err error) {
	daoError := dao.err
	if daoError == nil {
		daoError = err
	}
}

func (dao PanicDAO) createPanicDevice(database *sql.DB, device PanicDevice) (PanicDevice, error) {

	stmt, err := database.Prepare("INSERT INTO panic_devices (serial, birth_date) VALUES (?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(device.Serial, now)
	dao.handleError(err)

	return device, dao.err
}

func (dao PanicDAO) attachElderToPanicDevice(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("INSERT INTO devices_elders (serial, elder_id) VALUES (? ,?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	dao.handleError(err)

	wasSuccessful := false
	daoError := dao.err
	if daoError == nil {
		wasSuccessful, daoError = dao.addElderToPanicDeviceHistory(database, deviceSerial, elderId)
	}

	return wasSuccessful, daoError
}

func (dao PanicDAO) addElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("INSERT INTO devices_history (serial, elder_id, attachment_date) VALUES (?, ?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(deviceSerial, elderId, now)
	dao.handleError(err)

	daoError := dao.err
	wasSuccessful := (daoError == nil)
	return wasSuccessful, daoError
}

func (dao PanicDAO) detachElderFromPanicDevise(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("DELETE FROM devices_elders WHERE serial = ? AND elder_id = ?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	dao.handleError(err)

	wasSuccessful := false
	daoError := dao.err
	if daoError == nil {
		wasSuccessful, daoError = dao.removeElderToPanicDeviceHistory(database, deviceSerial, elderId)
	}

	return wasSuccessful, daoError
}

func (dao PanicDAO) removeElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("UPDATE devices_history SET detachment_date = ? WHERE serial = ? AND elder_id = ?);")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(now, deviceSerial, elderId)
	dao.handleError(err)

	daoError := dao.err
	wasSuccessful := (daoError == nil)
	return wasSuccessful, daoError
}
