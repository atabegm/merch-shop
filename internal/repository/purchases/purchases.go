package purchases

import (
	pgxv5 "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Repo for purchase create.
type Repo struct {
	getter *pgxv5.CtxGetter
	pool   *pgxpool.Pool
	logger *logrus.Logger
}

// New repo for merch create.
func New(
	getter *pgxv5.CtxGetter,
	pool *pgxpool.Pool,
	logger *logrus.Logger,
) Repo {
	return Repo{
		getter: getter,
		pool:   pool,
		logger: logger,
	}
}
