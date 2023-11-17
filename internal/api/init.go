package api

import (
	"skytrace/pkg/httpcore"
	"skytrace/pkg/util"

	"github.com/go-chi/chi/v5"
)

func InitService() chi.Router {
	util.InitLogger()
	router := httpcore.NewRouter()
	//env := todoservice.InitEnv() // get typed environment
	controller := NewTodoController()
	ApplyRoutes(router, controller)

	return router
}
