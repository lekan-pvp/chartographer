package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"internshipApplicationTemplate/internal/config"
	"internshipApplicationTemplate/internal/handlers"
	"internshipApplicationTemplate/internal/imageservice"
	"log"
	"net/http"
)

func main() {
	config.New()

	err := imageservice.New(config.Cfg)
	if err != nil {
		log.Fatal(err)
	}

	serverAddress := config.Cfg.ServerAddress
	log.Println("the server address is", serverAddress)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/chartas", func(r chi.Router) {
		r.Post("/", handlers.CreateImage)
		r.Post("/{id}/", handlers.SaveFragment)
	})

	log.Fatal(http.ListenAndServe(serverAddress, router))
}
