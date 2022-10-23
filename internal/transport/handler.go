package transport

import (
	"avito/internal/handlers"
	"avito/internal/user"
	"avito/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	transferURL = "/IncrementBalance/clientBalance"
	reserveURL  = "/reserveAmount"
)

type handler struct {
	logger     *logging.Logger
	repository user.Repository
}

func NewHandler(logger *logging.Logger, repository user.Repository) handlers.Handler {
	return &handler{
		logger:     logger,
		repository: repository,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(transferURL, h.IncrementBalance)
	router.POST(reserveURL, h.ReserveAmount)

}
