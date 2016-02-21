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

func (business ElderBusiness) createElder(elder Elder) (Elder, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.elderDAO
	return dao.createElder(database, elder)
}

func (business ElderBusiness) addRelativeToElder(elderId string, relative Relative) (bool, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.elderDAO
	return dao.addRelativeToElder(database, elderId, relative)
}

func (business ElderBusiness) removeRelativeFromElder(elderId string, relativeId string) (bool, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.elderDAO
	return dao.removeRelativeFromElder(database, elderId, relativeId)
}

func (business ElderBusiness) GetElderRelatives(elderId string) ([]Relative, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.elderDAO
	return dao.getElderRelatives(database, elderId)
}

func (business ElderBusiness) GetElderById(elderId string) (Elder, error) {
	database := getDatabase(business.dbConfig)
	defer database.Close()
	dao := business.elderDAO
	return dao.getElderById(database, elderId)
}
