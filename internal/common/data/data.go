package data

import (
	"fmt"
	"log"

	"github.com/3lur/go-mall/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(conf *config.Config) (*Data, func(), error) {
	// "root:123@/test?charset=utf8"
	// root:123456@tcp(127.0.0.1:3380)/user?charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci",
		conf.Database.Username,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.DbName,
	)
	db, err := xorm.NewEngine(conf.Database.Driver, dsn)

	if err != nil {
		log.Fatalf("failed connect to database: %s", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("database ping error: %s", err)
	}

	log.Println("successfully connected to the database!")

	return &Data{db}, func() {
		if err := db.Close(); err != nil {
			log.Fatalf("failed to close database: %s", err)
		}
	}, nil
}
