// Package web web/handlers.go
package web

import (
	"html/template"
	"net/http"
)

type Handler struct {
	templates *template.Template
}

func NewHandler() (*Handler, error) {
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Handler{
		templates: tmpl,
	}, nil
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handleIndex)
	mux.HandleFunc("/patterns/new", h.handleNewPattern)
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := h.templates.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (h *Handler) handleNewPattern(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "pattern_form", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
