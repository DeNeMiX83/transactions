package user

import (
	"api/app/core/user"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	User_id string `json:"user_id"`
	Balanc_id string `json:"balanc_id"`
	Amount int `json:"amount"`
}

type Handler struct {
	uc user.UserUseCase
}

func NewHandler(uc user.UserUseCase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) IncreaseBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	requesData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error reading request body: %v", err)
	}
	var request Request
	json.Unmarshal(requesData, &request)

	user_id := request.User_id
	balanc_id := request.Balanc_id
	amount := request.Amount
	
	err = h.uc.IncreaseBalance(user_id, balanc_id, amount)
	if err != nil {
		resp["error"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp["message"] = "success"
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DecreaseBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)

	requesData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Error reading request body: %v", err)
	}

	var request Request
	json.Unmarshal(requesData, &request)

	user_id := request.User_id
	balanc_id := request.Balanc_id

	amount := request.Amount

	err = h.uc.DecreaseBalance(user_id, balanc_id, amount)
	if err != nil {
		resp["error"] = err.Error()
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp["message"] = "success"
	json.NewEncoder(w).Encode(resp)
}