package elder

import (
	"database/sql"
	"time"

	inf "github.com/julianespinel/btn-server/infrastructure"
)

type ElderDAO struct {
}

func (dao ElderDAO) createElder(database *sql.DB, elder Elder) Elder {
	stmt, err := database.Prepare("INSERT INTO elders (id, name, last_name, cellphone, registration_date) VALUES (?, ?, ?, ? ,?);")
	inf.HandleDBError(err)
	defer stmt.Close()

	now := time.Now()
	_, err = stmt.Exec(elder.Id, elder.Name, elder.LastName, elder.Cellphone, now)
	inf.HandleDBError(err)

	return elder
}
