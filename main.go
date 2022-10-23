package main

import (
	"avito/internal/config"
	"avito/internal/transport"
	userDb "avito/internal/user/db"
	"avito/pkg/client/postgresql"
	"avito/pkg/logging"
	"context"
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

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	repository := userDb.NewRepository(postgreSQLClient, logger, postgreSQLClient)

	handler := transport.NewHandler(logger, repository)
	handler.Register(router)
	logger.Info("register user handler")

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
