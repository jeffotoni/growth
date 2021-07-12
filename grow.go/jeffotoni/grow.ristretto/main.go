package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
	"github.com/jeffotoni/grow.go/jeffotoni/grow.ristretto/pkg/ristretto"
)

var (
	mapGrow      sync.Map
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

func init() {
	ristretto.Set("count", "0")
	//mapGrowCount.Store("count", 0)
}

// Middleware Logger
func Logger(name string) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			h.ServeHTTP(w, r)
			log.Printf(
				"%s %s \u001B[1;36m%s\u001B[0m \u001B[1;33m%s\u001B[0m \u001B[0;34m%s\u001B[0m \033[1;32m%s\033[0m",
				r.Header.Get("User-Agent"),
				r.RequestURI,
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		})
	}
}

type Adapter func(http.Handler) http.Handler

// Middleware
func Middleware(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ping",
		Middleware(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("pong😍"))
			}),
			Logger(""),
		))
	mux.Handle("/api/v1/growth",
		Middleware(http.HandlerFunc(Route),
			Logger(""),
		))
	mux.Handle("/api/v1/growth/post/status",
		Middleware(http.HandlerFunc(GetStatus),
			Logger(""),
		))
	mux.Handle("/api/v1/growth/size",
		Middleware(http.HandlerFunc(GetSize),
			Logger(""),
		))
	mux.Handle("/",
		Middleware(http.HandlerFunc(Route),
			Logger(""),
		))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	log.Println("\033[1;44mRunning on http://0.0.0.0:8080 (Press CTRL+C to quit)\033[0m")
	log.Fatal(server.ListenAndServe())
}

func Route(w http.ResponseWriter, r *http.Request) {
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
		http.NotFound(w, r)
	}
}

func Put(w http.ResponseWriter, r *http.Request) {
	var err error
	var code int = 400
	elem := strings.Split(r.URL.Path, "/")
	if len(elem) != 7 {
		WriteService(w, r, code,`{"msg":"error in path"}`)
		return
	}

	type putGrow struct {
		Value float64 `json:"value"`
	}

	var putg putGrow
	err = json.NewDecoder(r.Body).Decode(&putg)
	if err != nil {
		WriteService(w, r, code,`{"msg":"error in decode json value has to be float"}`)
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
	} else {
		mapGrow.LoadOrStore(key, putg.Value)
		countStr := ristretto.Get("count")
		count, _ := strconv.Atoi(countStr)
		count = count + 1
		countStr = strconv.Itoa(count)
		ristretto.Set("count", countStr)
		code = http.StatusCreated
	}
	WriteService(w, r, code,"")
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var code int = 400
	elem := strings.Split(r.URL.Path, "/")
	if len(elem) != 7 {
		WriteService(w, r, code,`{"msg":"error in path"}`)
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
	WriteService(w, r, code,"")
}

func Get(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var code int = 400
	elem := strings.Split(r.URL.Path, "/")
	if len(elem) != 7 {
		WriteService(w, r, code,`{"msg":"error in path"}`)
		return
	}

	country := strings.ToUpper(elem[4])
	Indicator := strings.ToUpper(elem[5])
	year := elem[6]
	key := country + Indicator + year
	val := ristretto.Get(key)
	if len(val)>0 {
		var grow dataGrowth
		fval, err := strconv.ParseFloat(val,64)
		if err != nil {
			log.Println("error parse:" ,err.Error())
			WriteService(w, r, code,`{"msg":"error parse float"}`)
			return
		}
		grow.Value = fval
		grow.Country = country
		grow.Indicator = Indicator
		grow.Year, _ = strconv.Atoi(year)
		b, err = json.Marshal(&grow)
		if err != nil {
			WriteService(w, r, code,`{"msg":"error marshal:` + err.Error() + `"}`)
			return
		}
		code = http.StatusOK
	}
	WriteService(w, r, code,string(b))
}

func GetSize(w http.ResponseWriter, r *http.Request) {
	var sizeInt int = 0
	var sizeStr string
	size, ok := mapGrowCount.Load("count")
	if ok {
		sizeInt = size.(int)
	}
	sizeStr = strconv.Itoa(sizeInt)
	WriteService(w, r, 200,`{"size":` + sizeStr + `}`)
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	key, ok := mapGrow.Load("BRZNGDP_R2002")
	if !ok {
		WriteService(w, r, 400, `{"msg":"not finished"}`)
		return
	}

	var count_str string
	count, ok := mapGrowCount.Load("count")
	if ok {
		count_str = strconv.Itoa(count.(int))
	}
	result := fmt.Sprintf("%.2f", key.(float64))
	WriteService(w, r, 200, `{"msg":"complete","test value"":` + result + `, "count":` + count_str + `}`)
}

func Post(w http.ResponseWriter, r *http.Request) {
	var grow []dataGrowth
	err := json.NewDecoder(r.Body).Decode(&grow)
	if err != nil {
		WriteService(w, r, 400, `{"msg":"error in your json"}`)
		return
	}
	defer r.Body.Close()
	go func(grow []dataGrowth) {
		var cnew int = 0
		for _, v := range grow {
			year := strconv.Itoa(v.Year)
			key := strings.ToUpper(v.Country) + strings.ToUpper(v.Indicator) + year
			sold := ristretto.Get(key)
			if len(sold) == 0 {
				cnew++
			}
			ristretto.Set(key, sold)
		}
		countStr := ristretto.Get("count")
		count, _ := strconv.Atoi(countStr)
		count = count + cnew
		countStr = strconv.Itoa(count)
		ristretto.Set("count", countStr)
	}(grow)
		WriteService(w, r, 202, `{"msg":"In progress"}`)
}

func WriteService(w http.ResponseWriter, r *http.Request, code int, msg string){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Web.Server", "net/http")
	w.WriteHeader(code)
	if len(msg) == 0 {
		return
	}
	w.Write([]byte(msg))
}