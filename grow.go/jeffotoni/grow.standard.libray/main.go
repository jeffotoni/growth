package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// docker build -t jeffotoni/apigrow -f Dockerfile .
// docker run --rm -it -p 8080:8080 jeffotoni/apigrow
// curl localhost:8080/ping
var(
	mapGrow sync.Map
	mapGrowCount sync.Map
)
/* Example
[
   {
      "Country":"BRZ",
      "Indicator":"NGDP_R",
      "Value":183.26,
      "Year":2002
   },
   {
      "Country":"AFG",
      "Indicator":"NGDP_R",
      "Value":198.736,
      "Year":2003
   }
]
*/
type dataGrowth struct {
	Country   string  `json:"Country"`
	Indicator string  `json:"Indicator"`
	Value     float64 `json:"Value"`
	Year      int     `json:"Year"`
}

func init(){
	mapGrowCount.Store("count", 0)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong 1"))
		})
	mux.HandleFunc("/api/v1/growth",Route)
	mux.HandleFunc("/api/v1/growth/post/status",GetStatus)
	mux.HandleFunc("/api/v1/growth/size",GetSize)
	mux.HandleFunc("/",Route)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Print("Run Server:8080")
	log.Fatal(server.ListenAndServe())
}

func Route(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case http.MethodGet:
		Get(w, r)
	case http.MethodDelete:
		Delete(w, r)
	case http.MethodPut:
		Put(w, r)
	case http.MethodPost:
		Post(w, r)
	default:
		http.NotFound(w,r)
	}
}

func Put(w http.ResponseWriter, r *http.Request){
	var err error
	var code int = 400
	elem := strings.Split(r.URL.Path, "/")
	if len(elem) != 7 {
		log.Println("len:", len(elem), " path:", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"error in path"}`))
		return
	}

	type putGrow struct{
		Value float64 `json:"value"`
	}

	var putg putGrow
	err = json.NewDecoder(r.Body).Decode(&putg)
	if err!=nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"error in decode json value has to be float"}`))
		return
	}
	defer r.Body.Close()

	country := strings.ToUpper(elem[4])
	Indicator := strings.ToUpper(elem[5])
	year := elem[6]
	key := country + Indicator + year
	_, ok := mapGrow.Load(key)
	if ok {
		mapGrow.Store(key, putg.Value)
		code = http.StatusOK
	} else{
		mapGrow.LoadOrStore(key, putg.Value)
		countInt, _ := mapGrowCount.Load("count")
		count := countInt.(int)
		count = count + 1
		mapGrowCount.Store("count", count)
		log.Println("inserted new record in memory:", count)
		code = http.StatusCreated
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func Delete(w http.ResponseWriter, r *http.Request){
	var code int = 400
	elem := strings.Split(r.URL.Path, "/")
	if len(elem) != 7 {
		log.Println("len:", len(elem), " path:", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"error in path"}`))
		return
	}

	country := strings.ToUpper(elem[4])
	Indicator := strings.ToUpper(elem[5])
	year := elem[6]
	key := country + Indicator + year
	_, ok := mapGrow.Load(key)
	if ok {
		mapGrow.Delete(key)
		code = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
}

func Get(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var err error
	var code int = 400
	elem := strings.Split(r.URL.Path, "/")
	if len(elem) != 7 {
		log.Println("len:", len(elem), " path:", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"msg":"error in path"}`))
		return
	}

	country := strings.ToUpper(elem[4])
	Indicator := strings.ToUpper(elem[5])
	year := elem[6]
	key := country + Indicator + year
	val, ok := mapGrow.Load(key)
	if ok {
		var grow dataGrowth
		grow.Value = val.(float64)
		grow.Country = country
		grow.Indicator = Indicator
		grow.Year, _ = strconv.Atoi(year)
		b, err = json.Marshal(&grow)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"msg":"error marshal:` + err.Error() + `"}`))
			return
		}
		code = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

func GetSize(w http.ResponseWriter, r *http.Request) {
	var sizeInt int =0
	var sizeStr string
	size, ok := mapGrowCount.Load("count")
	if ok {
		sizeInt = size.(int)
	}
	sizeStr = strconv.Itoa(sizeInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"size":`+sizeStr+`}`))
}

func GetStatus(w http.ResponseWriter, r *http.Request){
	key, ok := mapGrow.Load("BRZNGDP_R2002")
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"msg":"not finished"}`))
		return
	}

	var count_str string
	count, ok := mapGrowCount.Load("count")
	if ok {
		count_str = strconv.Itoa(count.(int))
	}
	result := fmt.Sprintf("%.2f",key.(float64))
	log.Println("value valid:",result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"msg":"complete","test value"":` +result+ `, "count":`+ count_str +`}`))
}

func Post(w http.ResponseWriter, r *http.Request){
	var grow []dataGrowth
	err := json.NewDecoder(r.Body).Decode(&grow)
	if err!=nil{
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"msg":"error in your json"}`))
		return
	}
	defer r.Body.Close()

	go func(grow []dataGrowth){
		var cnew int = 0
		for _, v := range grow {
			year := strconv.Itoa(v.Year)
			key := strings.ToUpper(v.Country) + strings.ToUpper(v.Indicator) + year
			_, ok := mapGrow.LoadOrStore(key, v.Value)
			if !ok {
				cnew++
			}
		}
		countInt, _ := mapGrowCount.Load("count")
		count := countInt.(int)
		count = count + cnew
		mapGrowCount.Store("count", count)
		log.Println("successfully loaded data into memory:", count)
	}(grow)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"msg":"In progress"}`))
}