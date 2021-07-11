package main

import (
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${time} ${method} ${path} - ${ip} - \u001B[0;34m${status}\u001B[0m - \033[1;32m${latency}\033[0m\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     os.Stdout,
	}))

	app.Get("/ping", Ping)
	app.Post("/api/v1/growth", Post)

	app.Listen("0.0.0.0:8080")
}

func Ping(c *fiber.Ctx) error {
	return c.Status(200).SendString(`{"msg":"pong❤️"}`)
}

func Post(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")
	var grow []dataGrowth
	err := c.BodyParser(&grow)
	if err != nil {
		return c.Status(200).SendString(`{"msg":"error in your json"}`)
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

	return c.Status(200).SendString(`{"msg":"In progress"}`)
}
