package controllers

import (
	bd "api/src/BD"
	"api/src/auth"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"

	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequisition, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequisition, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := bd.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	userSaveDB, err := repository.SearchForEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.CheckPassword(userSaveDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(userSaveDB.Id)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)
	w.Write([]byte(token))

}
