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
func (dao ElderDAO) handleError(err error) {
	if dao.err == nil && err != nil {
		dao.err = err
		log.Error(err)
	}
}

// Clean dao.err and return the latest error if any.
func (dao ElderDAO) Error() error {
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

	return elder, dao.Error()
}

func (dao ElderDAO) createRelative(database *sql.DB, relative Relative) (Relative, error) {

	stmt, err := database.Prepare("INSERT INTO relatives (id, name, last_name, email, cellphone, relationship) VALUES (?, ?, ?, ?, ?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(relative.Id, relative.Name, relative.LastName, relative.Email, relative.Cellphone, relative.Relationship)
	dao.handleError(err)

	return relative, dao.Error()
}

func (dao ElderDAO) addRelativeToElder(database *sql.DB, elderId string, relative Relative) (bool, error) {

	_, err := dao.createRelative(database, relative)
	dao.handleError(err)

	stmt, err := database.Prepare("INSERT INTO elders_relatives (elder_id, relative_id) VALUES (?, ?);")
	dao.handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(elderId, relative.Id)
	dao.handleError(err)

	daoError := dao.Error()
	relativeWasAdded := (daoError == nil)

	return relativeWasAdded, daoError
}
