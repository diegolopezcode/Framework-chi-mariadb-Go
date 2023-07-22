package routes

import (
	"net/http"

	"github.com/diegolopezcode/api-crud-complete-chi/api/paths"
	"github.com/diegolopezcode/api-crud-complete-chi/configs"
	Auth "github.com/diegolopezcode/api-crud-complete-chi/handler/login"
	HandlerPermissions "github.com/diegolopezcode/api-crud-complete-chi/handler/permissions"
	HandlerRolePermissions "github.com/diegolopezcode/api-crud-complete-chi/handler/role_permissions"
	HandlerRoles "github.com/diegolopezcode/api-crud-complete-chi/handler/roles"
	HandlerTasks "github.com/diegolopezcode/api-crud-complete-chi/handler/tasks"
	HandlerUsers "github.com/diegolopezcode/api-crud-complete-chi/handler/users"
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

		r.Post("/createrole", HandlerRoles.CreateRole)
		r.Post("/login", Auth.Login)
		r.Get("/getrole", HandlerRoles.GetRoleById)
		r.Put("/updaterole", HandlerRoles.UpdateRole)
		r.Post("/createpermission", HandlerPermissions.CreatePermission)
		r.Get("/getpermission", HandlerPermissions.GetPermissionById)
		r.Put("/updatepermission", HandlerPermissions.UpdatePermission)
		r.Post("/createrolepermission", HandlerRolePermissions.CreateRolePermission)
		r.Get("/getrolepermission", HandlerRolePermissions.GetRolePermissionById)
		r.Post("/createuser", HandlerUsers.CreateUser)
		r.Get("/getuser", HandlerUsers.GetUsers)
		r.Patch("/updateuser", HandlerUsers.UpdateUser)
		r.Post("/createtask", HandlerTasks.CreateTask)
		r.Get("/gettask", HandlerTasks.GetTasks)
		r.Patch("/updatetask", HandlerTasks.UpdateTask)
		r.Delete("/deletetask", HandlerTasks.DeleteTask)

	})

	// Creating a route that is only accessible if the user is logged in.
	app.Route(paths.PRIVATE, func(r chi.Router) {

		r.Use(jwtauth.Verifier(jwtauth.New("HS256", []byte(configs.Config("JWT_SECRET")), nil)))
		r.Use(jwtauth.Authenticator)

		// r.Get("/{paths}", handler.Indexar)

	})
}
