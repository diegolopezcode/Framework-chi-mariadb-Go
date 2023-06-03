package main

import (
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/api/routes"
	"github.com/diegolopezcode/api-crud-complete-chi/configs"
	database "github.com/diegolopezcode/api-crud-complete-chi/models/database/config"
	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/go-chi/render"
)

func main() {

	// Creating a new router.
	app := chi.NewMux()

	// A middleware that is going to be executed before the routes.
	app.Use(middleware.Logger)
	app.Use(middleware.RequestID)
	app.Use(middleware.URLFormat)
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{configs.Config("CORS_ALLOW_ORIGINS")},
		AllowedMethods: []string{configs.Config("CORS_ALLOW_METHODS")},
		AllowedHeaders: []string{configs.Config("CORS_ALLOW_HEADERS")},
	}))
	app.Use(render.SetContentType(render.ContentTypeJSON))
	//Routes
	routes.SetupRoutes(app)
	//Start Database
	database.Connect()

	//Listen
	http.ListenAndServe(":"+configs.Config("PORT"), app)

}
