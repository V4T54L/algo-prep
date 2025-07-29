
package router

import (
	"app/internal/infrastructure/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(postHandler *handlers.PostHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", postHandler.List)
	r.Get("/posts/new", postHandler.NewForm)
	r.Post("/posts", postHandler.Create)
	return r
}
