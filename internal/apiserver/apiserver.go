package apiserver

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

const (
	defaultPingContext = 5 * time.Second
)

// Start server func create.
func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:              config.BindAddr,
		Handler:           mux,
		ReadTimeout:       defaultReadTimeout,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		WriteTimeout:      defaultWriteTimeout,
		IdleTimeout:       defaultIdleimeout,
	}

	err = srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("cant connect to db. error:%w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultPingContext)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
