package db

import (
	"embed"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/incrusio21/nikahmi/config"
	"gorm.io/gorm"
)

//go:embed migrations/*.sql
var migrationFs embed.FS

var Username, Password, Database, Host string
var Port int

func init() {
	conf, err := config.Read()
	if err != nil {
		panic(err)
	}
	Username = conf.DB.Username
	Password = conf.DB.Password
	Database = conf.DB.Database
	Host = conf.DB.Host
	Port = conf.DB.Port
}

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
