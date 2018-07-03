package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hey, you've requested: %s</h1>\n", r.URL.Path)
	})

	defaultPort := "8080"
	log.Printf("Listening on default port: %s", defaultPort)
	log.Fatal(http.ListenAndServe(":" + defaultPort, nil))
}
