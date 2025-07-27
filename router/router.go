package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

func SetRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Get("/ping", controllers.PingHandler)

	return router
}