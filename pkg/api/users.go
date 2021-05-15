package api

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
)

func (a *API) getUsersHandler(l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := a.store.GetUsers(r.Context())
		if err != nil {
			respondInteralError(w, l, fmt.Errorf("fetching users from store: %w", err))
			return
		}

		respondSuccess(w, l, "fetched users", users)
	}
}
