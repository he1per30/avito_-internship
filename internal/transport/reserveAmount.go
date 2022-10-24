package transport

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type Reserve struct {
	ID        int     `json:"id"`
	Cost      float64 `json:"cost"`
	OrderId   int     `json:"orderId"`
	ServiceId int     `json:"serviceId"`
}

type Resp struct {
	Message string  `json:"message"`
	Balance float64 `json:"balance"`
}

func (h *handler) ReserveAmount(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	b, _ := io.ReadAll(r.Body)
	var reserve Reserve
	err := json.Unmarshal(b, &reserve)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg 1 " + err.Error()))
		return
	}

	err = h.repository.ReserveAmount(reserve.ID, reserve.Cost, reserve.ServiceId, reserve.OrderId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg 2 " + err.Error()))
		return
	}

	balance, err := h.repository.GetBalance(reserve.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	var amount Resp
	amount.Message = "Success!"
	amount.Balance = balance
	fmt.Println(amount)

	resp, err := json.Marshal(&amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
