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
	r.Route("/profile", func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.Get("/", userHandler.GetUser)
		// r.Get("/{id}", userHandler.GetUser)
		// r.Post("/", userHandler.CreateUser)
		// r.Put("/{id}", userHandler.UpdateUser)
		// r.Delete("/{id}", userHandler.DeleteUser)
		r.Put("/update", userHandler.UpdateUser)
		r.Put("/image", userHandler.UpdateUserProfile)
	})

	r.Post("/registration", userHandler.CreateUser)

	loginHandler := handlers.NewLoginHandler(db.DB)
	// r.Route("/auth", func(r chi.Router) {
	// 	r.Post("/", loginHandler.Login)
	// 	r.Get("/logout", loginHandler.Logout)
	// })
	r.Post("/login", loginHandler.Login)

	fileServer := http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))
	r.Handle("/images/*", fileServer)

	bannerHandler := handlers.NewBannerHandler(db.DB)
	r.Get("/banner", bannerHandler.GetBanners)
	serviceHandler := handlers.NewServiceHandler(db.DB)
	r.Get("/service", serviceHandler.GetServices)

	balanceHandler := handlers.NewBalanceHandler(db.DB)
	r.Get("/balance", balanceHandler.GetBalance)
	r.Post("/topup", balanceHandler.TopUp)

	transactionHandler := handlers.NewTransactionHandler(db.DB)
	r.Post("/transaction", transactionHandler.Transaction)
	r.Get("/transaction/history", transactionHandler.GetTransactions)
}
