package main

import (
	"flag"
	"fmt"
	"net/http"
)

// Reflects the parameters of the request back into the response
func ReflectionHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "%#v\n%+v", r, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {
	bind := flag.String("bind", "127.0.0.1:8080", "listen address to use")
	flag.Parse()

	http.HandleFunc("/", ReflectionHandler)
	http.ListenAndServe(*bind, nil)
}
