package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type movie struct {
	name   string
	actors []string
}

func main() {
	fmt.Println("hei")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	http.ListenAndServe(":4000", r)
}

func greeter() {
	fmt.Println("greete")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang on Yt</h1>"))
}
