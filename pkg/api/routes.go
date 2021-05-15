package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/metalmatze/signal/server/signalhttp"
	"github.com/prometheus/client_golang/prometheus"
)

func (a *API) registerRoutes(r chi.Router, ins signalhttp.HandlerInstrumenter) {
	r.Get("/users", a.newHandler("getUsers", ins, a.getUsersHandler))
	r.Post("/user", a.newHandler("createUser", ins, a.createUserHandler))
	r.Delete("/user", a.newHandler("deleteUser", ins, a.deleteUserHandler))
	r.Patch("/user", a.newHandler("updateUser", ins, a.updateUserHandler))
}

func (a *API) newHandler(
	id string,
	ins signalhttp.HandlerInstrumenter,
	fn func(log.Logger) http.HandlerFunc,
) http.HandlerFunc {
	return ins.NewHandler(
		prometheus.Labels{"group": "api", "handler": id},
		fn(log.With(a.logger, "handler", id)),
	)
}
