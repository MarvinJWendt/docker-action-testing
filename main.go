package main

import "net/http"

func main() {
	// Hello World HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Start HTTP Server
	http.ListenAndServe(":80", nil)
}
