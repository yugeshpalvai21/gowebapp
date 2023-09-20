package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct{}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := make(map[string]string)
	resp["message"] = "Status OK"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("There is error in controler")
	}

	w.Write(jsonResp)
	return
}
