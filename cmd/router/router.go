package router

import (
	"dododo/domain/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
)

func New(conn *pgx.Conn) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	userControllers := user.Init(conn)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Word"))
	})
	router.Post("/api/v1/register", userControllers.Register)
	router.Post("/api/v1/login", userControllers.Login)
	router.Get("/api/v1/logout", userControllers.Logout)

	return router
}