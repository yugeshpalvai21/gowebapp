package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PagesController struct {
}

func (pc *PagesController) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	resp := make(map[string]string)
	resp["message"] = "Status OK"
	resp["AdditionalInfo"] = "Root Page COnfigured Successfully"

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("There is error in controler")
	}

	w.Write(jsonResp)
	return
}
