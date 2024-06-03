package main

import (
	"Media/handlers"
	"Media/infrastructure/routing"
	serve "Media/infrastructure/server"
	"Media/usecase"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	var logger = logrus.New()

	logger.Info("Program start")

	var usecase usecase.Case = *usecase.NewCase(logger)
	var handler *handlers.Handlers = handlers.NewHandlers(logger, &usecase)
	var router *routing.GoChiRouter = routing.NewGoChiRouting(logger, *handler)
	var server serve.Server = serve.NewServerHTTP(logger, ":8080", router)

	if err := server.Start(); err != nil {
		os.Exit(1)
	}

	// Запуск сервера в отдельной горутине
	go func() {
		if err := server.Start(); err != nil {
			logger.Fatalf("Error starting the server: %s", err)
		}
	}()

	// Создание канала для прослушивания сигналов системы
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Блокировка до получения сигнала
	<-stop

	logger.Info("Shutting down the server...")

	// Создание контекста с тайм-аутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Stop(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %s", err)
	}

	logger.Info("Server exiting")
}
