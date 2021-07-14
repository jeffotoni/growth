package api

import (
	"deryksr/model"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type server struct {
	echo     *echo.Echo
	port     string
	database *model.LocalDatabase
}

func StartServer(port int, database *model.LocalDatabase) {
	server := &server{
		echo:     echo.New(),
		port:     ":" + strconv.Itoa(port),
		database: database,
	}

	initRouter(server)

	server.echo.Logger.SetLevel(log.INFO)
	server.echo.Start(server.port)
}
