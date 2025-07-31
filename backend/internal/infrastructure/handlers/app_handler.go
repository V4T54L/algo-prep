package handlers

import (
	"app/internal/application/usecase"
	"app/internal/infrastructure/template"
	"app/web/templates"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

type AppHandler struct {
	usecase  usecase.PostUseCase
	renderer template.TemplRenderer
}

func NewAppHandler(u usecase.PostUseCase, r template.TemplRenderer) *AppHandler {
	return &AppHandler{usecase: u, renderer: r}
}

func (h *AppHandler) GetHomePage(w http.ResponseWriter, r *http.Request) {
	posts, _ := h.usecase.ListPosts(r.Context())
	log.Println("Posts count: ", len(posts))
	h.renderer.Render(w, r, templates.Homepage(posts), templates.Base(([]templ.Component{templates.Homepage(posts)})))
}
