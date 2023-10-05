package data

import (
	"log"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(db *xorm.Engine) (*Data, func(), error) {
	return &Data{db}, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed to close database: %s", err)
		}
	}, nil
}

func NewDB() (*xorm.Engine, error) {
	e, err := xorm.NewEngine("mysql", "dsn")
	if err != nil {
		return nil, err
	}

	if gin.Mode() == "debug" {
		e.ShowSQL(true)
	}

	if err := e.Ping(); err != nil {
		return nil, err
	}

	return e, err
}
