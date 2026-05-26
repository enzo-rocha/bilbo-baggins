package main

import (
	"log"
	"net/http"

	httpdelivery "bilbo-baggins/internal/delivery/http"
	"bilbo-baggins/internal/infra/web"
	"bilbo-baggins/internal/usecase"
)

func main() {
	homeRepository := web.NewHomePageRepository()
	showHomePage := usecase.NewShowHomePage(homeRepository)
	homeHandler := httpdelivery.NewHomeHandler(showHomePage)

	router := httpdelivery.NewRouter(homeHandler, web.StaticFileSystem())

	log.Println("server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
