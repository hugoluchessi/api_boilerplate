package controllers

import (
	"net/http"

	"github.com/hugoluchessi/api_boilerplate/config"
)

func HandleSomeDomainGetSomething(st *config.ServerTools, rw http.ResponseWriter, req *http.Request) {
	// Extract parameters and call business layer
}
