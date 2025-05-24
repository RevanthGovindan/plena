package routes

import (
	"net/http"
	"web3-tokeninfo/pkg/middlewares"

	"github.com/gorilla/mux"
)

func GetRoutes(r *mux.Router) *mux.Router {
	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.RequestAuthenticator)
	r.HandleFunc("/tokens/{tokenId}", fetchTokenInfo).Methods(http.MethodGet)
	return r
}
