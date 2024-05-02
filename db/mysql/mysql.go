package mysql

import (
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/incrusio21/nikahmi/config"
	"github.com/incrusio21/nikahmi/db"
)

var Db = config.Database

func init() {
	db.InitDB(Db)
}

func Migrate(migrate_name string) error {
	sqlDB := config.MysqlDB()

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	return db.GetMigrate(driver, migrate_name)
}
