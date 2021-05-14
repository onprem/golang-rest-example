package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/onprem/go-db-example/pkg/store"
)

type API struct {
	store  *store.Store
	logger log.Logger
}

func New(store *store.Store, logger log.Logger) *API {
	return &API{
		store:  store,
		logger: logger,
	}
}

func (a *API) Register(r chi.Router) {
	a.registerRoutes(r)
}
