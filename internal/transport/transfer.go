package transport

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type IncrementBalanceReq struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
}

type IncrementBalanceResp struct {
	Message string  `json:"message"`
	Balance float64 `json:"balance"`
}

func (h *handler) IncrementBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	b, _ := io.ReadAll(r.Body)
	var ibr IncrementBalanceReq
	err := json.Unmarshal(b, &ibr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	err = h.repository.IncrementBalance(ibr.ID, ibr.Amount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}
	balance, err := h.repository.GetBalance(ibr.ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	var incr IncrementBalanceResp
	incr.Message = "Success!"
	incr.Balance = balance
	fmt.Println(incr)

	resp, err := json.Marshal(&incr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)

}
