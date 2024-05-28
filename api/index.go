package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tobyrushton/go-chi-vercel-example/handlers"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	r := chi.NewRouter()

	r.Get("/", handlers.NewHelloWorldHandler().ServeHTTP)

	r.ServeHTTP(w, req)
}
