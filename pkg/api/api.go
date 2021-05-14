package api

import (
	"github.com/onprem/go-db-example/pkg/store"
)

type API struct {
	store *store.Store
}

func New(store *store.Store) *API {
	return &API{
		store: store,
	}
}
