package elder

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
)

type ElderDAO struct {
	err error
}

func CreateElderDAO() ElderDAO {
	return ElderDAO{err: nil}
}

var log = logrus.New()

// Keep the first error.
func (dao *ElderDAO) handleError(err error) {
	if dao.err == nil && err != nil {
		dao.err = err
		log.Error(err)
	}
}

// Clean dao.err and return the latest error if any.
func (dao *ElderDAO) error() error {
	daoError := dao.err
	dao.err = nil
	return daoError
}

func (dao ElderDAO) createElder(database *sql.DB, elder Elder) (Elder, error) {

	stmt, err := database.Prepare("INSERT INTO elders (id, name, last_name, cellphone, registration_date) VALUES (?, ?, ?, ? ,?);")
	dao.handleError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(elder.Id, elder.Name, elder.LastName, elder.Cellphone, now)
	dao.handleError(err)

	return elder, dao.error()
}

func (dao ElderDAO) createRelative(database *sql.DB, relative Relative) (Relative, error) {

	stmt, err := database.Prepare("INSERT INTO relatives (id, name, last_name, email, cellphone, relationship) VALUES (?, ?, ?, ?, ?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(relative.Id, relative.Name, relative.LastName, relative.Email, relative.Cellphone, relative.Relationship)
	dao.handleError(err)

	return relative, dao.error()
}

func (dao ElderDAO) addRelativeToElder(database *sql.DB, elderId string, relative Relative) (bool, error) {

	_, err := dao.createRelative(database, relative)
	dao.handleError(err)

	stmt, err := database.Prepare("INSERT INTO elders_relatives (elder_id, relative_id) VALUES (?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(elderId, relative.Id)
	dao.handleError(err)

	daoError := dao.error()
	relativeWasAdded := (daoError == nil)

	return relativeWasAdded, daoError
}

func (dao ElderDAO) removeRelativeFromElder(database *sql.DB, elderId string, relativeId string) (bool, error) {
	stmt, err := database.Prepare("DELETE FROM elders_relatives WHERE elder_id=? AND relative_id=?;")
	dao.handleError(err)
	defer stmt.Close()
	_, err = stmt.Exec(elderId, relativeId)
	dao.handleError(err)
	daoError := dao.error()
	relativeWasRemoved := (daoError == nil)
	return relativeWasRemoved, daoError
}

func (dao ElderDAO) getElderRelatives(database *sql.DB, elderId string) ([]Relative, error) {
	relatives := []Relative{}
	stmt, err := database.Prepare("SELECT r.* FROM elders_relatives er, relatives r WHERE er.elder_id = ? AND er.relative_id = r.id;")
	dao.handleError(err)
	defer stmt.Close()
	rows, err := stmt.Query(elderId)
	dao.handleError(err)
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var last_name string
		var email string
		var cellphone string
		var relationship string
		err = rows.Scan(&id, &name, &last_name, &email, &cellphone, &relationship)
		dao.handleError(err)
		relative := Relative{id, name, last_name, email, cellphone, relationship}
		relatives = append(relatives, relative)
	}
	err = rows.Err()
	dao.handleError(err)
	return relatives, dao.error()
}

func (dao ElderDAO) getElderById(database *sql.DB, elderId string) (Elder, error) {
	var elder Elder
	stmt, err := database.Prepare("SELECT * FROM elders WHERE id = ?;")
	dao.handleError(err)
	defer stmt.Close()
	rows, err := stmt.Query(elderId)
	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var last_name string
		var cellphone string
		var registration_date time.Time
		err = rows.Scan(&id, &name, &last_name, &cellphone, &registration_date)
		dao.handleError(err)
		elder = Elder{id, name, last_name, cellphone, registration_date}
	}
	err = rows.Err()
	dao.handleError(err)
	return elder, dao.error()
}
