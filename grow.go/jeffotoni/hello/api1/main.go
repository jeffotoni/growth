package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Growth struct {
	Country   string  `json:"country,omitempty"`
	Indicator string  `json:"indicator,omitempty"`
	Value     float32 `json:"value,omitempty"`
	Year      int     `json:"year,omitempty"`
}

func main() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/v1/growth", Add)
	log.Println("Run Server port:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var growth Growth
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"msg":"error body parse}`))
		return
	}

	err = json.Unmarshal(b, &growth)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"msg":"error unmarshal}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Engine", "Go")
	w.Header().Add("Country", growth.Country)
	w.Header().Add("Indicator", growth.Indicator)
	w.Write(b)
	return
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"msg":"success"}`))
	return
}
