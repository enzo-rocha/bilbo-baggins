package http

import "net/http"

func NewRouter(homeHandler HomeHandler, staticFileSystem http.FileSystem) http.Handler {
	router := http.NewServeMux()
	router.Handle("/", homeHandler)
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(staticFileSystem)))

	return router
}
