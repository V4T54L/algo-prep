package handlers

import (
	"app/internal/application/usecase"
	"app/internal/infrastructure/template"
	"app/web/templates"
	"net/http"

	"github.com/a-h/templ"
)

type ProblemHandler struct {
	usecase  usecase.ProblemUseCase
	renderer template.TemplRenderer
}

func NewProblemHandler(u usecase.ProblemUseCase, r template.TemplRenderer) *ProblemHandler {
	return &ProblemHandler{usecase: u, renderer: r}
}

func (h *ProblemHandler) GetProblemsPage(w http.ResponseWriter, r *http.Request) {
	problems, _ := h.usecase.ListProblems(r.Context())
	h.renderer.Render(w, r, templates.ProblemListPage(problems), templates.Base([]templ.Component{templates.ProblemListPage(problems)}))
}
