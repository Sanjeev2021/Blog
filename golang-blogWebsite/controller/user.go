package Controller

import (
	"encoding/json"
	"net/http"

	"blog/model"
	"blog/service"

)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	//parse request body
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//service call for new user
	user, err := service.NewUser(newUser.ID, newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// update username handler
func UpdateUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var updateUser model.User
	err := json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.UpdateUsername(&updateUser, updateUser.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Delete user handler
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var deleteUser model.User
	err := json.NewDecoder(r.Body).Decode(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.DeleteUser(&deleteUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// get user
func GetUser(w http.ResponseWriter, r *http.Request) {
	var getUser model.User
	err := json.NewDecoder(r.Body).Decode(&getUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	getuser, err := service.GetUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(getuser)
}
