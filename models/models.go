package models

import (
	"fmt"
	"path/filepath"

	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"upper.io/db.v3/mongo"
)

// init a database instance
func NewDbEngine() (err error) {
	switch DATA_TYPE {
	case "sqlite":
		cur, _ := filepath.Abs("..")
		dataSourceName := fmt.Sprintf("%v/%v/%v.db", cur, DATA_NAME, DATA_NAME)
		Engine, err = xorm.NewEngine("sqlite3", dataSourceName)
		err = Engine.Ping()

	case "mysql":
		dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8",
			USERNAME, PASSWORD, DATA_HOST, DATA_PORT, DATA_NAME)

		Engine, err = xorm.NewEngine("mysql", dataSourceName)
		err = Engine.Ping()
	case "postgres":
		dbSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", USERNAME, PASSWORD, DATA_HOST,
			DATA_PORT, DATA_NAME, SSL_MODE)
		Engine, err = xorm.NewEngine("postgres", dbSourceName)
		err = Engine.Ping()

	case "mongodb":
		err = NewMongodbClient()
		err = MongodbClient.Ping()

	default:
		cur, _ := filepath.Abs("..")
		dataSourceName := fmt.Sprintf("%v/%v/%v.db", cur, DATA_NAME, DATA_NAME)
		Engine, err = xorm.NewEngine("sqlite3", dataSourceName)
		err = Engine.Ping()
	}

	return err
}

// return a mongodb session
func NewMongodbClient() (err error) {
	setting := mongo.ConnectionURL{Host: fmt.Sprintf("%v:%v", DATA_HOST, DATA_PORT), Database: DATA_NAME, User: USERNAME, Password: PASSWORD}
	MongodbClient, err = mongo.Open(setting)
	return err
}
