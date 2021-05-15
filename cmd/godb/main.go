package main

import (
	"context"
	"flag"
	"net/http"
	"os"

	"github.com/metalmatze/signal/server/signalhttp"
	"github.com/onprem/go-db-example/ent"
	"github.com/onprem/go-db-example/pkg/api"
	"github.com/onprem/go-db-example/pkg/store"

	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/mattn/go-sqlite3"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var dbPath, address string
	flag.StringVar(&dbPath, "db", "godb.db", "Path to the sqlite database file.")
	flag.StringVar(&address, "address", "0.0.0.0:8080", "The address to bind the HTTP server.")

	flag.Parse()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	client, err := ent.Open("sqlite3", "file:"+dbPath+"?cache=shared&mode=rwc&_fk=1")
	if err != nil {
		level.Error(logger).Log("err", err)
		return
	}
	level.Info(logger).Log("msg", "successfully connected to database")

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		level.Error(logger).Log("msg", "auto migration", "err", err)
	}

	r := chi.NewRouter()

	str := store.New(client)
	apy := api.New(str, log.With(logger, "component", "api"), prometheus.DefaultRegisterer)

	r.Handle("/metrics", promhttp.Handler())

	ins := signalhttp.NewHandlerInstrumenter(prometheus.DefaultRegisterer, []string{"group", "handler"})

	r.Route("/api", func(r chi.Router) {
		apy.Register(r, ins)
	})

	level.Info(logger).Log("msg", "starting web server", "addr", address)
	if err := http.ListenAndServe(address, r); err != nil && err != http.ErrServerClosed {
		level.Error(logger).Log("msg", "http server listen", "err", err)
	}
}
