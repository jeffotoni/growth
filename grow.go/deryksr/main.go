package main

import (
	"deryksr/api"
	"deryksr/model"
)

func main() {
	api.StartServer(8080, model.NewDatabase())
}
