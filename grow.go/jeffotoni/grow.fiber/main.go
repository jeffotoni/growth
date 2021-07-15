package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
}

func main() {
	app := fiber.New(fiber.Config{BodyLimit: 10 * 1024 * 1024})
	app.Use(cors.New())
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

	go func(grow []dataGrowth) {
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
	}(grow)

	return c.Status(202).SendString(`{"msg":"In progress"}`)
}

func GetStatus(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	key, ok := mapGrow.Load("BRZNGDP_R2002")
	if !ok {
		return c.Status(400).SendString(`{"msg":"not finished"}`)
	}
	var count_str string
	count, ok := mapGrowCount.Load("count")
	if ok {
		count_str = strconv.Itoa(count.(int))
	}
	result := fmt.Sprintf("%.2f", key.(float64))
	return c.Status(200).SendString(`{"msg":"complete","test value"":` + result + `, "count":` + count_str + `}`)
}

func GetSize(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var sizeInt int = 0
	var sizeStr string
	size, ok := mapGrowCount.Load("count")
	if ok {
		sizeInt = size.(int)
	}
	sizeStr = strconv.Itoa(sizeInt)
	return c.Status(200).SendString(`{"size":` + sizeStr + `}`)
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
		//log.Println("inserted new record in memory:", count)
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
