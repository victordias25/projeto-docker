package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8081", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá"))
}
