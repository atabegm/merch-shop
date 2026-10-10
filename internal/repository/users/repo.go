package users

import (
	pgxv5 "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// Repo for users create.
type Repo struct {
	pool   *pgxpool.Pool
	logger *logrus.Logger
	getter *pgxv5.CtxGetter
}

// New repo for users create.
func New(
	pool *pgxpool.Pool,
	logger *logrus.Logger,
	getter *pgxv5.CtxGetter,
) Repo {
	return Repo{
		pool:   pool,
		getter: getter,
		logger: logger,
	}
}
