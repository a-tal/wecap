package internal

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func newMux() *mux.Router {
	m := mux.NewRouter()

	m.Handle("/metrics", promhttp.Handler())
	m.PathPrefix(handlerPath).HandlerFunc(reportHandler)
	m.PathPrefix("/").HandlerFunc(catchAllHandler)

	m.Use(midware)

	return m
}

func newServer() *http.Server {
	addr := "0.0.0.0:8000"

	srv := &http.Server{
		Handler: newMux(),
		Addr:    addr,

		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			ll.Println(err)
		}
	}()

	ll.Printf("listening on %s", addr)

	return srv
}
