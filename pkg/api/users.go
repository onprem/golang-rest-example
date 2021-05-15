package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/onprem/go-db-example/ent"
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

func (a *API) createUserHandler(l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			respondErr(w, l, http.StatusBadRequest, "invalid input")
			return
		}

		usr, err := a.store.CreateUser(r.Context(), input.Name, input.Age)
		if err != nil {
			respondInteralError(w, l, fmt.Errorf("creating user: %w", err))
			return
		}

		a.userCreateCounter.Inc()
		respondSuccess(w, l, "successfully created user", usr)
	}
}

func (a *API) deleteUserHandler(l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID int `json:"id"`
		}

		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			respondErr(w, l, http.StatusBadRequest, "invalid input")
			return
		}

		if err := a.store.DeleteUser(r.Context(), input.ID); err != nil {
			if ent.IsNotFound(err) {
				respondErr(w, l, http.StatusBadRequest, "invalid user id")
				return
			}

			respondInteralError(w, l, fmt.Errorf("deleting user: %w", err))
			return
		}

		respondSuccess(w, l, "successfully deleted user", nil)
	}
}

func (a *API) updateUserHandler(l log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			ID   int     `json:"id"`
			Name *string `json:"name"`
			Age  *int    `json:"age"`
		}

		defer r.Body.Close()
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			respondErr(w, l, http.StatusBadRequest, "invalid input")
			return
		}

		usr, err := a.store.UpdateUser(r.Context(), input.ID, input.Name, input.Age)
		if err != nil {
			if ent.IsNotFound(err) {
				respondErr(w, l, http.StatusBadRequest, "invalid user id")
				return
			}

			respondInteralError(w, l, fmt.Errorf("deleting user: %w", err))
			return
		}

		respondSuccess(w, l, "successfully updated user", usr)
	}
}
