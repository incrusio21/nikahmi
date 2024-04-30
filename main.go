package main

import (
	"embed"
	"flag"

	"github.com/incrusio21/nikahmi/db/mysql"
)

//go:embed app/*
var Templatesfs embed.FS

func main() {
	migrate := flag.String("migrate", "", "Menjalankan Migrate dan bukan Server")

	flag.Parse()

	if *migrate != "" {
		mysql.Migrate(*migrate)
		return
	}
}
