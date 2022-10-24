package transport

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

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

	w.Write([]byte(fmt.Sprintf("userBalance: %g", balance)))
}
