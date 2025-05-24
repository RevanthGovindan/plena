package routes

import (
	"net/http"
	"web3-tokeninfo/pkg/utils"
)

func fetchTokenInfo(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Context().Value(contextKey("userId")))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(utils.Response))
}
