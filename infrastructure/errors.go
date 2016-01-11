package infrastructure

import (
	"github.com/Sirupsen/logrus"
)

var log = logrus.New()

func HandleDBError(err error) {
	if err != nil {
		log.Error(err)
	}
}
