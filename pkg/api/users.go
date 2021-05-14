package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func (a *API) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	l := log.With(a.logger, "handler", "getusers")

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
