package http

import (
	"log"
	"net/http"

	"bilbo-baggins/internal/usecase"
)

type HomeHandler struct {
	showHomePage usecase.ShowHomePage
}

func NewHomeHandler(showHomePage usecase.ShowHomePage) HomeHandler {
	return HomeHandler{
		showHomePage: showHomePage,
	}
}

func (handler HomeHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.Header().Set("Allow", http.MethodGet)
		http.Error(response, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	page, err := handler.showHomePage.Execute(request.Context())
	if err != nil {
		log.Printf("show home page: %v", err)
		http.Error(response, "internal server error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", page.ContentType)
	response.WriteHeader(http.StatusOK)
	_, _ = response.Write(page.Content)
}
