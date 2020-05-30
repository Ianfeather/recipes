package main

import (
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Ok there!"))
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
