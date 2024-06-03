package routing

import "net/http"

type Router interface {
	Route(w http.ResponseWriter, r *http.Request)
}
