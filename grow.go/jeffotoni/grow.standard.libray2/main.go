package main

import (
	"encoding/json"
	"hash/maphash"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	// sy "github.com/jeffotoni/api1/syncmap"
	"github.com/patrickmn/go-cache"
)

var c = cache.New(25*time.Minute, 10*time.Minute)

// var m = sy.NewSyncMap()

type Growth struct {
	Id        string  `json:"id,omitempty"`
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
	err := json.NewDecoder(r.Body).Decode(&growth)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"msg":"error NewDecoder}`))
		return
	}

	// b, err := json.Marshal(growth)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{"msg":"error Marshal"}`))
	// 	return
	// }

	// b, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{"msg":"error body parse}`))
	// 	return
	// }

	// var growth sy.Growth
	// err = json.Unmarshal(b, &growth)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{"msg":"error unmarshal}`))
	// 	return
	// }

	key := strconv.Itoa(rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64()))).Int())
	growth.Id = key
	c.Set(key, growth, cache.NoExpiration)
	// m.Store(key, growth)

	w.Header().Set("Engine", "Go")
	w.Header().Set("key", key)
	w.Header().Set("Country", growth.Country)
	w.Header().Set("Indicator", growth.Indicator)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(""))
	return
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"msg":"success"}`))
	return
}
