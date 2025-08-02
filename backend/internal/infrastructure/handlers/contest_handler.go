package handlers

import (
	"app/internal/application/usecase"
	"app/internal/infrastructure/template"
	"app/web/templates"
	"net/http"

	"github.com/a-h/templ"
)

type ContestHandler struct {
	usecase  usecase.ContestUseCase
	renderer template.TemplRenderer
}

func NewContestHandler(u usecase.ContestUseCase, r template.TemplRenderer) *ContestHandler {
	return &ContestHandler{usecase: u, renderer: r}
}

func (h *ContestHandler) GetContestsPage(w http.ResponseWriter, r *http.Request) {
	contests, _ := h.usecase.ListContests(r.Context())
	h.renderer.Render(w, r, templates.ContestListPage(contests, "Upcoming"), templates.Base([]templ.Component{templates.ContestListPage(contests, "Upcoming")}))
}
