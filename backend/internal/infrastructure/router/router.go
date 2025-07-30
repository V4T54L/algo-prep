package router

import (
	"app/internal/infrastructure/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(postHandler *handlers.PostHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", postHandler.List)
	r.Get("/posts/new", postHandler.NewForm)
	r.Post("/posts", postHandler.Create)
	return r
}
