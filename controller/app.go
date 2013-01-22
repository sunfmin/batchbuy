package main

import (
	"net/http"
	"time"
	"log"
	"strings"
	// "model"
)

func serverFile(w http.ResponseWriter, r *http.Request) {
	if strings.Contains(r.URL.Path, "/assets") {
		http.ServeFile(w, r, ".." + r.URL.Path)
	} else {
		http.ServeFile(w, r, "../view" + r.URL.Path + ".html")
	}
}

func main() {
	http.HandleFunc("/assets/", serverFile)
	http.HandleFunc("/profile", serverFile)
	http.HandleFunc("/product", serverFile)
	http.HandleFunc("/order", serverFile)
	http.HandleFunc("/order_list", serverFile)
	// http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
	// })
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