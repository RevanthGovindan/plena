package routes

import (
	"net/http"
	"web3-tokeninfo/pkg/utils"
)

func fetchTokenInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(utils.Response))
}
