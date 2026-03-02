package http

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"art-interface/internal/art"
)

const testTemplate = `{{define "index.html"}}HTTP {{.StatusCode}} {{.Result}} {{.ErrorText}}{{end}}`

func newTestHandler(t *testing.T) *Handler {
	t.Helper()
	tmpl := template.Must(template.New("index.html").Parse(testTemplate))
	return NewHandler(tmpl, art.NewService())
}

func TestHomeOK(t *testing.T) {
	h := newTestHandler(t)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	h.Home(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("got %d want %d", w.Code, http.StatusOK)
	}
}

func TestDecodeAccepted(t *testing.T) {
	h := newTestHandler(t)
	form := url.Values{}
	form.Set("input", "[3 A][2 B]")
	form.Set("mode", "decode")

	req := httptest.NewRequest(http.MethodPost, "/decoder", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	h.Decode(w, req)

	if w.Code != http.StatusAccepted {
		t.Fatalf("got %d want %d", w.Code, http.StatusAccepted)
	}
}

func TestDecodeBadRequest(t *testing.T) {
	h := newTestHandler(t)
	form := url.Values{}
	form.Set("input", "[x #]")
	form.Set("mode", "decode")

	req := httptest.NewRequest(http.MethodPost, "/decoder", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	h.Decode(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("got %d want %d", w.Code, http.StatusBadRequest)
	}
}
