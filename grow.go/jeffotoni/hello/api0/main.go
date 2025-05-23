package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("0.0.0.0:8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			w.Header().Add("Engine", "Go")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"msg":"success"}`))
		})))
}
