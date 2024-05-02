package db

import (
	"embed"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//go:embed migrations/*.sql
var migrationFs embed.FS

var Username, Password, Database, Host string
var Port int

func InitDB(db_driver *gorm.DB) {
	sqlDB, err := db_driver.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
}

func GetMysqlDriver(db *DB) *gorm.DB {
	Database = db.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, Database)

	dialect := mysql.Open(dsn)
	gorm_db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return gorm_db
}

func GetMigrate(db_driver database.Driver, migrate_name string) error {
	d, err := iofs.New(migrationFs, "migrations")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance("iofs", d, Database, db_driver)
	if err != nil {
		panic(err)
	}

	switch migrate_name {
	case "down":
		err = m.Down()
	default:
		err = m.Up()
	}

	return err
}
