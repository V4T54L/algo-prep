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

func (t *TemplRenderer) Render(w http.ResponseWriter, r *http.Request, partial, full templ.Component) error {
	var component templ.Component

	if val, ok := r.Header["Hx-Request"]; ok && val[0] == "true" {
		component = partial
	} else {
		component = full
	}
	templ.Handler(component).ServeHTTP(w, r)
	return nil
}
