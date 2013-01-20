package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
	// "model"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hi there, I love %s!", r.Method)
		fmt.Fprintf(w, r.URL.Path)
	})
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hi there, I love %s!", r.Method)
		fmt.Fprintf(w, r.URL.Path)
	})
	// http.ListenAndServe(":8080", nil)
	s := &http.Server{
		Addr:           ":8080",
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}