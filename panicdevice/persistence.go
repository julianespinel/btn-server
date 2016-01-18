package panicdevice

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
)

type PanicDAO struct {
	err error
}

var log = logrus.New()

// Keep the first error.
func (dao PanicDAO) handleError(err error) {
	if dao.err == nil {
		dao.err = err
		log.Error(err)
	}
}

// clean dao.err and return the latest error if any.
func (dao PanicDAO) Error() error {
	daoError := dao.err
	dao.err = nil
	return daoError
}

func (dao PanicDAO) createPanicDevice(database *sql.DB, device PanicDevice) (PanicDevice, error) {

	stmt, err := database.Prepare("INSERT INTO panic_devices (serial, birth_date) VALUES (?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(device.Serial, now)
	dao.handleError(err)

	return device, dao.Error()
}

func (dao PanicDAO) attachElderToPanicDevice(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("INSERT INTO devices_elders (serial, elder_id) VALUES (? ,?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	dao.handleError(err)

	wasSuccessful := false
	wasSuccessful, err = dao.addElderToPanicDeviceHistory(database, deviceSerial, elderId)
	dao.handleError(err)

	return wasSuccessful, dao.Error()
}

func (dao PanicDAO) addElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("INSERT INTO devices_history (serial, elder_id, attachment_date) VALUES (?, ?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(deviceSerial, elderId, now)
	dao.handleError(err)

	daoError := dao.Error()
	wasSuccessful := (daoError == nil)
	return wasSuccessful, daoError
}

func (dao PanicDAO) detachElderFromPanicDevice(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("DELETE FROM devices_elders WHERE serial = ? AND elder_id = ?;")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	dao.handleError(err)

	wasSuccessful := false
	wasSuccessful, err = dao.updateElderFromPanicDeviceHistory(database, deviceSerial, elderId)
	dao.handleError(err)

	return wasSuccessful, dao.Error()
}

func (dao PanicDAO) updateElderFromPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) (bool, error) {

	stmt, err := database.Prepare("UPDATE devices_history SET detachment_date = ? WHERE serial = ? AND elder_id = ?;")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(now, deviceSerial, elderId)
	dao.handleError(err)

	daoError := dao.Error()
	wasSuccessful := (daoError == nil)
	return wasSuccessful, daoError
}
