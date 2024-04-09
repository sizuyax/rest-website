package main

import (
	"simple/backend/logger"
	"simple/backend/server/http"
)

func main() {
	server, err := http.InitWebServer()
	if err != nil {
		logger.Logger.Fatal(err)
	}

	if err := server.StartServer(); err != nil {
		logger.Logger.Fatal(err)
	}
}
