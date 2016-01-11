package elder

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	inf "github.com/julianespinel/btn-server/infrastructure"
)

type ElderBusiness struct {
	dbConfig inf.DBConfig
	elderDAO ElderDAO
}

func CreateElderBusiness(dbConfig inf.DBConfig, elderDAO ElderDAO) ElderBusiness {
	return ElderBusiness{dbConfig: dbConfig, elderDAO: elderDAO}
}

func getDatabase(dbConfig inf.DBConfig) *sql.DB {
	database, err := sql.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@/"+dbConfig.DbName)
	inf.HandleDBError(err)
	return database
}

func (elderBusiness ElderBusiness) createElder(elder Elder) Elder {
	database := getDatabase(elderBusiness.dbConfig)
	defer database.Close()
	dao := elderBusiness.elderDAO
	return dao.createElder(database, elder)
}
