package routes

import (
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/api/paths"
	"github.com/diegolopezcode/api-crud-complete-chi/configs"
	"github.com/diegolopezcode/api-crud-complete-chi/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

func SetupRoutes(app *chi.Mux) {

	// Creating a route that is accessible to everyone.
	app.Route(paths.PUBLIC, func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hi"))
		})

		r.Post("/createrole", handler.CreateRole)
		r.Post("/login", handler.Login)
		r.Get("/getrole", handler.GetRoleById)
		r.Put("/updaterole", handler.UpdateRole)
		r.Post("/createpermission", handler.CreatePermission)
		r.Get("/getpermission", handler.GetPermissionById)
		r.Put("/updatepermission", handler.UpdatePermission)

	})

	// Creating a route that is only accessible if the user is logged in.
	app.Route(paths.PRIVATE, func(r chi.Router) {

		r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(configs.Config("JWT_SECRET")), nil)))
		r.Use(jwtauth.Authenticator)

		r.Get("/{paths}", handler.Indexar)

	})
}
