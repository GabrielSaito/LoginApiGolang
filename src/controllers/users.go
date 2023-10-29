package controllers

import (
	bd "api/src/BD"
	"api/src/auth"
	"api/src/models"

	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare("register"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := bd.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repository.NewRepositoryUsers(db)
	user.Id, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := bd.Conect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repository := repository.NewRepositoryUsers(db)
	users, err := repository.Search(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func SearchUserId(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
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
	user, err := repository.SearchId(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	userIdNoToken, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userId != userIdNoToken {
		responses.Error(w, http.StatusForbidden, errors.New("Não foi possivel atualizar!"))
		return
	}

	fmt.Println(userIdNoToken)

	bodyRequisition, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequisition, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	w.Write([]byte("Update user!"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdNoToken, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userId != userIdNoToken {
		responses.Error(w, http.StatusForbidden, errors.New("nao foi possivel deletar usuario"))
	}

	db, err := bd.Conect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repository.NewRepositoryUsers(db)
	if err = repository.Delete(userId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusInternalServerError, err)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	userIdNoToken, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if userIdNoToken != userId {
		responses.Error(w, http.StatusForbidden, errors.New("Nao foi possivel atualizar"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)

	var password models.PasswordModel
	if err = json.Unmarshal(bodyRequest, &password); err != nil {
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
	passwordInDB, err := repository.SearchPassword(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.CheckPassword(passwordInDB, password.OldPassword); err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Nao autorizado"))
		return
	}
	passwordWithHash, err := security.Hash(password.NewPassword)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = repository.NewPassword(userId, string(passwordWithHash)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)

}
