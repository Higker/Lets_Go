package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	FORMAT = "2006 01/02 15:04:05 PM"
)

func runServer() {
	http.HandleFunc("/", welcome)
	log.Fatal(http.ListenAndServe("localhost:8090", nil))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Golang!\nServerTime: "+time.Now().Format(FORMAT)+"</h1>")
}
func main() {
	runServer()
}
