package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/code/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested")
	})
	fs := http.FileServer(http.Dir("template/"))
	http.Handle("/template/", http.StripPrefix("/template/", fs))
	http.ListenAndServe(":8080", nil)

}
