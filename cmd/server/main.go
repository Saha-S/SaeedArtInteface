package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"art-interface/internal/art"
	httpHandler "art-interface/internal/http"
)

func main() {
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}

	tmpl := template.Must(template.ParseFiles("web/templates/index.html"))
	svc := art.NewService()
	h := httpHandler.NewHandler(tmpl, svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/decoder", h.Decode)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Printf("art-interface server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
