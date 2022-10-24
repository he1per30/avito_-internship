package transport

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type Response struct {
	UserBalance float64 `json:"userBalance"`
}

func (h *handler) GetUserBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("id")
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid userID " + err.Error()))
		return
	}

	balance, err := h.repository.GetBalance(intUserId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	var Response Response
	Response.UserBalance = balance
	b, err := json.Marshal(&Response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
