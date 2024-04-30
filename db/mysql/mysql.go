package mysql

import (
	"fmt"

	mysql_migrate "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/incrusio21/nikahmi/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dsn string
var Db *gorm.DB

func init() {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, db.Database)
	Db = GetMysqlDriver()
	db.InitDB(Db)
}

func GetMysqlDriver() *gorm.DB {
	dialect := mysql.Open(dsn)
	gorm_db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return gorm_db
}

func Migrate(migrate_name string) error {
	sqlDB, err := Db.DB()
	if err != nil {
		panic(err)
	}

	driver, _ := mysql_migrate.WithInstance(sqlDB, &mysql_migrate.Config{})

	return db.GetMigrate(driver, migrate_name)
}
