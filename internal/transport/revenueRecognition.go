package transport

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type Revenue struct {
	ID        int     `json:"id"`
	Cost      float64 `json:"cost"`
	OrderId   int     `json:"orderId"`
	ServiceId int     `json:"serviceId"`
}

func (h *handler) RevenueRecognition(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	b, _ := io.ReadAll(r.Body)
	var reserve Reserve
	err := json.Unmarshal(b, &reserve)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	err = h.repository.RevenueRecognition(reserve.ID, reserve.Cost, reserve.ServiceId, reserve.OrderId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}
