package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")
	APP := os.Getenv("APP")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s, %q", APP, html.EscapeString(r.URL.Path))
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "APP2, %q", html.EscapeString(r.URL.Path))
	// })

	log.Fatal(http.ListenAndServe(":"+PORT, nil))

}
