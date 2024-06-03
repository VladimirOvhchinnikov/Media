package serve

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ServerHTTP struct {
	Logger  *logrus.Logger
	Address string
	server  *http.Server
}

func NewServerHTTP(logger *logrus.Logger, address string, handler http.Handler) *ServerHTTP {
	return &ServerHTTP{
		Logger:  logger,
		Address: address,
		server: &http.Server{
			Addr:    address,
			Handler: handler,
		},
	}
}

func (sh *ServerHTTP) Start() error {
	sh.Logger.Info("Starting the server at the address ", sh.Address)
	err := http.ListenAndServe(sh.Address, sh.server.Handler)
	if err != nil {
		sh.Logger.Error("Error in starting the HTTP server ", err)
		return err
	}
	sh.Logger.Info("The server has been successfully started")
	return nil
}

func (sh *ServerHTTP) Stop(ctx context.Context) error {
	sh.Logger.Info("Shutting down the server...")
	err := sh.server.Shutdown(ctx)
	if err != nil {
		sh.Logger.Error("Error shutting down the server: ", err)
		return err
	}
	sh.Logger.Info("Server has been shut down successfully")
	return nil
}

func (sh *ServerHTTP) Restart() error {
	return nil
}

func (sh *ServerHTTP) SetRouter() error {
	return nil
}

func (sh *ServerHTTP) Configure(options map[string]interface{}) error {
	return nil
}
