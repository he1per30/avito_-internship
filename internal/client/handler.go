package client

import (
	"avito/internal/handlers"
	"avito/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	transferURL = "/moneyTransfer/clientBalance"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(transferURL, h.moneyTransfer)
}

func (h *handler) moneyTransfer(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is moneyTransfer"))
}
