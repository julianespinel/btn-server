package elder

import (
	"database/sql"
	"time"

	"github.com/Sirupsen/logrus"
)

type ElderDAO struct {
	err error
}

var log = logrus.New()

// Keep the first error.
func (dao ElderDAO) handleError(err error) {
	if dao.err == nil {
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
