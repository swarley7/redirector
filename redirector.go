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
var statusCode int

func redirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request from: ", r.RemoteAddr, "Status code:", statusCode, " to: ", redirUrl)
	http.Redirect(w, r, redirUrl, statusCode)
}

func main() {
	flag.StringVar(&redirUrl, "redir", "https://example.com/aaa", "Redirect URL")
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.StringVar(&host, "host", "localhost", "Host to listen on")
	flag.IntVar(&statusCode, "code", 302, "Status code to respond with")

	flag.Parse()

	http.HandleFunc("/", redirect)
	log.Printf("Listening on %s:%d - Redirecting (Status Code: %d) requests to %s", host, port, statusCode, redirUrl)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
