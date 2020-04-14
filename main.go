package main

import (
	"net/http"

	"github.com/hugoluchessi/api_boilerplate/config"
)

func main() {
	st, err := config.NewServerTools()

	if err != nil {
		panic(err)
	}

	mux := ConfigureRoutes(st)

	http.ListenAndServe(":8080", mux)
}
