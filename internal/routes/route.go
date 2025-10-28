package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/masfuulaji/store/internal/app/handlers"
	"github.com/masfuulaji/store/internal/database"
)

func SetupRoutes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	db, err := database.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	userHandler := handlers.NewUserHandler(db.DB)
	r.Route("/user", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.CreateUser)
		r.Put("/{id}", userHandler.UpdateUser)
		r.Delete("/{id}", userHandler.DeleteUser)
	})

	loginHandler := handlers.NewLoginHandler(db.DB)
	r.Route("/auth", func(r chi.Router) {
		r.Post("/", loginHandler.Login)
		r.Get("/logout", loginHandler.Logout)
	})
}
