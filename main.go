package main

import "net/http"

func main() {

	r := newManager()
	http.Handle("/", http.FileServer(http.Dir("webpage")))
	http.HandleFunc("/ws", r.serveWS)

	http.ListenAndServe(":3000", nil)
}
