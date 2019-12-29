package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServerMux struct {
}

func (p *CustomServerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random is: %f", rand.Float64())
}

func main() {
	// Any struct that has serveHTTP function can be a multiplexer
	// mux := &CustomServerMux{}
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
 		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
 		fmt.Fprintln(w, rand.Intn(100) )
	})

	http.ListenAndServe(":8000", newMux)
}