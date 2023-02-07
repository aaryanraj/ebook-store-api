package controllers

import (
	"net/http"

	"github.com/aaryanraj/ebook-store-api/pkg/auth"
	"github.com/aaryanraj/ebook-store-api/pkg/clients"
	"github.com/aaryanraj/ebook-store-api/pkg/common"
	"github.com/aaryanraj/ebook-store-api/pkg/responses"
	"github.com/aaryanraj/ebook-store-api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &clients.User{}
	if err := utils.ParseBody(r, user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	common.Prepare(user)
	if err := common.Validate(user, "login"); err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusOK, token)
}

func SignIn(email, password string) (string, error) {
	var err error
	user := &clients.User{}

	foundUser, err := user.FindUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = common.VerifyPassword(foundUser.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(uint32(foundUser.ID))
}
