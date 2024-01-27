package router

import (
	"mocky-blog/domain/user"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *chi.Mux {
	router := chi.NewRouter()
	userControllers := user.Init(db)

	router.Use(middleware.Logger)
	
	router.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Word"))
		})
	})

	router.Route("/api/v1", func(routes chi.Router) {
		routes.Group(func(r chi.Router) {
			r.Post("/register", userControllers.Register)
			r.Post("/login", userControllers.Login)
			r.Get("/logout", userControllers.Logout)
		})
	})

	return router
}