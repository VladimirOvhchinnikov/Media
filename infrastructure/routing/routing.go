package routing

import (
	"Media/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

type GoChiRouter struct {
	Logger *logrus.Logger
	Router chi.Router
	//Handler handlers.Handlers
}

func NewGoChiRouting(logger *logrus.Logger, handler handlers.Handlers) *GoChiRouter {
	router := chi.NewRouter()
	router.Post("/culc", handler.CulcHandler)

	return &GoChiRouter{
		Logger: logger,
		Router: router,
	}
}

func (gc *GoChiRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gc.Router.ServeHTTP(w, r)
}
