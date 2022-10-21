package main

import (
	"avito/internal/client"
	"avito/internal/config"
	"avito/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	handler := client.NewHandler(logger)
	handler.Register(router)
	logger.Info("register client handler")

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start app")
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Infof("start server on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	logger.Fatalln(server.Serve(listener))
}
