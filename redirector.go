package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var redirUrl string
var port int
var host string
var statusCode int

func redirect(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := redirUrl
	code := statusCode
	var err error

	if r.Form.Has("u") {
		url = r.Form.Get("u")
	}
	if r.Form.Has("c") {
		code, err = strconv.Atoi(r.Form.Get("c"))
		if err != nil {
			fmt.Println("Supplied status code is not a valid int - ", r.Form.Get("c"), " - ", err)
			code = statusCode
		}
	}
	fmt.Println("Request from: ", r.RemoteAddr, "Status code:", code, " to: ", url)
	http.Redirect(w, r, url, code)
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
