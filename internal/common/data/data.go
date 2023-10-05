package data

import (
	"fmt"

	"github.com/3lur/go-mall/pkg/config"
	"github.com/3lur/go-mall/pkg/console"
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
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	console.Success("successfully connected to the database!")

	return &Data{db}, func() {
		if err := db.Close(); err != nil {
			console.ExitIf(err)
		}
	}, nil
}
