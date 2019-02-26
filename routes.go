package main

import (
	"net/http"

	"github.com/hugoluchessi/api_boilerplate/config"
	"github.com/hugoluchessi/api_boilerplate/controllers"
	"github.com/hugoluchessi/badger"
	"github.com/hugoluchessi/gobservable/logging"
	"github.com/hugoluchessi/gobservable/tctx"
)

func ConfigureRoutes(st *config.ServerTools) *badger.Mux {
	// Create new Mux
	mux := badger.NewMux()

	// Available Middlewares
	transMw := tctx.NewTransactionContextMiddleware()
	logMw := logging.NewContextLoggerMiddleware(st.Logger)

	// Create new router group
	mainRouter := mux.AddRouter("/")
	mainRouter.Use(logMw.Handler)
	mainRouter.Use(transMw.Handler)

	mainRouter.Get("", http.HandlerFunc(healthCheckHandler))
	mainRouter.Get("status", http.HandlerFunc(statusHandler))

	mainRouter.Get("someDomain", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		controllers.HandleSomeDomainGetSomething(st, res, req)
	}))

	return mux
}

func healthCheckHandler(res http.ResponseWriter, req *http.Request) {
	// Check external dependencies and return 500 in case anything goes wrong
}

func statusHandler(res http.ResponseWriter, req *http.Request) {
	// Check external dependencies and return a report with all
	//  dependencies statuses with code 200
}
