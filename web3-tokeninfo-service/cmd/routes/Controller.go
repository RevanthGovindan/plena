package routes

import "net/http"

func fetchTokenInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("from web3"))
}
