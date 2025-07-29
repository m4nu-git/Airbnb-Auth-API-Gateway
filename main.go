package main

import (
	"AuthInGo/app"
	dbConfig "AuthInGo/config/db"
)

func main() {

	// cfg := app.Config {
	// 	Addr: ":3001",
	// }

	// app := app.Application {
	// 	Config: cfg,
	// }



	cfg := app.NewConfig() // Set the server to listen on port 8080
	app := app.NewApplication(cfg)
	dbConfig.SetupDB()
	app.Run()
}