package server

import (
	"fmt"
	"net/http"
)

func Init() {
	http.HandleFunc("/", Handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome")
}

func StartServer() {
	http.ListenAndServe(":1357", nil)
}