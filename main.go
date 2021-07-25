package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World")
}

func MyRouter() http.Handler {
	router := chi.NewRouter()
	router.Get("/", HelloWorld)
	return router
}

func main() {
	handler := MyRouter()
	server := &http.Server{
		Addr: ":8000",
		Handler: handler,
	}
	_ = server.ListenAndServe()
}
