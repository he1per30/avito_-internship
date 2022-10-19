package client

import (
	"avito/internal/handlers"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	transferURL = "/moneyTransfer/clientBalance"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(transferURL, h.moneyTransfer)
}

func (h *handler) moneyTransfer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is moneyTransfer"))
}
