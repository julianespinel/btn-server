package panicdevice

import (
	"database/sql"
	"time"

	inf "github.com/julianespinel/btn-server/infrastructure"
)

type PanicDAO struct {
}

func (dao PanicDAO) createPanicDevice(database *sql.DB, device PanicDevice) PanicDevice {

	stmt, err := database.Prepare("INSERT INTO panic_devices (serial, birth_date) VALUES (?, ?);")
	inf.HandleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(device.Serial, now)
	inf.HandleDBError(err)

	return device
}

func attachElderToPanicDevice(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("INSERT INTO devices_elders (serial, elder_id) VALUES (? ,?);")
	inf.HandleDBError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	inf.HandleDBError(err)

	addElderToPanicDeviceHistory(database, deviceSerial, elderId)
}

func addElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("INSERT INTO devices_history (serial, elder_id, attachment_date) VALUES (?, ?, ?);")
	inf.HandleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(deviceSerial, elderId, now)
	inf.HandleDBError(err)
}

func detachElderFromPanicDevise(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("DELETE FROM devices_elders WHERE serial = ? AND elder_id = ?);")
	inf.HandleDBError(err)
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	inf.HandleDBError(err)

	removeElderToPanicDeviceHistory(database, deviceSerial, elderId)
}

func removeElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("UPDATE devices_history SET detachment_date = ? WHERE serial = ? AND elder_id = ?);")
	inf.HandleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(now, deviceSerial, elderId)
	inf.HandleDBError(err)
}
