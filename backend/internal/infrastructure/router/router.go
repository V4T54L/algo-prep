package router

import (
	"app/internal/infrastructure/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(postHandler *handlers.PostHandler, appHandler *handlers.AppHandler, problemHandler *handlers.ProblemHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", appHandler.GetHomePage)
	r.Get("/posts", postHandler.List)
	r.Get("/posts/new", postHandler.NewForm)
	r.Post("/posts", postHandler.Create)
	r.Get("/problems", problemHandler.GetProblemsPage)

	staticDir := http.Dir("static")
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(staticDir)))
	return r
}
