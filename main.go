package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if word1 := r.FormValue("word1"); word1 != "" {
		log.Printf("word1: %s", r.FormValue("word1"))
	}
	if word2 := r.FormValue("word2"); word2 != "" {
		log.Printf("word2: %s", r.FormValue("word2"))
	}
	if word3 := r.FormValue("word3"); word3 != "" {
		log.Printf("word3: %s", r.FormValue("word3"))
	}

	t.templ.Execute(w, r)
}

func main() {
	//handle functions
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/story", &templateHandler{filename: "story.html"})

	// Start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Ran into an error:", err)
	} else {
		log.Println("Serving on http://localhost:8080")
	}
}
