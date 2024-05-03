package config

import "github.com/incrusio21/nikahmi/db"

type App struct {
	Name string
}

type Config struct {
	App     *App
	DB      *db.DB
	Session string
}

type Session struct {
	Name  string
	Value string
}
