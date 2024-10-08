package main

import (
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"main.go/internal/user"
	"main.go/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()
	logger.Info("start application")
	listener, err := net.Listen("tcp", "127.0.0.1:1234")

	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Info("server is listening port 0.0.0.0:1234")
	logger.Fatal(server.Serve(listener))
}
