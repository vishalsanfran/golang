package main

import (
	"log"
	"net/http"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(t.filename))
	})
	t.templ.Execute(w, nil)
}

func main() {
	g := newGroup()
	http.Handle("/", &templateHandler{filename: "chat_app.html"})
	http.Handle("/group", g)
	go g.run()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
