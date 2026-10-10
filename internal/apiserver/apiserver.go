package apiserver

import (
	"avito/internal/api"
	"avito/internal/api/middleware"
	"avito/internal/kafka"
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/merch"
	"avito/internal/repository/purchases"
	"avito/internal/repository/users"
	"avito/internal/service"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

const (
	defaultReadTimeout       = 5 * time.Second
	defaultReadHeaderTimeout = 6 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleTimeout       = 7 * time.Second
)

// Start create.
func Start(ctx context.Context, config *Config, dbCfg *DBConfig, logger *logrus.Logger) error {
	conn, err := openDB(ctx, dbCfg)
	if err != nil {
		return fmt.Errorf("error with connect with db: %w", err)
	}

	trManager := manager.Must(pgxv5.NewDefaultFactory(conn))
	getter := pgxv5.DefaultCtxGetter

	usersRepo := users.New(conn, logger, getter)
	merchRepo := merch.New(conn, logger)
	purchasesRepo := purchases.New(
		getter,
		conn,
		logger,
	)

	coinTransfersRepo := coinstransfer.New(conn, logger)

	jwtSecret := "my-secret-key"

	producer := kafka.NewProducer(
		config.KafkaBroker,
		config.KafkaTopic,
	)

	defer producer.Close()

	service := service.New(
		&usersRepo,
		&coinTransfersRepo,
		&purchasesRepo,
		&merchRepo,
		trManager,

		producer,

		jwtSecret,
	)

	handler := api.New(
		service,
		logger,
	)

	router := http.NewServeMux()
	router.Handle("POST /api/send", middleware.Auth(http.HandlerFunc(handler.Send), []byte(jwtSecret)))
	router.Handle("GET /api/buy/{item}", middleware.Auth(http.HandlerFunc(handler.Buy), []byte(jwtSecret)))
	router.HandleFunc("POST /api/auth", handler.Auth)
	router.Handle("GET /api/info", middleware.Auth(http.HandlerFunc(handler.Info), []byte(jwtSecret)))

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
		return nil, fmt.Errorf("apiserver: %w", err)
	}

	config.ConnConfig.Host = dbCfg.PgHost
	config.ConnConfig.Port = dbCfg.PgPort
	config.ConnConfig.Database = dbCfg.PgDatabase
	config.ConnConfig.User = dbCfg.PgUser
	config.ConnConfig.Password = dbCfg.PgPassword

	pool, err := pgxpool.NewWithConfig(
		ctx,
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("apiserver: %w", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("apiserver: %w", err)
	}

	return pool, nil
}
