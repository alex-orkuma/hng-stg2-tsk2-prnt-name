package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello my name is SHIAONDO ORKUMA ALEX"))
	fmt.Println("Hello my name is SHIAOND ORKUMA ALEX")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Starting a server on port :9000")
	http.ListenAndServe(":9000", mux)
}
