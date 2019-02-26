package main

import (
	"net/http"
	"os"

	"github.com/hugoluchessi/api_boilerplate/config"
	"github.com/hugoluchessi/gobservable/logging"
)

func main() {
	cfg := logging.LoggerConfig{Output: os.Stdout}
	cfgs := []logging.LoggerConfig{cfg}

	zap := logging.NewZapLogger(cfgs)
	l := logging.NewContextLogger(zap)

	st := config.NewServerTools(l)

	mux := ConfigureRoutes(st)

	http.ListenAndServe(":8080", mux)
}
