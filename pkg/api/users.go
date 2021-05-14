package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func (a *API) getUsersHandler(l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := a.store.GetUsers(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "error",
				"msg":    "something went wrong",
			})
			level.Error(l).Log("msg", "getting users from store", "err", err)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"users":  users,
		})
	}
}
