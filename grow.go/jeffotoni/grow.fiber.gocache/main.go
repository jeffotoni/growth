package main

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
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
	mapGrowCount.Store("count", 0)
	mapGrow.Store("BRZNGDPX_R2002", "183.26")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app := fiber.New(fiber.Config{
		BodyLimit:    10 * 1024 * 1024,
		Prefork:      true,
		ServerHeader: "Fiber",
	})
	//app.Use(cors.New())
	// app.Use(logger.New(logger.Config{
	// 	Format:     "${pid} ${time} ${method} ${path} - ${ip} - \u001B[0;34m${status}\u001B[0m - \033[1;32m${latency}\033[0m\n",
	// 	TimeFormat: "02-Jan-2006 15:04:05",
	// 	Output:     os.Stdout,
	// }))

	app.Get("/ping", Ping)
	app.Post("/api/v1/growth", Post)
	app.Get("/api/v1/growth/post/status", GetStatus)
	app.Get("/api/v1/growth/size", GetSize)
	app.Get("/api/v1/growth/:country/:indicator/:year", Get)
	app.Put("/api/v1/growth/:country/:indicator/:year", Put)
	app.Delete("/api/v1/growth/:country/:indicator/:year", Delete)
	app.Listen("0.0.0.0:8080")
}

func Ping(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	return c.Status(200).SendString(`{"msg":"pong❤"}`)
}

func Post(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var grow []dataGrowth
	err := c.BodyParser(&grow)
	if err != nil {
		return c.Status(400).SendString(`{"msg":"error in your json"}`)
	}

	// for _, v := range grow {
	// 	year := strconv.Itoa(v.Year)
	// 	bs := make([]byte, 100)
	// 	bl := 0
	// 	bl += copy(bs[bl:], strings.ToUpper(v.Country))
	// 	bl += copy(bs[bl:], strings.ToUpper(v.Indicator))
	// 	bl += copy(bs[bl:], year)
	// 	mapGrow.Store(string(bs), v.Value)
	// }

	var numJobs = len(grow)
	var jobs = make(chan dataGrowth, numJobs)
	for w := 0; w < 11; w++ {
		go worker(w, jobs)
	}
	for _, tgrow := range grow {
		jobs <- tgrow
	}
	close(jobs)
	return c.Status(202).SendString(`{"msg":"In progress"}`)
}

func worker(id int, grow <-chan dataGrowth) {
	for v := range grow {
		var strBuilder strings.Builder
		year := strconv.Itoa(v.Year)
		strBuilder.WriteString(strings.ToUpper(v.Country))
		strBuilder.WriteString(strings.ToUpper(v.Indicator))
		strBuilder.WriteString(strings.ToUpper(year))

		// bs := make([]byte, 100)
		// bl := 0
		// bl += copy(bs[bl:], strings.ToUpper(v.Country))
		// bl += copy(bs[bl:], strings.ToUpper(v.Indicator))
		// bl += copy(bs[bl:], year)
		mapGrow.Store(strBuilder.String(), v.Value)
	}

}

func GetStatus(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	length := 0
	mapGrow.Range(func(_, _ interface{}) bool {
		length++
		return true
	})
	count_str := strconv.Itoa(length)
	return c.Status(200).SendString(`{"msg":"complete","count":` + count_str + `}`)
}

func GetSize(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	length := 0
	mapGrow.Range(func(key interface{}, value interface{}) bool {
		length++
		println("here")
		return true
	})
	count_str := strconv.Itoa(length)
	return c.Status(200).SendString(`{"msg":"complete","count":` + count_str + `}`)
}

func Put(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	var code int = 400
	country := strings.ToUpper(c.Params("country"))
	Indicator := strings.ToUpper(c.Params("indicator"))
	year := c.Params("year")
	if len(country) == 0 || len(Indicator) == 0 || len(year) != 4 {
		//log.Println("len:", len(elem), " path:", r.URL.Path)
		c.Status(400).SendString(`{"msg":"error in path url"}`)
		return
	}

	type putGrow struct {
		Value float64 `json:"value"`
	}

	var putg putGrow
	err = c.BodyParser(&putg)
	if err != nil {
		return c.Status(400).SendString(`{"msg":"error in decode json value has to be float"}`)
	}

	key := country + Indicator + year
	_, ok := mapGrow.Load(key)
	if ok {
		mapGrow.Store(key, putg.Value)
		code = 200
	} else {
		mapGrow.LoadOrStore(key, putg.Value)
		countInt, _ := mapGrowCount.Load("count")
		count := countInt.(int)
		count = count + 1
		mapGrowCount.Store("count", count)
		code = 201
	}
	return c.Status(code).SendString("")
}

func Delete(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var code int = 400
	country := strings.ToUpper(c.Params("country"))
	Indicator := strings.ToUpper(c.Params("indicator"))
	year := c.Params("year")
	if len(country) == 0 || len(Indicator) == 0 || len(year) != 4 {
		//log.Println("len:", len(elem), " path:", r.URL.Path)
		return c.Status(400).SendString(`{"msg":"error in path url"}`)
	}
	key := country + Indicator + year
	_, ok := mapGrow.Load(key)
	if ok {
		mapGrow.Delete(key)
		countInt, _ := mapGrowCount.Load("count")
		count := countInt.(int)
		count = count - 1
		mapGrowCount.Store("count", count)
		code = 202
	}
	return c.Status(code).SendString("")
}

func Get(c *fiber.Ctx) (err error) {
	c.Set("Content-Type", "application/json")
	var b []byte
	var code int = 400
	country := strings.ToUpper(c.Params("country"))
	Indicator := strings.ToUpper(c.Params("indicator"))
	year := c.Params("year")
	if len(country) == 0 || len(Indicator) == 0 || len(year) != 4 {
		return c.Status(400).SendString(`{"msg":"error in path url"}`)
	}
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
			return c.Status(400).SendString(`{"msg":"error marshal:` + err.Error() + `"}`)
		}
		code = 200
	}
	return c.Status(code).SendString(string(b))
}
