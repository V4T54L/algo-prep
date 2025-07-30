package handlers

import (
	"app/internal/application/usecase"
	"app/internal/infrastructure/template"
	"app/web/templates"
	"net/http"
)

type PostHandler struct {
	usecase  usecase.PostUseCase
	renderer template.TemplRenderer
}

func NewPostHandler(u usecase.PostUseCase, r template.TemplRenderer) *PostHandler {
	return &PostHandler{usecase: u, renderer: r}
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	posts, _ := h.usecase.ListPosts(r.Context())
	h.renderer.Render(w, r, templates.PostList(posts), templates.Base(templates.PostList(posts)))
}

func (h *PostHandler) NewForm(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, r, templates.PostForm(), templates.Base(templates.PostForm()))
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	authorID := "1" // Hardcoded for demo
	h.usecase.CreatePost(r.Context(), title, content, authorID)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
