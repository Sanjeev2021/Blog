package Login

import (
	"encoding/json"
	"net/http"

	Login "blog/login/service"

)

type LoginInfo struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var logininfo LoginInfo

	err := json.NewDecoder(r.Body).Decode(&logininfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call for login
	login, err := Login.Authenticate(logininfo.Email, logininfo.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(login)
}
