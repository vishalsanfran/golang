package main

import (
	"log"
	"net/http"
)

func main() {
	g := newGroup()
	http.Handle("/group", g)
	go g.run()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
