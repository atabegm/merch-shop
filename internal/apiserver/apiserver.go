package apiserver

import (
	"avito/internal/api/auth"
	"avito/internal/api/info"
	sendcoin "avito/internal/api/send_coin"
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/merch"
	"avito/internal/repository/purchases"
	"avito/internal/repository/users"
	"context"
	"time"

	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	defaultReadTimeout       = 5 * time.Second
	defaultReadHeaderTimeout = 6 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleimeout        = 7 * time.Second
)

// Start server func create.
func Start(ctx context.Context, dbCfg *DBConfig, config *Config) error {
	logger := logrus.New()
	conn, err := openDB(ctx, dbCfg)
	if err != nil {
		return err
	}

	defer conn.Close()

	coinsTransferRepo := coinstransfer.New(conn, logger)
	userRepo := users.New(conn, logger)
	purchasesRepo := purchases.New(conn, logger)
	merchRepo := merch.New(conn, logger)

	jwtSecret := "secret key"
	userHandler := info.New(&userRepo, &coinsTransferRepo, &purchasesRepo, &merchRepo,
		logger)

	authHandler := auth.New(&userRepo, logger, jwtSecret)
	sendCoinHandler := sendcoin.New(&userRepo, &coinsTransferRepo, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/info", userHandler.Info)
	mux.HandleFunc("POST /api/auth", authHandler.Auth)
	mux.HandleFunc("POST /api/sendCoin", sendCoinHandler.Send)
	srv := &http.Server{
		Addr:              config.BindAddr,
		Handler:           mux,
		ReadTimeout:       defaultReadTimeout,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		WriteTimeout:      defaultWriteTimeout,
		IdleTimeout:       defaultIdleimeout,
	}

	return srv.ListenAndServe()
}

func openDB(ctx context.Context, dbCfg *DBConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("openDB config parse  error:%w", err)
	}

	config.ConnConfig.Host = dbCfg.PgHost
	config.ConnConfig.Port = dbCfg.PgPort
	config.ConnConfig.Database = dbCfg.PgDatabase
	config.ConnConfig.User = dbCfg.PgUser
	config.ConnConfig.Password = dbCfg.PgPassword

	fmt.Println("config user:", dbCfg.PgUser, "config database: ", dbCfg.PgDatabase)

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return pool, nil
}
