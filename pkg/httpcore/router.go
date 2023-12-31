package httpcore

import (
	"skytrace/pkg/util"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func NewRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(cors.AllowAll().Handler)
	router.Use(LoggerMiddleware(&util.RichLog))
	router.Use(middleware.Recoverer)

	return router
}
