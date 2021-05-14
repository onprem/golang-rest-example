package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
)

func (a *API) registerRoutes(r chi.Router) {
	r.Get("/users", a.getUsersHandler(log.With(a.logger, "handler", "getusers")))
}
