package main

import (
	"log"
	"net/http"
	"webService/internal"
)

func main() {
	hub := internal.NewHub()
	go hub.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		internal.ServeWs(hub, w, r)
	})
	err := http.ListenAndServe(":8087", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
