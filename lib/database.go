package lib

import (
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	conn, err := leveldb.OpenFile("balance.db", nil)
	if err != nil {
		logrus.Errorf("Error connection db: %v", err)
	}

	db = conn
}

func DatabaseConnection() *leveldb.DB {
	return db
}
