package api

import (
	"deryksr/model"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
)

var srv *server
var process sync.Map

func initRouter(sv *server) {
	process.Store("busy", false)

	srv = sv
	group := srv.echo.Group("/api/v1/growth")

	group.POST("", createMemoryDb)
	group.GET("/size", dbSize)
	group.GET("/post/status", loadStatus)
	group.PUT("/:country/:indicator/:year", upsertRecord)
	group.GET("/:country/:indicator/:year", getRecord)
	group.DELETE("/:country/:indicator/:year", deleteRecord)
}

func createMemoryDb(ctx echo.Context) error {
	payload := make([]model.Growth, 0)

	if err := ctx.Bind(&payload); err != nil {
		return err
	}

	go func(payload []model.Growth) {
		process.Store("busy", true)
		for _, element := range payload {
			srv.database.Save(element)
		}
		process.Store("busy", false)
	}(payload)

	return ctx.JSON(http.StatusAccepted, model.Response{
		Message: "In progress",
	})
}

func upsertRecord(ctx echo.Context) error {
	var receive struct {
		Value float64 `json:"value"`
	}

	if err := ctx.Bind(&receive); err != nil {
		return err
	}

	year, _ := strconv.Atoi(ctx.Param("year"))
	record := model.Growth{
		Country:   ctx.Param("country"),
		Indicator: ctx.Param("indicator"),
		Year:      year,
		Value:     receive.Value,
	}

	srv.database.Upsert(record)
	return nil
}

func loadStatus(ctx echo.Context) error {
	var message string
	value, _ := process.Load("busy")

	if value.(bool) {
		message = "still in processing..."
	} else {
		message = "processing completed"
	}

	return ctx.JSON(http.StatusOK, struct {
		Message string `json:"msg"`
		Count   int    `json:"count"`
	}{
		Message: message,
		Count:   srv.database.Size(),
	})

}

func getRecord(ctx echo.Context) error {
	year, _ := strconv.Atoi(ctx.Param("year"))
	record := model.Growth{
		Country:   ctx.Param("country"),
		Indicator: ctx.Param("indicator"),
		Year:      year,
	}
	key := model.GenerateKey(record)
	result := srv.database.Read(key)
	if result == nil {
		srv.echo.Logger.Debug(result)
		return ctx.JSON(http.StatusOK, model.Response{
			Message: "No records have been found :(",
		})
	}
	return ctx.JSON(http.StatusOK, *result)
}

func deleteRecord(ctx echo.Context) error {
	year, _ := strconv.Atoi(ctx.Param("year"))
	record := model.Growth{
		Country:   ctx.Param("country"),
		Indicator: ctx.Param("indicator"),
		Year:      year,
	}
	key := model.GenerateKey(record)
	err := srv.database.Delete(key)
	if err != nil {
		srv.echo.Logger.Debug(err)
	}
	return nil
}

func dbSize(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, struct {
		Size int `json:"size"`
	}{
		Size: srv.database.Size(),
	})
}
