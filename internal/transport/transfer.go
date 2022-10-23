package transport

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type IncrementBalanceReq struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
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

	w.Write([]byte("Ok"))
}
