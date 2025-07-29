// -- internal/infrastructure/template/templ_renderer.go

package template

import (
	"net/http"

	"github.com/a-h/templ"
)

type TemplRenderer struct{}

func NewTemplRenderer() *TemplRenderer {
	return &TemplRenderer{}
}

func (t *TemplRenderer) Render(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	// if c, ok := component.(interface {
	// 	Render(context.Context, http.ResponseWriter) error
	// }); ok {
	// 	return c.Render(r.Context(), w)
	// }
	// http.Error(w, "invalid template component", http.StatusInternalServerError)

	templ.Handler(component).ServeHTTP(w, r)
	return nil
}
