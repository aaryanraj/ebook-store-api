package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/aaryanraj/ebook-store-api/pkg/clients"
	"github.com/aaryanraj/ebook-store-api/pkg/common"
	"github.com/aaryanraj/ebook-store-api/pkg/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &clients.User{}
	if err := utils.ParseBody(r, user); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	common.Prepare(user)
	common.Validate(user, "login")
	book, _ := user.SaveUser()
	resp, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
