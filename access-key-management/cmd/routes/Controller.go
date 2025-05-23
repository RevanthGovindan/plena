package routes

import "net/http"

func createNewAccessKeys(w http.ResponseWriter, r *http.Request) {

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
