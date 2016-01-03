package panicdevice

import (
	"time"
	"database/sql"
)

type PanicDAO struct {
}

func createPanicDevice(database *sql.DB, device PanicDevice) PanicDevice {

	stmt, err := database.Prepare("INSERT INTO panic_devices (serial, birth_date) VALUES (?, ?);")
	if err != nil {
		// Handle error properly.
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(device.serial, now)
	if err != nil {
		// Handle error properly.
	}

	return device
}

func attachElderToPanicDevice(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("INSERT INTO devices_elders (serial, elder_id) VALUES (? ,?);")
	if err != nil {
		//
	}
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	if err != nil {
		//
	}

	addElderToPanicDeviceHistory(database, deviceSerial, elderId)
}

func addElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("INSERT INTO devices_history (serial, elder_id, attachment_date) VALUES (?, ?, ?);")
	if err != nil {
		//
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(deviceSerial, elderId, now)
	if err != nil {
		//
	}
}

func detachElderFromPanicDevise(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("DELETE FROM devices_elders WHERE serial = ? AND elder_id = ?);")
	if err != nil {
		//
	}
	defer stmt.Close()

	_, err = stmt.Exec(deviceSerial, elderId)
	if err != nil {
		//
	}

	removeElderToPanicDeviceHistory(database, deviceSerial, elderId)
}

func removeElderToPanicDeviceHistory(database *sql.DB, deviceSerial string, elderId string) {

	stmt, err := database.Prepare("UPDATE devices_history SET detachment_date = ? WHERE serial = ? AND elder_id = ?);")
	if err != nil {
		//
	}
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(now, deviceSerial, elderId)
	if err != nil {
		//
	}
}
