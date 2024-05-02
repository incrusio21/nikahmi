package main

import (
	"flag"

	"github.com/incrusio21/nikahmi/app"
	"github.com/incrusio21/nikahmi/db/mysql"
)

func main() {
	migrate := flag.String("migrate", "", "Menjalankan Migrate dan bukan Server")

	flag.Parse()

	if *migrate != "" {
		mysql.Migrate(*migrate)
		return
	}

	router := app.Router
	Routing(router)

	err := router.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
