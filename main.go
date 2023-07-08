package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		form := `<html>
					<head>
					<title>Chat</title>
					</head>
					<body>
						Real time chat
					</body>
				</html>`
		fmt.Fprintf(w, form)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
