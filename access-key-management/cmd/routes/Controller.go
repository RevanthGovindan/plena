package routes

import (
	"access-key-management/internal/services"
	"encoding/json"
	"net/http"
)

func createNewAccessKeys(w http.ResponseWriter, r *http.Request) {
	resp, err := services.CreateNewAccessKeys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	respStr, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(respStr)
}

func deleteAccessKeys(w http.ResponseWriter, r *http.Request) {

}

func updateAccessKeys(w http.ResponseWriter, r *http.Request) {

}

func getAllAccessKeys(w http.ResponseWriter, r *http.Request) {

}

func fetchAccessKeys(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello"))
}

func disableAccessKeys(w http.ResponseWriter, r *http.Request) {

}
