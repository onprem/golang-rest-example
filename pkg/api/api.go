package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/metalmatze/signal/server/signalhttp"
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

func (a *API) Register(r chi.Router, ins signalhttp.HandlerInstrumenter) {
	a.registerRoutes(r, ins)
}

type res struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}

func respond(w io.Writer, l log.Logger, resp res) {
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		level.Warn(l).Log("msg", "json encoding failed", "err", err)
	}
}

func respondErr(w http.ResponseWriter, l log.Logger, code int, msg string) {
	w.WriteHeader(code)
	respond(w, l, res{
		Status: "error",
		Msg:    msg,
	})
}

func respondInteralError(w http.ResponseWriter, l log.Logger, err error) {
	respondErr(w, l, http.StatusInternalServerError, "something went wrong")
	level.Warn(l).Log("msg", "internal server error", "err", err)
}

func respondSuccess(w http.ResponseWriter, l log.Logger, msg string, data interface{}) {
	respond(w, l, res{
		Status: "success",
		Msg:    msg,
		Data:   data,
	})
}
