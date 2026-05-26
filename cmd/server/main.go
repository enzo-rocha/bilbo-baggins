package main

import (
	"log"
	"net/http"
	"os"

	httpdelivery "bilbo-baggins/internal/delivery/http"
	"bilbo-baggins/internal/infra/web"
	"bilbo-baggins/internal/usecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	homeRepository := web.NewHomePageRepository()
	showHomePage := usecase.NewShowHomePage(homeRepository)
	homeHandler := httpdelivery.NewHomeHandler(showHomePage)

	router := httpdelivery.NewRouter(homeHandler, web.StaticFileSystem())

	log.Printf("server running at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
