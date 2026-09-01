package apiserver

import (
	"avito/internal/api/info"
	coinstransfer "avito/internal/repository/coins_transfer"
	"avito/internal/repository/users"
	"context"

	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Start server func create.
func Start(ctx context.Context, dbCfg *DbConfig, config *Config) error {
	logger := logrus.New()
	conn, err := openDB(ctx, dbCfg)
	if err != nil {
		return err
	}

	defer conn.Close()

	coinsTransferRepo := coinstransfer.New(conn, logger)
	userRepo := users.New(conn, logger)
	userHandler := info.New(&userRepo, &coinsTransferRepo, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/info", userHandler.Info)
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

func openDB(ctx context.Context, dbCfg *DbConfig) (*pgxpool.Pool, error) {
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

// type modelsTransfer struct {
// 	id         string
// 	senderId   string
// 	receiverId string
// 	amount     int
// 	createdAt  time.Time
// }

// type handlerTransfer struct {
// 	amount    int
// 	direction string // out, in
// 	person    string
// }

// func FromModelToHandler(mList []modelsTransfer, id string) []handlerTransfer {
// 	htList := make([]handlerTransfer, 0)
// 	_ = htList
// 	for _, t := range mList {
// 		ht := handlerTransfer{
// 			amount: t.amount,
// 		}
// 		if id == t.receiverId {
// 			ht.direction = "in"
// 			ht.person = t.senderId
// 		}
// 		if id == t.senderId {
// 			ht.direction = "out"
// 			ht.person = t.receiverId
// 		}
// 	}

// 	return nil
// }

// type TransferService interface {
// 	GetByUserID(ctx context.Context, userID string) ([]modelsTransfer, error)
// }

// type Handler struct {
// 	transferService TransferService
// }

// func (h *Handler) TransfersByUserID(w http.ResponseWriter, r *http.Request) {
// 	id := r.PathValue("id")               // "users/:id/tranfers/:direction"
// 	direction := r.PathValue("direction") // "users/:id/tranfers/:direction"

// 	transfers, err := h.transferService.GetByUserID(r.Context(), id)
// 	if err != nil {

// 	}

// 	allHandlerTransfers := FromModelToHandler(transfers, id)

// 	responseTransfers := make([]handlerTransfer, 0)
// 	for _, ht := range allHandlerTransfers {
// 		if ht.direction != direction {
// 			continue
// 		}
// 		responseTransfers = append(responseTransfers, ht)
// 	}

// 	resp, err := json.Marshal(responseTransfers)
// 	if err != nil {
// 		panic(err)
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)
// }
