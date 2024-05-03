package config

import (
	"bytes"
	"database/sql"
	_ "embed"

	"github.com/incrusio21/nikahmi/app/middleware/session"
	"github.com/incrusio21/nikahmi/config/storage"
	"github.com/incrusio21/nikahmi/config/storage/memory"
	"github.com/incrusio21/nikahmi/config/storage/mysql"
	"github.com/incrusio21/nikahmi/db"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//go:embed .yaml
var defaultConfiguration []byte
var store *session.Store
var Database *gorm.DB
var Yaml *Config

var viper_config = viper.New()

func init() {
	// Configuration file
	viper_config.SetConfigType("yaml")

	// Read configuration
	if err := viper_config.ReadConfig(bytes.NewBuffer(defaultConfiguration)); err != nil {
		panic(err)
	}

	conf, err := Read()
	if err != nil {
		panic(err)
	}

	Yaml = conf

	switch conf.DB.Driver {
	default:
		Database = db.GetMysqlDriver(conf.DB)
	}

	var storage storage.Storage
	switch conf.Session {
	case "mysql":
		var db_session *sql.DB
		if conf.DB.Driver == "mysql" {
			db_session = MysqlDB()
		} else {
			panic("Maaf Driver mysql tidak dapat di gunakan untuk Session")
		}

		storage = mysql.New(mysql.Config{
			Db:    db_session,
			Reset: false,
			// GCInterval:      10 * time.Second,
		})
	default:
		storage = memory.New()
	}

	store = session.New(session.Config{
		Storage: storage,
	})
}

func MysqlDB() *sql.DB {
	sqlDB, err := Database.DB()
	if err != nil {
		panic(err)
	}

	return sqlDB
}

func Read() (*Config, error) {
	var config Config

	// Unmarshal the configuration
	if err := viper_config.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
