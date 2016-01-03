package panicdevice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	inf "github.com/julianespinel/btn-server/infrastructure"
)

type PanicBusiness struct {
	dbConfig inf.DBConfig
	panicDAO PanicDAO
}

func CreatePanicBusiness(dbConfig inf.DBConfig, panicDAO PanicDAO) PanicBusiness {
	return PanicBusiness{dbConfig: dbConfig, panicDAO: panicDAO}
}

func getDatabase(dbConfig inf.DBConfig) *sql.DB {
	database, err := sql.Open("mysql", dbConfig.Username+":"+dbConfig.Password+"@/"+dbConfig.DbName)
	if err != nil {
		// Handle error properly.
	}
	return database
}

func (business PanicBusiness) createPanicDevice(device PanicDevice) PanicDevice {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	return createPanicDevice(database, device)
}
