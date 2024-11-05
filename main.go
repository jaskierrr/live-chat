package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

type templHandler struct {
	once  sync.Once
	templ template.Template
}

func (t *templHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = *template.Must(template.ParseFiles("templates/index.html"))
	})
	t.templ.Execute(w, r)
}

func main() {
	r := newRoom()

	http.Handle("/", &templHandler{})
	http.Handle("/room", r)

	go r.run()

	log.Println("Starting web server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
