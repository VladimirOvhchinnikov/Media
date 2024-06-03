package serve

import (
	"context"
	"net/http"
)

type Server interface {
	Start() error
	Stop(context.Context) error
	Restart() error

	SetRouter() error
	Configure(options map[string]interface{}) error
}

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
