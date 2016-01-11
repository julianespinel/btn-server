package panicdevice

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

type PanicDAO struct {
}

func handleDBError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func createPanicDevice(database *sql.DB, device PanicDevice) PanicDevice {

	stmt, err := database.Prepare("INSERT INTO panic_devices (serial, birth_date) VALUES (?, ?);")
	handleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(device.Serial, now)
	handleDBError(err)

	return device
}

func attachElderToPanicDevice(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("INSERT INTO devices_elders (serial, elder_id) VALUES (? ,?);")
	handleDBError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	handleDBError(err)

	addElderToPanicDeviceHistory(database, deviceSerial, elderId)
}

func addElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("INSERT INTO devices_history (serial, elder_id, attachment_date) VALUES (?, ?, ?);")
	handleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(deviceSerial, elderId, now)
	handleDBError(err)
}

func detachElderFromPanicDevise(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("DELETE FROM devices_elders WHERE serial = ? AND elder_id = ?);")
	handleDBError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	handleDBError(err)

	removeElderToPanicDeviceHistory(database, deviceSerial, elderId)
}

func removeElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("UPDATE devices_history SET detachment_date = ? WHERE serial = ? AND elder_id = ?);")
	handleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(now, deviceSerial, elderId)
	handleDBError(err)
}
