// Package web web/handlers.go
package web

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/tech-sam/GoPatternLab/pkg/db"
	"github.com/tech-sam/GoPatternLab/pkg/db/models"
)

type Handler struct {
	templates *template.Template
	patterns  *models.PatternModel
}

func NewHandler(db *db.DB) (*Handler, error) {
	tmpl, err := template.ParseGlob("web/templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Handler{
		templates: tmpl,
		patterns:  models.NewPatternModel(db),
	}, nil
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.handleIndex)
	mux.HandleFunc("/patterns/new", h.handleNewPattern)
	mux.HandleFunc("/patterns/create", h.handleCreatePattern)
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Get patterns from database
	patterns, err := h.patterns.GetPatterns()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute base.html which will automatically include index.html's content
	// because index.html defines the "content" template that base.html references
	err = h.templates.ExecuteTemplate(w, "base.html", map[string]interface{}{
		"Patterns": patterns,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handleNewPattern(w http.ResponseWriter, r *http.Request) {
	err := h.templates.ExecuteTemplate(w, "pattern_form", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) handleCreatePattern(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := strings.TrimSpace(r.FormValue("name"))
	description := strings.TrimSpace(r.FormValue("description"))

	if name == "" {
		http.Error(w, "Pattern name is required", http.StatusBadRequest)
		return
	}

	// Create the pattern and get back the created pattern
	pattern, err := h.patterns.Create(name, description)
	if err != nil {
		http.Error(w, "Failed to create pattern", http.StatusInternalServerError)
		return
	}

	err = h.templates.ExecuteTemplate(w, "success_message", pattern)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
