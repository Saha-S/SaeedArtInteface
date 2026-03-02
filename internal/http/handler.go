package http

import (
	"html/template"
	"net/http"
	"strings"

	"art-interface/internal/art"
)

type PageData struct {
	Input      string
	Result     string
	Mode       string
	Multiline  bool
	StatusCode int
	ErrorText  string
}

type Handler struct {
	tmpl *template.Template
	svc  *art.Service
}

func NewHandler(t *template.Template, svc *art.Service) *Handler {
	return &Handler{tmpl: t, svc: svc}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = h.tmpl.ExecuteTemplate(w, "index.html", PageData{Mode: string(art.ModeDecode), StatusCode: http.StatusOK})
}

func (h *Handler) Decode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		h.renderError(w, http.StatusBadRequest, "Malformed form data")
		return
	}

	input := r.FormValue("input")
	mode := art.Mode(r.FormValue("mode"))
	multiline := r.FormValue("multiline") == "on"

	if strings.TrimSpace(input) == "" {
		h.renderError(w, http.StatusBadRequest, "Input is required")
		return
	}
	if mode == "" {
		mode = art.ModeDecode
	}
	if mode != art.ModeDecode && mode != art.ModeEncode {
		h.renderError(w, http.StatusBadRequest, "Unsupported mode")
		return
	}

	out, err := h.svc.Execute(mode, input, multiline)
	if err != nil {
		h.renderError(w, http.StatusBadRequest, "Malformed encoded string")
		return
	}

	w.WriteHeader(http.StatusAccepted)
	_ = h.tmpl.ExecuteTemplate(w, "index.html", PageData{
		Input:      input,
		Result:     out,
		Mode:       string(mode),
		Multiline:  multiline,
		StatusCode: http.StatusAccepted,
	})
}

func (h *Handler) renderError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	_ = h.tmpl.ExecuteTemplate(w, "index.html", PageData{
		Mode:       string(art.ModeDecode),
		StatusCode: code,
		ErrorText:  msg,
	})
}
