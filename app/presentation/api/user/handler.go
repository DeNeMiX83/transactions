package user

import (
	"api/app/core/user"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Request struct {
	User_id string `json:"user_id"`
	Balanc_id string `json:"balanc_id"`
	Amount int `json:"amount"`
}

type handler struct {
	uc user.UserUseCase
}

func NewHandler(uc user.UserUseCase) *handler {
	return &handler{uc: uc}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST("/user/balance/increase", h.IncreaseBalance)
	router.POST("/user/balance/decrease", h.DecreaseBalance)
}

func (h *handler) IncreaseBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		w.WriteHeader(http.StatusBadRequest)
		resp["error"] = err.Error()
	} else {
		resp["message"] = "success"
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *handler) DecreaseBalance(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
		w.WriteHeader(http.StatusBadRequest)
		resp["error"] = err.Error()
	} else {
		resp["message"] = "success"
	}
	json.NewEncoder(w).Encode(resp)
}