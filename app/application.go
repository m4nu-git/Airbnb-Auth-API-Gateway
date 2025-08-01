package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
	"AuthInGo/router"
	"AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

// Config holds the configuration for the server
type Config struct {
	Addr string //PORT
}

type Application struct {
	Config Config
}

// Constructor for config
func NewConfig() Config {

	port := config.GetString("PORT", ":8080")
	
	return Config{
		Addr: port,
	}
}

// Constructor for Application
func NewApplication(cfg Config) *Application {
	return &Application {
		Config: cfg,
	}
}

func (app *Application) Run() error {

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error settings up database:", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	us := services.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(uc)

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetRouter(uRouter), // TODO: Setup a chi router and put it here
		ReadTimeout: 10 * time.Second, // Set read timeout to 10 seconds
		WriteTimeout: 10 * time.Second, // Set write timeout to 10 seconds
	}

	fmt.Println("Starting server on", app.Config.Addr)

	return server.ListenAndServe()
}