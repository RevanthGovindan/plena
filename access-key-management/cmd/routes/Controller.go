package routes

import (
	"access-key-management/internal/models"
	"access-key-management/internal/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func createNewAccessKeys(w http.ResponseWriter, r *http.Request) {
	resp, err := services.CreateNewAccessKeys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	respStr, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(respStr)
}

func deleteAccessKeys(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.DeleteAccessKeys(vars["keyId"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func updateAccessKeys(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var keyData models.UpdateAccessKeyRequest
	err := json.NewDecoder(r.Body).Decode(&keyData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	err = services.UpdateAccessKeys(vars["keyId"], keyData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getAllAccessKeys(w http.ResponseWriter, r *http.Request) {
	data, err := services.GetAllAccessKeys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed"))
		return
	}
	respStr, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(respStr)
}

func fetchAccessKeys(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data, err := services.GetDataByAccessKey(vars["keyId"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	respStr, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(respStr)
}

func disableAccessKeys(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := services.DisableAccessKey(vars["keyId"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}
