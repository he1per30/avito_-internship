package transport

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strings"
)

type Report struct {
	Date      string `json:"date"`
	ServiceId int    `json:"serviceId"`
}

func (h *handler) GetReport(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	b, err := io.ReadAll(r.Body)
	var reportReq Report
	err = json.Unmarshal(b, &reportReq)
	fmt.Println(reportReq.Date)
	splitDate := strings.Split(reportReq.Date, "-")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}
	err = h.repository.GetReport(splitDate[0], splitDate[1], reportReq.ServiceId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid msg " + err.Error()))
		return
	}

	w.Write([]byte("Ok"))
}
