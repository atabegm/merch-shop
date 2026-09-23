package apiserver

import (
	"avito/internal/api/auth"
	"avito/internal/api/auth/middleware"
	"avito/internal/api/buy"
	"avito/internal/api/info"
	sendcoin "avito/internal/api/send_coin"
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/merch"
	"avito/internal/repository/purchases"
	"avito/internal/repository/users"
	authservice "avito/internal/service/auth_service"
	buyservice "avito/internal/service/buy"
	sendcoinservice "avito/internal/service/send_coin_service"
	"net/http"

	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	defaultReadTimeout       = 5 * time.Second
	defaultReadHeaderTimeout = 6 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleTimeout       = 7 * time.Second
)

func Start(ctx context.Context, config *Config, dbCfg *DBConfig, logger *logrus.Logger) error {
	conn, err := openDB(ctx, dbCfg)
	if err != nil {
		return fmt.Errorf("error with connect with db: %w", err)
	}

	usersRepo := users.New(conn, logger)
	merchRepo := merch.New(conn, logger)
	purchasesRepo := purchases.New(conn, logger)
	coinTransfersRepo := coinstransfer.New(conn, logger)

	jwtSecret := "my-secret-key"

	authService := authservice.New(&usersRepo, jwtSecret)
	buyService := buyservice.New(&purchasesRepo, &merchRepo)
	sendCoinsService := sendcoinservice.New(&usersRepo, &coinTransfersRepo)

	authHandler := auth.New(authService, logger, jwtSecret)
	infoHandler := info.New(&usersRepo, &coinTransfersRepo, &purchasesRepo, &merchRepo, logger)
	buyHandler := buy.New(buyService, logger)
	sendCoinHandler := sendcoin.New(sendCoinsService, logger)

	router := http.NewServeMux()
	router.Handle("GET /api/info", middleware.Auth(http.HandlerFunc(infoHandler.Info), []byte(jwtSecret)))
	router.Handle("POST /api/buy/{item}", middleware.Auth(http.HandlerFunc(buyHandler.Buy), []byte(jwtSecret)))
	router.Handle("POST /api/sendCoin", middleware.Auth(http.HandlerFunc(sendCoinHandler.Send), []byte(jwtSecret)))
	router.HandleFunc("POST /api/auth", authHandler.Auth)

	srv := http.Server{
		Addr:              config.BindAddr,
		Handler:           router,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		ReadHeaderTimeout: defaultReadHeaderTimeout,
	}

	logger.Info("start server")

	return srv.ListenAndServe()
}

func openDB(ctx context.Context, dbCfg *DBConfig) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, fmt.Errorf("error with pool parse config: %w", err)
	}

	config.ConnConfig.Host = dbCfg.PgHost
	config.ConnConfig.Port = dbCfg.PgPort
	config.ConnConfig.Database = dbCfg.PgDatabase
	config.ConnConfig.User = dbCfg.PgUser
	config.ConnConfig.Password = dbCfg.PgPassword

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("error with ping pool: %w", err)
	}

	return pool, nil
}
