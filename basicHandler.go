package main

import (
	"io"
	"net/http"
	"log"
)

func MyServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}