package main

import (
	"github.com/sirupsen/logrus"
	"simple/backend/server/http"
)

func main() {
	if err := http.InitWebServer(); err != nil {
		logrus.Fatal(err)
	}
}
