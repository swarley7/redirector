package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var redirUrl string
var port int
var host string

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, redirUrl, 301)
}

func main() {
	flag.StringVar(&redirUrl, "redir", "https://example.com/aaa", "Redirect URL")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&host, "host", "localhost", "Host to listen on")
	http.HandleFunc("/", redirect)
	log.Printf("Listening on %s:%d - Redirecting requests to %s", host, port, redirUrl)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
