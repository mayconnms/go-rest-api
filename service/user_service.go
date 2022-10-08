package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest-api/domain"
	"rest-api/repository"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	repository.DeleteUser(uint(id))

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	user := domain.User{ID: uint(id)}
	json.NewEncoder(w).Encode(user)

}

func CreatUser(w http.ResponseWriter, r *http.Request) {

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	user := domain.User{}
	json.Unmarshal(d, &user)

	repository.SaveUser(&user)

}
