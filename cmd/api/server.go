package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func (app *application) server() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", app.config.domain, app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	return srv.ListenAndServe()
}
